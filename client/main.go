package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	serverAddr := "http://localhost:8080/"
	requestsPerSecond := 10
	// useNoDelay := os.Getenv("NO_DELAY") == "true"

	client := &http.Client{}

	for i := 0; i < requestsPerSecond; i++ {
		go func() {
			start := time.Now() // Start timing

			resp, err := client.Get(serverAddr)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			elapsed := time.Since(start) // Measure latency
			fmt.Printf("Response received: %s, Latency: %s\n", resp.Status, elapsed)
		}()
		time.Sleep(1 * time.Second / time.Duration(requestsPerSecond))
	}
}
