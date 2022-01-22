package media

import (
	"errors"

	"github.com/fajryhamzah/reddit-downloader/src/data"
)

var imageHandler ImageHandler

func GetHandler(response data.MainResponse) (MediaHandlerInterface, error) {
	if !response.Data.Children[0].Data.IsVideo {
		return &imageHandler, nil
	}

	return nil, errors.New("Media is not supported")
}
