package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}
	normalized := parsed.Host + parsed.Path
	normalized = strings.TrimRight(normalized, "/")
	return normalized, nil
}
