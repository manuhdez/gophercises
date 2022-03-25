package link_parser

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(content string) []Link {
	r := strings.NewReader(content)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("Error parsing html: ", err)
		log.Fatal(err)
	}

	return traverse(doc)
}

func traverse(node *html.Node) []Link {
	var links []Link

	if node.Type == html.ElementNode && node.Data == "a" {
		l := node.FirstChild
		text := getLinkText(l)

		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, Link{Href: attr.Val, Text: text})
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, traverse(c)...)
	}

	return links
}

func getLinkText(n *html.Node) string {
	text := html.EscapeString(n.Data)

	next := n.NextSibling
	for {
		if next == nil {
			break
		}

		var extra string
		if next.FirstChild != nil && next.FirstChild.Type == html.TextNode {
			extra += html.EscapeString(next.FirstChild.Data)
		} else if next.Type == html.TextNode {
			extra += html.EscapeString(next.Data)
		}
		text += extra

		if next.NextSibling == nil {
			break
		}

		next = next.NextSibling
	}

	return cleanText(text)
}

func cleanText(t string) string {
	text := strings.Trim(t, "\n")
	return strings.TrimSpace(text)
}
