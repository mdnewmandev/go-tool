package main

import (
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