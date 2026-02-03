// Package main indicates this is an executable Go program
package main

// Import fmt package for printing output
import "fmt"

func main() {

	// -------------------------
	// 1. While-style for loop
	// -------------------------

	// Initialize variable i
	i := 1

	// Loop runs while i is less than or equal to 3
	for i <= 3 {
		fmt.Println(i) // Print current value of i
		i = i + 1      // Increase i by 1
	}

	// -------------------------
	// 2. Classic for loop
	// -------------------------

	// j starts from 0, runs while j < 3, increments each loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// -------------------------
	// 3. Range-based loop
	// -------------------------

	// range 3 generates values: 0, 1, 2
	for i := range 3 {
		fmt.Println("range", i)
	}

	// -------------------------
	// 4. Infinite loop with break
	// -------------------------

	// for {} creates an infinite loop
	for {
		fmt.Println("loop")
		break // break exits the loop immediately
	}

	// -------------------------
	// 5. Continue example
	// -------------------------

	// range 6 generates values: 0 to 5
	for n := range 6 {

		// If n is even, skip this iteration
		if n%2 == 0 {
			continue
		}

		// This line runs only for odd numbers
		fmt.Println(n)
	}
}
