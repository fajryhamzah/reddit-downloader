package media

import (
	"fmt"
	"path"

	"github.com/fajryhamzah/reddit-downloader/src/client"
	"github.com/fajryhamzah/reddit-downloader/src/data"
	"github.com/fajryhamzah/reddit-downloader/src/log"
	"github.com/fajryhamzah/reddit-downloader/src/semaphore"
	"github.com/fajryhamzah/reddit-downloader/src/writer"
)

type ImageHandler struct {
}

const DEFAULT_IMAGE_DOWNLOAD_PATH string = "result/images"

func (i *ImageHandler) Handle(response data.MainResponse) {
	childrenResponse := response.Data.Children[0].Data
	imageLink := childrenResponse.UrlDestination
	filename := path.Base(imageLink)

	filePath := fmt.Sprintf("%s/%s_%s_%s_%s", DEFAULT_IMAGE_DOWNLOAD_PATH, childrenResponse.Subreddit, childrenResponse.Title, childrenResponse.Author, filename)

	log.Logf("Downloading from %s/%s", childrenResponse.SubredditPrefix, childrenResponse.Title)
	log.Logf("With Filename : %s", filePath)

	i.downloadFile(imageLink, filePath)
}

func (i *ImageHandler) downloadFile(imageLink string, fileName string) {
	response, err := client.Get(imageLink)

	if nil != err {
		log.Error("Failed to retrieve image", imageLink)
		semaphore.GetWaitGroup().Done()

		return
	}

	err = writer.Write(fileName, response.Body)

	if err != nil {
		log.Error("Failed to create file", fileName)
		semaphore.GetWaitGroup().Done()

		return
	}

	log.Successf("%s downloaded.", fileName)
	semaphore.GetWaitGroup().Done()
}
