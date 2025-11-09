package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1fromHTML(html string) string {
	headerContent := ""
	
	if len(html) == 0 {
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		headerContent = s.Text()
	})

	return headerContent
}

func getFirstParagraphFromHTMLMainPriority(html string) string {
	paragraphContent := ""
	
	if len(html) == 0 {
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	
	doc.Find("main p").EachWithBreak(func(i int, s *goquery.Selection) bool {
		paragraphContent = s.Text()
		return false // Break after the first match
	})

	if paragraphContent != "" {
		return paragraphContent
	}

	doc.Find("p").EachWithBreak(func(i int, s *goquery.Selection) bool {
		paragraphContent = s.Text()
		return false // Break after the first match
	})

	return paragraphContent
}

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

func main() {
	fmt.Println("Hello, World!")
}