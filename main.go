package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseURL := args[0]
	fmt.Printf("starting crawl of: %s", baseURL)
	cfg, err := NewConfig(baseURL, 5)
	if err != nil {
		log.Fatal("Error creating config struct")
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	fmt.Println("\n--- Crawl Results ---")
	for url, count := range cfg.pages {
		fmt.Printf("- %s: %d\n", url, count)
	}
}
