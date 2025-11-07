package main

func normalizeURL(url string) (string, error) {
	// Remove scheme (http:// or https://)
	if len(url) >= 8 && (url[:8] == "https://" || url[:7] == "http://") {
		if url[:8] == "https://" {
			url = url[8:]
		} else {
			url = url[7:]
		}
	}

	return url, nil
}