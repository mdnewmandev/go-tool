package main

import (
	"net/url"
	"testing"
)

func TestGetH1fromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "empty string",
			html:     "",
			expected: "",
		},
		{
			name:     "simple h1",
			html:     "<h1>Welcome to My Site</h1>",
			expected: "Welcome to My Site",
		},
		{
			name:     "h1 with other tags",
			html:     "<html><body><h1>Test title</h1></body></html>",
			expected: "Test title",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1fromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "empty string",
			html:     "",
			expected: "",
		},
		{
			name:     "simple p",
			html:     "<p>This is the first paragraph.</p><p>This is the second paragraph.</p>",
			expected: "This is the first paragraph.",
		},
		{
			name:     "main with p",
			html:     "<p>Outside paragraph.</p><main><p>Main paragraph.</p></main>",
			expected: "Main paragraph.",
		},
		{
			name:     "no p tags",
			html:     "<div>No paragraphs here.</div>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTMLMainPriority(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		baseURL  string
		expected []string
	}{
		{
			name:     "empty string",
			html:     "",
			baseURL:  "http://example.com",
			expected: []string{},
		},
		{
			name:     "simple links",
			html:     `<a href="page1.html">Page 1</a><a href="/page2.html">Page 2</a>`,
			baseURL:  "http://example.com",
			expected: []string{"http://example.com/page1.html", "http://example.com/page2.html"},
		},
		{
			name:     "absolute and relative links",
			html:     `<a href="http://other.com/page3.html">Page 3</a><a href="page4.html">Page 4</a>`,
			baseURL:  "http://example.com",
			expected: []string{"http://other.com/page3.html", "http://example.com/page4.html"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			base, err := url.Parse(tc.baseURL)
			if err != nil {
				t.Fatalf("Failed to parse base URL: %v", err)
			}
			actual, err := getURLsFromHTML(tc.html, base)
			if err != nil {
				t.Fatalf("getURLsFromHTML returned error: %v", err)
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected length: '%v', actual length: '%v'", i, tc.name, len(tc.expected), len(actual))
				return
			}
			for j := range actual {
				if actual[j] != tc.expected[j] {
					t.Errorf("Test %v - %s FAIL: expected[%d]: '%v', actual[%d]: '%v'", i, tc.name, j, tc.expected[j], j, actual[j])
				}
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		baseURL  string
		expected []string
	}{
		{
			name:     "empty string",
			html:     "",
			baseURL:  "http://example.com",
			expected: []string{},
		},
		{
			name:     "simple images",
			html:     `<img src="image1.jpg"/><img src="/image2.png"/>`,
			baseURL:  "http://example.com",
			expected: []string{"http://example.com/image1.jpg", "http://example.com/image2.png"},
		},
		{
			name:     "absolute and relative images",
			html:     `<img src="http://other.com/image3.gif"/><img src="image4.svg"/>`,
			baseURL:  "http://example.com",
			expected: []string{"http://other.com/image3.gif", "http://example.com/image4.svg"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			base, err := url.Parse(tc.baseURL)
			if err != nil {
				t.Fatalf("Failed to parse base URL: %v", err)
			}
			actual, err := getImagesFromHTML(tc.html, base)
			if err != nil {
				t.Fatalf("getImagesFromHTML returned error: %v", err)
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected length: '%v', actual length: '%v'", i, tc.name, len(tc.expected), len(actual))
				return
			}
			for j := range actual {
				if actual[j] != tc.expected[j] {
					t.Errorf("Test %v - %s FAIL: expected[%d]: '%v', actual[%d]: '%v'", i, tc.name, j, tc.expected[j], j, actual[j])
				}
			}
		})
	}
}

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		pageURL  string
		expected PageData
	}{
		{
			name:    "basic page",
			html:    `<html><body><h1>Title</h1><main><p>Main paragraph.</p></main><a href="page1.html">Link</a><img src="image1.jpg"/></body></html>`,
			pageURL: "http://example.com",
			expected: PageData{
				URL:            "http://example.com",
				H1:             "Title",
				FirstParagraph: "Main paragraph.",
				OutgoingLinks:  []string{"http://example.com/page1.html"},
				ImageURLs:      []string{"http://example.com/image1.jpg"},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.html, tc.pageURL)
			if actual.URL != tc.expected.URL ||
				actual.H1 != tc.expected.H1 ||
				actual.FirstParagraph != tc.expected.FirstParagraph ||
				len(actual.OutgoingLinks) != len(tc.expected.OutgoingLinks) ||
				len(actual.ImageURLs) != len(tc.expected.ImageURLs) {
				t.Errorf("Test %v - %s FAIL: expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
				return
			}
			for j := range actual.OutgoingLinks {
				if actual.OutgoingLinks[j] != tc.expected.OutgoingLinks[j] {
					t.Errorf("Test %v - %s FAIL: expected OutgoingLinks[%d]: '%v', actual[%d]: '%v'", i, tc.name, j, tc.expected.OutgoingLinks[j], j, actual.OutgoingLinks[j])
				}
			}
			for j := range actual.ImageURLs {
				if actual.ImageURLs[j] != tc.expected.ImageURLs[j] {
					t.Errorf("Test %v - %s FAIL: expected ImageURLs[%d]: '%v', actual[%d]: '%v'", i, tc.name, j, tc.expected.ImageURLs[j], j, actual.ImageURLs[j])
				}
			}
		})
	}
}