// Package main indicates this is an executable Go program
package main

// Import fmt package for printing output
import "fmt"

func main() {

	// -------------------------
	// 1. Even or odd check
	// -------------------------

	// Check if 7 is divisible by 2
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// -------------------------
	// 2. if without else
	// -------------------------

	// Check if 8 is divisible by 4
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// -------------------------
	// 3. Logical OR condition
	// -------------------------

	// If either 8 or 7 is even, condition becomes true
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// -------------------------
	// 4. if with short declaration
	// -------------------------

	// num is declared and used only inside this if block
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
