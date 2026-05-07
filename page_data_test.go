package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		html     string
		expected PageData
	}{
		{
			name:     "full page extraction",
			inputURL: "https://boot.dev",
			html: `
<html>
	<body>
		<h1>Main Title</h1>
		<main><p>Core content paragraph.</p></main>
		<a href="/about">About</a>
		<img src="/logo.png">
	</body>
</html>`,
			expected: PageData{
				URL:            "https://boot.dev",
				Heading:        "Main Title",
				FirstParagraph: "Core content paragraph.",
				OutgoingLinks:  []string{"https://boot.dev/about"},
				ImageURLs:      []string{"https://boot.dev/logo.png"},
			},
		},
		{
			name:     "extraction with fallbacks (h2 and no main p)",
			inputURL: "https://example.com",
			html: `
<html>
	<body>
		<h2>Secondary Heading</h2>
		<p>General paragraph outside main.</p>
		<a href="https://google.com">External</a>
	</body>
</html>`,
			expected: PageData{
				URL:            "https://example.com",
				Heading:        "Secondary Heading",
				FirstParagraph: "General paragraph outside main.",
				OutgoingLinks:  []string{"https://google.com"},
				ImageURLs:      nil, // o []string{} según tu implementación
			},
		},
		{
			name:     "empty or minimal html",
			inputURL: "https://empty.com",
			html:     `<html><body></body></html>`,
			expected: PageData{
				URL:            "https://empty.com",
				Heading:        "",
				FirstParagraph: "",
				OutgoingLinks:  nil,
				ImageURLs:      nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.html, tc.inputURL)

			// Usamos reflect.DeepEqual para comparar los slices y el struct completo
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test '%s' FAIL:\nExpected: %+v\nActual:   %+v", tc.name, tc.expected, actual)
			}
		})
	}
}
