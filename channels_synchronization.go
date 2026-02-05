package main

import (
	"fmt"
	"time"
)

// worker simulates some work and notifies when done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second) // simulate work
	fmt.Println("done")
	done <- true // send signal that work is finished
}

func main() {
	// create a channel to receive completion signal
	done := make(chan bool, 1)

	// start the worker goroutine
	go worker(done)

	// wait here until worker signals completion
	<-done
}
