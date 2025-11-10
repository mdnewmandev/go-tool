package main

import (
	"fmt"
	"os"
	"sort"
)

func printPages(pages map[string]int) {
	if len(pages) == 0 {
		fmt.Println("(no pages)")
		return
	}

	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, pages[k])
	}
}

func main() {
	rawBaseURL := os.Args[1]
	
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(os.Args[1:]) == 1 {
		fmt.Printf("starting crawl of: %s", rawBaseURL)
		fmt.Println()
	}

	pages := make(map[string]int)
	crawlPage(rawBaseURL, "", pages)

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("Crawl completed successfully.")
	printPages(pages)
	fmt.Println()
}