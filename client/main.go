package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	serverAddr := "http://localhost:8080/"
	client := &http.Client{}

	// to save latency logs
	fileIndex := 1
	for {
		if _, err := os.Stat(fmt.Sprintf("log%d.csv", fileIndex)); os.IsNotExist(err) {
			break
		}
		fileIndex++
	}

	fileName := fmt.Sprintf("log%d.csv", fileIndex)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating CSV file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Request", "Latency (ms)"})

	for i := 0; i < 100; i++ {
		go func(requestNumber int) {
			start := time.Now()
			req, err := http.NewRequest("GET", serverAddr, nil)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			latency := time.Since(start).Milliseconds()

			// logging
			fmt.Printf("Response: %s, Latency: %dms\n", resp.Status, latency)
			writer.Write([]string{
				strconv.Itoa(requestNumber + 1),
				strconv.FormatInt(latency, 10),
			})
		}(i)

		time.Sleep(50 * time.Millisecond) // simulate tinygram
	}

	fmt.Printf("Latency data has been saved to %s\n", fileName)
}
