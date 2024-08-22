package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	htmlReader := strings.NewReader(htmlBody)
	htmlNode, err := html.Parse(htmlReader)

	if err != nil {
		return []string{}, fmt.Errorf("error parsing htmlReader: %v", err)
	}

	var traverseNodes func(*html.Node)
	var urls []string
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
						continue
					}
					resolvedURL := baseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
					break
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			traverseNodes(c)
		}
	}
	traverseNodes(htmlNode)
	// fmt.Printf("%v", urls)

	return urls, nil
}
