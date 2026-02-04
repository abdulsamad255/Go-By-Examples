package main

import (
	"fmt"
	"time"
)

// Simple function that prints a message and index
func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// -------------------------------
	// 1. Normal function call (synchronous)
	// -------------------------------
	f("direct") // blocks until finished

	// -------------------------------
	// 2. Goroutine function call (asynchronous)
	// -------------------------------
	go f("goroutine") // runs concurrently

	// -------------------------------
	// 3. Anonymous goroutine with argument
	// -------------------------------
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// -------------------------------
	// 4. Wait for goroutines to finish
	// -------------------------------
	time.Sleep(time.Second) // simple way to wait

	fmt.Println("done")
}
