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

	log.Logf("Downloading from %s", imageLink)
	log.Logf("With Filename : %s", filePath)

	i.downloadFile(imageLink, filePath)
}

func (i *ImageHandler) IsSupported() bool {
	return true
}

func (i *ImageHandler) NotSupportedMessage() string {
	return "Image not supported."
}

func (i *ImageHandler) downloadFile(imageLink string, fileName string) {
	response, err := client.Get(imageLink)

	if nil != err {
		log.Error("Failed to retrieve image", imageLink, err)
		semaphore.GetWaitGroup().Done()

		return
	}

	defer response.Body.Close()

	err = writer.Write(fileName, response.Body)

	if err != nil {
		log.Error("Failed to create file", fileName, err)
		semaphore.GetWaitGroup().Done()

		return
	}

	log.Successf("%s downloaded.", fileName)
	semaphore.GetWaitGroup().Done()
}
