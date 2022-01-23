package media

import (
	"errors"

	"github.com/fajryhamzah/reddit-downloader/src/data"
)

var imageHandler MediaHandlerInterface = &ImageHandler{}
var videoHandler MediaHandlerInterface = &VideoHandler{}

func GetHandler(response data.MainResponse) (MediaHandlerInterface, error) {
	childrenData := response.Data.Children[0].Data
	var handler MediaHandlerInterface

	if childrenData.IsVideo {
		handler = videoHandler
	} else if childrenData.UrlDestination != "" {
		handler = imageHandler
	}

	if nil != handler {
		if !handler.IsSupported() {
			return nil, errors.New(handler.NotSupportedMessage())
		}

		return handler, nil
	}

	return nil, errors.New("Media is not supported for " + childrenData.Permalink)
}
