package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	returnedImages := []string{}
	
	if len(htmlBody) == 0 {
		return returnedImages, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return returnedImages, err
	}
	
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			parsedURL, err := url.Parse(src)
			if err == nil {
				resolvedURL := baseURL.ResolveReference(parsedURL)
				returnedImages = append(returnedImages, resolvedURL.String())
			}
		}
	})

	return returnedImages, nil
}