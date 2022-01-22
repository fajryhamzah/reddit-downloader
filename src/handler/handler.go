package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fajryhamzah/reddit-downloader/src/data"
)

func Handle(links []string) {
	for _, link := range links {
		jsonLink := strings.TrimSuffix(link, "/") + ".json"

		go getResponse(jsonLink)
	}
}

func getResponse(link string) {
	req, err := http.NewRequest("GET", link, nil)

	if nil != err {
		fmt.Printf("Failed to create request %s \n", link)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")

	response, err := new(http.Client).Do(req)

	if nil != err {
		fmt.Printf("Failed to retrieve data from %s \n", link)
		fmt.Println("Detail :", err)
		return
	}

	defer response.Body.Close()

	var decodedResponse []data.MainResponse
	json.NewDecoder(response.Body).Decode(&decodedResponse)

	fmt.Println(response.StatusCode, decodedResponse[0].Data.Children[0].Data)
}
