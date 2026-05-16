package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		inputBody string
		expected []string
	}{
		{
			name:      "absolute URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{
			name:      "relative URL converted to absolute",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/blog/post">Blog Post</a></body></html>`,
			expected:  []string{"https://crawler-test.com/blog/post"},
		},
		{
			name:      "multiple links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/about">About</a><a href="/contact">Contact</a></body></html>`,
			expected:  []string{"https://crawler-test.com/about", "https://crawler-test.com/contact"},
		},
		{
			name:      "no links returns nil",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>No links here</p></body></html>`,
			expected:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("couldn't parse input URL: %v", err)
			}
			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "relative image URL converted to absolute",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "absolute image URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="https://cdn.example.com/photo.jpg" alt="Photo"></body></html>`,
			expected:  []string{"https://cdn.example.com/photo.jpg"},
		},
		{
			name:      "multiple images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/a.png"><img src="/b.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/a.png", "https://crawler-test.com/b.png"},
		},
		{
			name:      "img without src is skipped",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img alt="no src"><img src="/valid.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/valid.png"},
		},
		{
			name:      "no images returns nil",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>No images here</p></body></html>`,
			expected:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("couldn't parse input URL: %v", err)
			}
			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
