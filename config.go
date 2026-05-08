package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func NewConfig(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	return &config{
		pages:              make(map[string]PageData),
		baseURL:            base,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}
