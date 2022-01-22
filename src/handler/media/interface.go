package media

import "github.com/fajryhamzah/reddit-downloader/src/data"

type MediaHandlerInterface interface {
	Handle(response data.MainResponse)
}
