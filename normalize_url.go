package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	
	normalizedURL := u.Host + strings.TrimSuffix(u.Path, "/")
	fmt.Println(normalizedURL)
	
	return normalizedURL, nil
}