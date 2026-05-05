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
