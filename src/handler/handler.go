package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fajryhamzah/reddit-downloader/src/client"
	"github.com/fajryhamzah/reddit-downloader/src/data"
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
		fmt.Printf("Failed to retrieve data from %s \n", link)
		fmt.Println("Detail :", err)
		semaphore.GetWaitGroup().Done()
		return
	}

	defer response.Body.Close()

	var decodedResponse []data.MainResponse
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	fmt.Println(response.StatusCode, decodedResponse[0].Data.Children[0].Data)
	semaphore.GetWaitGroup().Done()
}
