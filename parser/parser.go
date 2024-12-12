package parser

import (
	"golang.org/x/net/html"
)

func Parser(node *html.Node) (cleaned string, err error) {
	if node.Type == html.ElementNode {
		if node.Data == "h1" || node.Data == "p" {
			if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
				cleaned += node.FirstChild.Data
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		returned, _ := Parser(c)
		cleaned += returned
	}
	return cleaned, err
}
