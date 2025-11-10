package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) { 
	if rawCurrentURL == "" {
		rawCurrentURL = rawBaseURL
	}

	rawCurrentURL = strings.TrimSuffix(rawCurrentURL, "/")
	
	if _, exists := pages[rawCurrentURL]; exists {
		pages[rawCurrentURL]++
		return
	} else {
		pages[rawCurrentURL] = 1
	}

	html, err := getHTML(rawCurrentURL)

	if err != nil {
		fmt.Printf("error fetching HTML: %v", err)
		fmt.Println()
	}
	
	pageData := extractPageData(html, rawCurrentURL)

	fmt.Printf("Crawled URL: %s\n", pageData.URL)
	fmt.Printf("Number of Outgoing Links: %d\n", len(pageData.OutgoingLinks))
	fmt.Println("================================")

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %v", err)
		os.Exit(1)
	}

	for _, link := range pageData.OutgoingLinks {
		parsedLink, err := url.Parse(link)
		if err != nil {
			continue
		}
		if parsedLink.Host == base.Host {
			crawlPage(rawBaseURL, link, pages)
		}
	}
}