package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type results struct {
	URL     string
	Err     error
	Latency time.Duration
}

func get(url string) results {
	now := time.Now()
	resp, err := http.Get(url)
	latency := time.Since(now)
	if err != nil {
		return results{url, err, latency}
	}
	defer resp.Body.Close()
	return results{url, nil, latency}
}

func main() {
	urls := []string{
		"https://amazon.com",
		"https://apple.com",
		"https://example.com",
		"https://google.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			result := get(url)
			if result.Err != nil {
				fmt.Printf("Error: %v - %s\n", result.Err, result.URL)
			} else {
				fmt.Printf("%v: %s\n", result.Latency.Round(time.Millisecond), result.URL)
			}
		}(url)
	}

	wg.Wait()
}
