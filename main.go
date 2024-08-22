package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		return
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := args[0]
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	pages := make(map[string]int)
	crawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
