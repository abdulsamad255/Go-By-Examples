package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Atomic counter (starts at 0)
	var ops atomic.Uint64

	// WaitGroup waits for all goroutines to finish
	var wg sync.WaitGroup

	// We will start 50 goroutines
	for i := 0; i < 50; i++ {

		// Tell WaitGroup: one goroutine is starting
		wg.Add(1)

		go func() {
			// Tell WaitGroup: this goroutine is done at the end
			defer wg.Done()

			// Each goroutine increments the counter 1000 times
			for j := 0; j < 1000; j++ {

				// Atomic increment (SAFE)
				ops.Add(1)
			}
		}()
	}

	// Wait until all goroutines finish
	wg.Wait()

	// Safely read the final value
	fmt.Println("ops:", ops.Load())
}
