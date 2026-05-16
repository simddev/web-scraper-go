package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		resolved, err := baseURL.Parse(href)
		if err == nil {
			urls = append(urls, resolved.String())
		}
	})
	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	var images []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		resolved, err := baseURL.Parse(src)
		if err == nil {
			images = append(images, resolved.String())
		}
	})
	return images, nil
}
