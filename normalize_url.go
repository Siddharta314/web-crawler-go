package main

import (
	"net/url"
	"strings"
)

func normalizeURL(url_string string) (string, error) {
	parsed, err := url.Parse(url_string)
	if err != nil {
		return "", err
	}

	fullPath := parsed.Host + parsed.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
