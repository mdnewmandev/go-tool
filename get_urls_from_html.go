package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	returnedURLs := []string{}
	
	if len(htmlBody) == 0 {
		return returnedURLs, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return returnedURLs, err
	}
	
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			parsedURL, err := url.Parse(href)
			if err == nil {
				resolvedURL := baseURL.ResolveReference(parsedURL)
				returnedURLs = append(returnedURLs, resolvedURL.String())
			}
		}
	})

	return returnedURLs, nil
}