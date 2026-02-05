package main

import (
	"fmt"
	"sync"
)

// Shared counter
var counter = 0

// Mutex to protect the counter
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// Lock before accessing shared data
	mu.Lock()
	counter++
	// Unlock after done
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// Start 5 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Safe final value
	fmt.Println("Final Counter:", counter)
}
