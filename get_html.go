package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
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