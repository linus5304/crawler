package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got Network error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %v", res.Status)
	}
	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", res.Header.Get("Content-Type"))
	}
	htmlBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	return string(htmlBody), nil
}
