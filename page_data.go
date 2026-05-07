package main

import "net/url"

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)
	paragraph := getFirstParagraphFromHTML(html)
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			Heading:        heading,
			FirstParagraph: paragraph,
			OutgoingLinks:  []string{},
			ImageURLs:      []string{},
		}
	}
	links, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		links = []string{}
	}
	images, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		images = []string{}
	}
	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: paragraph,
		OutgoingLinks:  links,
		ImageURLs:      images,
	}
}
