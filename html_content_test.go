package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "h1 present",
			input:    "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "h2 fallback when no h1",
			input:    "<html><body><h2>Secondary Title</h2></body></html>",
			expected: "Secondary Title",
		},
		{
			name:     "h1 takes priority over h2",
			input:    "<html><body><h1>Main</h1><h2>Sub</h2></body></html>",
			expected: "Main",
		},
		{
			name:     "no heading returns empty string",
			input:    "<html><body><p>Just a paragraph</p></body></html>",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "main tag priority",
			input:    `<html><body><p>Outside.</p><main><p>Main paragraph.</p></main></body></html>`,
			expected: "Main paragraph.",
		},
		{
			name:     "fallback to first p when no main",
			input:    "<html><body><p>First.</p><p>Second.</p></body></html>",
			expected: "First.",
		},
		{
			name:     "no paragraph returns empty string",
			input:    "<html><body><h1>Title only</h1></body></html>",
			expected: "",
		},
		{
			name:     "first p inside main when multiple exist",
			input:    `<html><body><main><p>Alpha.</p><p>Beta.</p></main></body></html>`,
			expected: "Alpha.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
