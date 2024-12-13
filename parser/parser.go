package parser

import (
	"fmt"

	"golang.org/x/net/html"
)

func Parser(node *html.Node) (cleaned string, err error) {
	// processing
	if node.Type == html.ElementNode {
		// here the data is for sure the name of the tag
		cleaned += getText(node, node.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		// making the recursive call to ensure we traverse the tree
		response, _ := Parser(c)
		cleaned += response
	}

	return cleaned, err
}

// TODO: I think the function here only handles anchor tags that are the parent node,
// and does not handle cases where the anchor tag is a child within something like a list (li tag)
// function to get text element
func getText(node *html.Node, elementType string) (text string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		// to handle anchor tags
		if elementType == "a" {
			for _, attribute := range node.Attr {
				if attribute.Key == "href" {
					var attr string
					if attribute.Val != "undefined" {
						attr = attribute.Val
						text = fmt.Sprintf("[%v](%v)", text, attr)
					}
				}
			}
		}
		if c.Type == html.TextNode {
			text += html2md(c, elementType)
		}
	}
	return text
}

// function to parse the HTML to Markdown
func html2md(node *html.Node, elementType string) (md string) {
	text := node.Data
	switch elementType {
	case "h1":
		md = fmt.Sprintf("#%v\n", text)
	case "h2":
		md = fmt.Sprintf("##%v\n", text)
	case "h3":
		md = fmt.Sprintf("###%v\n", text)
	case "h4":
		md = fmt.Sprintf("####%v\n", text)
	case "h5":
		md = fmt.Sprintf("#####%v\n", text)
	case "h6":
		md = fmt.Sprintf("######%v\n", text)
	case "li":
		md = "- " + text
	case "p":
		md = text
	default:
		md = ""
	}
	return md
}
