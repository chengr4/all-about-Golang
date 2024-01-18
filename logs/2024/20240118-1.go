package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)
	lists := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.yahoo.com",
	}

	for _, url := range lists {
		go get(url, results)
	}
	println("break point one")

	// it will read until the channel is closed
	for range lists {
		res := <-results
		if res.err != nil {
			// handle error
			log.Printf("%s %s\n", res.err, res.url)
		} else {
			// handle result
			log.Printf("%s %s\n", res.latency, res.url)
		}
	}
	println("break point two")
}
