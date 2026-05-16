package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	if h1 := doc.Find("h1").First().Text(); h1 != "" {
		return h1
	}
	return doc.Find("h2").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	if main := doc.Find("main"); main.Length() > 0 {
		return main.Find("p").First().Text()
	}
	return doc.Find("p").First().Text()
}
