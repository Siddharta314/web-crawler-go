package main

import "net/url"

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
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
