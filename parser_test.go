package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "priority h1",
			input:    `<html><body><h1>Main Title</h1><h2>Subtitle</h2></body></html>`,
			expected: "Main Title",
		},
		{
			name:     "fallback to h2",
			input:    `<html><body><h2>Only Subtitle</h2><p>Content</p></body></html>`,
			expected: "Only Subtitle",
		},
		{
			name:     "no headings found",
			input:    `<html><body><p>Just a paragraph</p></body></html>`,
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("Test '%s' FAIL: expected '%s', actual '%s'", tc.name, tc.expected, actual)
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
			name:     "prefer p inside main",
			input:    `<html><body><p>Outside</p><main><p>Inside Main</p></main></body></html>`,
			expected: "Inside Main",
		},
		{
			name:     "fallback to any p",
			input:    `<html><body><div><p>General Paragraph</p></div></body></html>`,
			expected: "General Paragraph",
		},
		{
			name:     "no paragraph found",
			input:    `<html><body><header>No content here</header></body></html>`,
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("Test '%s' FAIL: expected '%s', actual '%s'", tc.name, tc.expected, actual)
			}
		})
	}
}
