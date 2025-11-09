package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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

type PageData struct {
	URL           	string
	H1            	string
	FirstParagraph string
	OutgoingLinks 	[]string
	ImageURLs     	[]string
}

func extractPageData(html, pageURL string) PageData {
	base, err := url.Parse(pageURL)
	if err != nil {
		base = &url.URL{}
	}

	return PageData{
		URL:	   		pageURL,
		H1:        		getH1fromHTML(html),
		FirstParagraph: getFirstParagraphFromHTMLMainPriority(html),
		OutgoingLinks:  func() []string { urls, _ := getURLsFromHTML(html, base); return urls }(),
		ImageURLs: 		func() []string { imgs, _ := getImagesFromHTML(html, base); return imgs }(),
	}
}

func getHTML(rawURL string) (string, error) {
	// Use http.NewRequest with an http.Client to fetch the webpage of the rawURL. Set a User-Agent header (e.g. BootCrawler/1.0) to avoid being blocked by servers.
	client := &http.Client{}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return "", fmt.Errorf("failed to fetch URL: %s, status code: %d", rawURL, resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "" && !strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("content type is not text/html: %s", resp.Header.Get("Content-Type"))
	}
	if resp.ContentLength == 0 {
		return "", fmt.Errorf("content length is zero")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(os.Args[1:]) == 1 {
		fmt.Printf("starting crawl of: %s", os.Args[1])
		fmt.Println()
	}
	html, err := getHTML(os.Args[1])
	if err != nil {
		fmt.Printf("error fetching HTML: %v", err)
		os.Exit(1)
	}
	// pageData := extractPageData(html, os.Args[1])
	// fmt.Printf("Crawled Page Data: %+v\n", pageData)
	fmt.Println("Crawl completed successfully.")
	fmt.Printf("Fetched HTML: %d\n", html)
}