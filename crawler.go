package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error status code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if baseURL.Host != currentURL.Host {
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if count, ok := pages[normalizedURL]; ok {
		pages[normalizedURL] = count + 1
		return
	}

	pages[normalizedURL] = 1

	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("Error parsing URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, nextURL := range urls {
		crawlPage(rawBaseURL, nextURL, pages)
	}
}
