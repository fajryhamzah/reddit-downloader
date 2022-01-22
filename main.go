package main

import (
	"flag"
	"fmt"

	"github.com/fajryhamzah/reddit-downloader/src/handler"
)

func main() {
	flag.Parse()
	links := flag.Args()

	if len(links) < 1 {
		panic("Reddit Url Needed!")
	}

	handler.Handle(links)

	fmt.Scanln()
	fmt.Println("done")
}
