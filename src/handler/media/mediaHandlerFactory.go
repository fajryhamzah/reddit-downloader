package media

import (
	"errors"

	"github.com/fajryhamzah/reddit-downloader/src/data"
)

var imageHandler ImageHandler

func GetHandler(response data.MainResponse) (MediaHandlerInterface, error) {
	childrenData := response.Data.Children[0].Data
	if !childrenData.IsVideo && childrenData.UrlDestination != "" {
		return &imageHandler, nil
	}

	return nil, errors.New("Media is not supported for " + childrenData.Permalink)
}
