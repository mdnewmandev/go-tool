package main

import "net/url"

type PageData struct {
	URL           	string
	H1            	string
	FirstParagraph  string
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