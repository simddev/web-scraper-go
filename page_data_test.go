package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		inputBody string
		expected PageData
	}{
		{
			name:     "full page with all fields",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
				<h1>Test Title</h1>
				<p>This is the first paragraph.</p>
				<a href="/link1">Link 1</a>
				<img src="/image1.jpg" alt="Image 1">
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://crawler-test.com/link1"},
				ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
			},
		},
		{
			name:     "page with no links or images",
			inputURL: "https://crawler-test.com/about",
			inputBody: `<html><body>
				<h2>About Us</h2>
				<p>We are a company.</p>
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com/about",
				Heading:        "About Us",
				FirstParagraph: "We are a company.",
				OutgoingLinks:  nil,
				ImageURLs:      nil,
			},
		},
		{
			name:     "page with multiple links and images",
			inputURL: "https://crawler-test.com/blog",
			inputBody: `<html><body>
				<h1>Blog</h1>
				<p>Welcome to the blog.</p>
				<a href="/post1">Post 1</a>
				<a href="/post2">Post 2</a>
				<img src="/thumb1.jpg">
				<img src="/thumb2.jpg">
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com/blog",
				Heading:        "Blog",
				FirstParagraph: "Welcome to the blog.",
				OutgoingLinks:  []string{"https://crawler-test.com/post1", "https://crawler-test.com/post2"},
				ImageURLs:      []string{"https://crawler-test.com/thumb1.jpg", "https://crawler-test.com/thumb2.jpg"},
			},
		},
		{
			name:     "empty page",
			inputURL: "https://crawler-test.com/empty",
			inputBody: `<html><body></body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com/empty",
				Heading:        "",
				FirstParagraph: "",
				OutgoingLinks:  nil,
				ImageURLs:      nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputBody, tc.inputURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %+v, got %+v", tc.expected, actual)
			}
		})
	}
}
