package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a channel with 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Ticker sends a signal every 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	// Process requests one by one
	for req := range requests {

		// Wait for the next tick (rate limit)
		<-limiter

		// Now handle the request
		fmt.Println("request", req, time.Now())
	}
}
