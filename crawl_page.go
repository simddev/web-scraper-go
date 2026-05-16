package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, exists := cfg.pages[normalizedURL]; exists {
		return false
	}
	cfg.pages[normalizedURL] = PageData{}
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl }()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}
	if !cfg.addPageVisit(normalizedCurrent) {
		return
	}

	fmt.Printf("crawling: %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching %s: %v\n", rawCurrentURL, err)
		return
	}

	pageData := extractPageData(html, rawCurrentURL)

	cfg.mu.Lock()
	cfg.pages[normalizedCurrent] = pageData
	cfg.mu.Unlock()

	links, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("error extracting URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}
