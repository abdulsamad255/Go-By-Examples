// Package main defines an executable program
package main

// Import required packages
import (
	"fmt"  // For printing output
	"time" // For working with time and dates
)

func main() {

	// -------------------------
	// 1. Simple value switch
	// -------------------------

	i := 2
	fmt.Print("Write ", i, " as ")

	// Switch compares the value of i
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// -------------------------
	// 2. Switch with multiple values
	// -------------------------

	// Get the current weekday
	switch time.Now().Weekday() {

	// If today is Saturday or Sunday
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")

	// Default runs if no case matches
	default:
		fmt.Println("It's a weekday")
	}

	// -------------------------
	// 3. Condition-based switch
	// -------------------------

	t := time.Now()

	// Switch without a value checks boolean conditions
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// -------------------------
	// 4. Type switch
	// -------------------------

	// Function that checks the type of input
	whatAmI := func(i interface{}) {

		// Type switch checks the actual data type
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	// Call function with different types
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
