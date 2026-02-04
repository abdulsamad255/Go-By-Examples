package main

import (
	"errors"
	"fmt"
)

// -------------------------------
// Custom error struct
// -------------------------------
type argError struct {
	arg     int    // value that caused error
	message string // explanation
}

// Error method makes argError satisfy the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// Function that returns a value or a custom error
func f(arg int) (int, error) {
	if arg == 42 {
		// Return a custom error
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Call function with bad value
	_, err := f(42)

	// Use errors.As to check if err is of type *argError
	var ae *argError
	if errors.As(err, &ae) {
		// Extract info from custom error
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
