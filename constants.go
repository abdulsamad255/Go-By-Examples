// Package main tells Go that this is an executable program
package main

// Import required packages
import (
	"fmt"  // fmt is used for printing output
	"math" // math is used for mathematical functions like Sin
)

// A constant string variable
// Constants cannot be changed once declared
const s string = "constant"

// main function: program execution starts here
func main() {

	// Print the constant string
	fmt.Println(s)

	// Declare a numeric constant
	// Go automatically decides its type
	const n = 500000000

	// Scientific notation:
	// 3e20 means 3 × 10^20
	// d = (3 × 10^20) / 500000000
	const d = 3e20 / n

	// Print the value of d
	// Go prints large numbers in scientific notation
	fmt.Println(d) // Output: 6e+11

	// Convert d (float) into int64
	// Decimal part (if any) is removed
	fmt.Println(int64(d)) // Output: 600000000000

	// Calculate sine of n
	// NOTE: math.Sin() uses radians, not degrees
	fmt.Println(math.Sin(n))
}
