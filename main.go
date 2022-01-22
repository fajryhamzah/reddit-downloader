package main

import (
	"flag"

	"github.com/fajryhamzah/reddit-downloader/src/handler"
	"github.com/fajryhamzah/reddit-downloader/src/log"
	"github.com/fajryhamzah/reddit-downloader/src/semaphore"
)

func main() {
	flag.Parse()
	links := flag.Args()

	if len(links) < 1 {
		panic("Reddit Url Needed!")
	}

	handler.Handle(links)

	semaphore.GetWaitGroup().Wait()
	log.Success("DONE")
}
