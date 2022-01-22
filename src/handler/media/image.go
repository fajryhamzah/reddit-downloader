package media

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/fajryhamzah/reddit-downloader/src/client"
	"github.com/fajryhamzah/reddit-downloader/src/data"
	"github.com/fajryhamzah/reddit-downloader/src/semaphore"
)

type ImageHandler struct {
}

const DEFAULT_IMAGE_DOWNLOAD_PATH = "result/images"

func (i *ImageHandler) Handle(response data.MainResponse) {
	childrenResponse := response.Data.Children[0].Data
	imageLink := childrenResponse.UrlDestination
	filename := path.Base(imageLink)

	filePath := fmt.Sprintf("%s/%s_%s_%s_%s", DEFAULT_IMAGE_DOWNLOAD_PATH, childrenResponse.Subreddit, childrenResponse.Title, childrenResponse.Author, filename)

	i.downloadFile(imageLink, filePath)

	semaphore.GetWaitGroup().Done()
}

func (i *ImageHandler) downloadFile(imageLink string, fileName string) {
	response, err := client.Get(imageLink)

	if nil != err {
		fmt.Println("Failed to retrieve image", imageLink)
		semaphore.GetWaitGroup().Done()

		return
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Failed to create file", fileName)
		semaphore.GetWaitGroup().Done()

		return
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Failed to write file", fileName)
		semaphore.GetWaitGroup().Done()

		return
	}
}
