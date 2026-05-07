package main

import (
	"net/url"
	"reflect"
	"testing"
)

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

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "convert relative to absolute",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one"><span>Link 1</span></a>
		<a href="https://other.com/path/one"><span>Link 2</span></a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "multiple links and nested elements",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<div>
			<a href="/first">First</a>
			<section>
				<a href="/second">Second</a>
				<a href="https://google.com">Google</a>
			</section>
		</div>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/first", "https://blog.boot.dev/second", "https://google.com"},
		},
		{
			name:     "ignore empty and missing href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="">Empty</a>
		<a>No Href</a>
		<a href="/valid">Valid</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/valid"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' SETUP FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, actual %v", i, tc.name, tc.expected, actual)
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
			name:     "absolute and relative image sources",
			inputURL: "https://boot.dev",
			inputBody: `
<html>
	<body>
		<img src="/assets/logo.png" alt="Logo">
		<img src="https://external.com/dog.jpg">
	</body>
</html>
`,
			expected: []string{"https://boot.dev/assets/logo.png", "https://external.com/dog.jpg"},
		},
		{
			name:     "missing src attribute",
			inputURL: "https://boot.dev",
			inputBody: `
<html>
	<body>
		<img alt="No src here">
		<img src="/valid-image.png">
	</body>
</html>
`,
			expected: []string{"https://boot.dev/valid-image.png"},
		},
		{
			name:     "images inside different containers",
			inputURL: "https://boot.dev",
			inputBody: `
<html>
	<body>
		<header>
			<img src="/icon.svg">
		</header>
		<main>
			<div>
				<img src="https://cdn.com/banner.png">
			</div>
		</main>
	</body>
</html>
`,
			expected: []string{"https://boot.dev/icon.svg", "https://cdn.com/banner.png"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' SETUP FAIL: %v", i, tc.name, err)
				return
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
