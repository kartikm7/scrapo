package scraper

import (
	"net/http"

	"golang.org/x/net/html"
)

func Scraper(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
