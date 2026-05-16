package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if _, seen := pages[normalizedCurrent]; seen {
		pages[normalizedCurrent]++
		return
	}

	pages[normalizedCurrent] = 1
	fmt.Printf("crawling: %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching %s: %v\n", rawCurrentURL, err)
		return
	}

	links, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("error extracting URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, link := range links {
		crawlPage(rawBaseURL, link, pages)
	}
}
