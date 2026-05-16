package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)

	const maxConcurrency = 5
	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())
	cfg.wg.Wait()

	fmt.Println("\n--- crawl results ---")
	for page, data := range cfg.pages {
		fmt.Printf("%s (heading: %q)\n", page, data.Heading)
	}
}
