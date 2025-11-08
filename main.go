package main

import (
	"fmt"
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

func main() {
	fmt.Println("Hello, World!")
}