package client

import (
	"errors"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	if nil != err {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")

	response, err := new(http.Client).Do(req)

	if nil != err {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Got status code " + response.Status)
	}

	return response, nil
}
