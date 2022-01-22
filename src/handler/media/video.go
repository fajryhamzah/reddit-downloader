package media

import (
	"fmt"

	"github.com/fajryhamzah/reddit-downloader/src/client"
	"github.com/fajryhamzah/reddit-downloader/src/data"
	"github.com/fajryhamzah/reddit-downloader/src/log"
	"github.com/fajryhamzah/reddit-downloader/src/semaphore"
	"github.com/fajryhamzah/reddit-downloader/src/writer"
)

type VideoHandler struct{}

const VIDEOS_PATH string = "result/videos"

func (v *VideoHandler) Handle(response data.MainResponse) {
	childrenResponse := response.Data.Children[0].Data

	if childrenResponse.SecureMedia.RedditVideo.IsGIF {
		v.downloadGIF(childrenResponse)
		return
	}

	semaphore.GetWaitGroup().Done()
}

func (v *VideoHandler) downloadGIF(response data.DataResponse) {
	resp, err := client.Get(response.SecureMedia.RedditVideo.FallbackUrl)

	if nil != err {
		log.Error("Failed to retrieve GIF", response.SecureMedia.RedditVideo.FallbackUrl)
		semaphore.GetWaitGroup().Done()

		return
	}

	filePath := fmt.Sprintf("%s/%s_%s_%s.mp4", VIDEOS_PATH, response.Subreddit, response.Title, response.Author)

	log.Logf("Downloading GIF from %s/%s", response.SubredditPrefix, response.Title)
	writer.Write(filePath, resp.Body)

	log.Successf("GIF %s downloaded.", filePath)
	semaphore.GetWaitGroup().Done()
}
