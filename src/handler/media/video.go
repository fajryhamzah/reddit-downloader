package media

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fajryhamzah/reddit-downloader/src/client"
	"github.com/fajryhamzah/reddit-downloader/src/data"
	"github.com/fajryhamzah/reddit-downloader/src/log"
	"github.com/fajryhamzah/reddit-downloader/src/semaphore"
	"github.com/fajryhamzah/reddit-downloader/src/writer"
)

type VideoHandler struct{}

const VIDEOS_PATH string = "result/videos"

const TEMP_HLS_PATH string = "temp"

func (v *VideoHandler) Handle(response data.MainResponse) {
	childrenResponse := response.Data.Children[0].Data

	if childrenResponse.SecureMedia.RedditVideo.IsGIF {
		v.downloadGIF(childrenResponse)
		return
	}

	v.downloadVideo(childrenResponse)
}

func (v *VideoHandler) IsSupported() bool {
	cmd := exec.Command("/bin/sh", "-c", "command", "-v", "ffmpeg")

	if err := cmd.Run(); err != nil {
		log.Error(err)
		return false
	}

	return true
}

func (v *VideoHandler) NotSupportedMessage() string {
	return "Video download disable. Please install ffmpeg."
}

func (v *VideoHandler) downloadGIF(response data.DataResponse) {
	resp, err := client.Get(response.SecureMedia.RedditVideo.FallbackUrl)

	if nil != err {
		log.Error("Failed to retrieve GIF", response.SecureMedia.RedditVideo.FallbackUrl)
		semaphore.GetWaitGroup().Done()

		return
	}

	defer resp.Body.Close()

	filePath := fmt.Sprintf("%s/%s_%s_%s.mp4", VIDEOS_PATH, response.Subreddit, response.Title, response.Author)

	log.Logf("Downloading GIF from %s/%s", response.SubredditPrefix, response.Title)
	err = writer.Write(filePath, resp.Body)

	if nil != err {
		log.Errorf("Failed to write file to %s", filePath)
		semaphore.GetWaitGroup().Done()

		return
	}

	log.Successf("GIF %s downloaded.", filePath)
	semaphore.GetWaitGroup().Done()
}

func (v *VideoHandler) downloadVideo(response data.DataResponse) {
	err := os.MkdirAll(TEMP_HLS_PATH, os.ModePerm)

	if nil != err {
		log.Error("Failed to create temp file for video downloader")
		semaphore.GetWaitGroup().Done()

		return
	}

	semaphore.GetWaitGroup().Done()
}
