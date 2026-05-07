package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	heading := strings.TrimSpace(doc.Find("h1").First().Text())
	if heading != "" {
		return heading
	}

	return strings.TrimSpace(doc.Find("h2").First().Text())
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	p := strings.TrimSpace(doc.Find("main p").First().Text())
	if p != "" {
		return p
	}

	return strings.TrimSpace(doc.Find("p").First().Text())
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if href == "" {
			return
		}
		resolvedURL, err := baseURL.Parse(href)
		if err != nil {
			return
		}

		urls = append(urls, resolvedURL.String())
	})
	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if src == "" {
			return
		}
		resolvedURL, err := baseURL.Parse(src)
		if err != nil {
			return
		}

		urls = append(urls, resolvedURL.String())
	})
	return urls, nil
}
