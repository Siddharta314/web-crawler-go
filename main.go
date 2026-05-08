package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("Error converting maxConcurrency to int")
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("Error converting maxPages to int")
	}
	fmt.Printf("starting crawl of: %s", baseURL)
	cfg, err := NewConfig(baseURL, maxConcurrency, maxPages)
	if err != nil {
		log.Fatal("Error creating config struct")
	}

	cfg.wg.Add(1)
	startTime := time.Now()
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()
	fmt.Printf("Tiempo total: %s\n", time.Since(startTime))

	fmt.Printf("Pages crawled: %d\n", len(cfg.pages))
	fmt.Println("\n--- Crawl Results ---")
	for url, count := range cfg.pages {
		fmt.Printf("- %s: %d\n", url, count)
	}
}
