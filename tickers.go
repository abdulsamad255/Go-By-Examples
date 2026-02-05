package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 500 milliseconds
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Goroutine to receive ticker events
	go func() {
		for {
			select {
			case <-done:
				return // stop the goroutine when done
			case t := <-ticker.C: // receive the tick
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Let the ticker run for ~1.6 seconds
	time.Sleep(1600 * time.Millisecond)

	// Stop the ticker
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
