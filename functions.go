package main

import "fmt"

// -------------------------
// 1. Function plus
// -------------------------
// Takes two integers and returns their sum
func plus(a int, b int) int {
	return a + b
}

// -------------------------
// 2. Function plusPlus
// -------------------------
// Takes three integers and returns their sum
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// -------------------------
	// 3. Call plus function
	// -------------------------
	res := plus(1, 2)         // res = 1 + 2
	fmt.Println("1+2 =", res) // Output: 1+2 = 3

	// -------------------------
	// 4. Call plusPlus function
	// -------------------------
	res = plusPlus(1, 2, 3)     // res = 1 + 2 + 3
	fmt.Println("1+2+3 =", res) // Output: 1+2+3 = 6
}
