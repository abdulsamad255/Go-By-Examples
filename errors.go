package main

import (
	"errors"
	"fmt"
)

// -------------------------------
// Function that returns a result or an error
// -------------------------------
func f(arg int) (int, error) {
	if arg == 42 {
		// Return an error if argument is 42
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// -------------------------------
// Predefined custom errors
// -------------------------------
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

// -------------------------------
// Function that simulates making tea
// -------------------------------
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// Wrap an existing error with context
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	// -------------------------------
	// Example 1: calling f() with different arguments
	// -------------------------------
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e) // handle error
		} else {
			fmt.Println("f worked:", r) // use result
		}
	}

	// -------------------------------
	// Example 2: calling makeTea() and checking error type
	// -------------------------------
	for i := range 5 {
		if err := makeTea(i); err != nil {

			// Check specific error types
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		// No error, tea is ready
		fmt.Println("Tea is ready!")
	}
}
