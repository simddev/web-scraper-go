package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("usage: crawler <url> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("maxConcurrency must be a positive integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil || maxPages < 1 {
		fmt.Println("maxPages must be a positive integer")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)

	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())
	cfg.wg.Wait()

	fmt.Println("\n--- crawl results ---")
	for page, data := range cfg.pages {
		fmt.Printf("%s (heading: %q)\n", page, data.Heading)
	}

	if err := writeJSONReport(cfg.pages, "report.json"); err != nil {
		fmt.Printf("error writing JSON report: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("report written to report.json")
}
