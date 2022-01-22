package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

type mainResponse struct {
	Data childrenResponse `json:"data"`
}

type childrenResponse struct {
	Children []struct {
		Data mainDataResponse
	} `json:"children"`
}

type mainDataResponse struct {
	Subreddit       string `json:"subreddit"`
	Title           string `json:"title"`
	SubredditPrefix string `json:"subreddit_name_prefixed"`
	Author          string `json:"author"`
	IsVideo         bool   `json:"is_video"`
	UrlDestination  string `json:"url_overridden_by_dest"`
}

func main() {
	flag.Parse()
	links := flag.Args()

	for _, link := range links {
		jsonLink := strings.TrimSuffix(link, "/") + ".json"

		req, err := http.NewRequest("GET", jsonLink, nil)

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

		var decodedResponse []mainResponse
		json.NewDecoder(response.Body).Decode(&decodedResponse)

		fmt.Println(response.StatusCode, decodedResponse[0].Data.Children[0].Data)
	}
}
