package handler

import (
	"encoding/json"
	"strings"

	"github.com/fajryhamzah/reddit-downloader/src/client"
	"github.com/fajryhamzah/reddit-downloader/src/data"
	"github.com/fajryhamzah/reddit-downloader/src/handler/media"
	"github.com/fajryhamzah/reddit-downloader/src/log"
	"github.com/fajryhamzah/reddit-downloader/src/semaphore"
)

func Handle(links []string) {
	for _, link := range links {
		jsonLink := strings.TrimSuffix(link, "/") + ".json"

		semaphore.GetWaitGroup().Add(1)
		go getResponse(jsonLink)
	}
}

func getResponse(link string) {
	response, err := client.Get(link)

	if nil != err || response.StatusCode != 200 {
		log.Errorf("Failed to retrieve data from %s", link)
		log.Error("Detail :", err)
		semaphore.GetWaitGroup().Done()
		return
	}

	defer response.Body.Close()

	var decodedResponse []data.MainResponse
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	var handler media.MediaHandlerInterface
	handler, err = media.GetHandler(decodedResponse[0])

	if nil != err {
		log.Error(err)
		semaphore.GetWaitGroup().Done()
		return
	}

	go handler.Handle(decodedResponse[0])
}
