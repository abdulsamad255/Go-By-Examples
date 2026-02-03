package main

import "fmt"

// -------------------------
// 1. Variadic function
// -------------------------
// Can take any number of int arguments
func sum(nums ...int) {
	fmt.Print(nums, " ") // print all numbers received

	total := 0
	for _, num := range nums { // loop through slice
		total += num
	}
	fmt.Println(total) // print the sum
}

func main() {

	// -------------------------
	// 2. Call with 2 arguments
	// -------------------------
	sum(1, 2) // [1 2] 3

	// -------------------------
	// 3. Call with 3 arguments
	// -------------------------
	sum(1, 2, 3) // [1 2 3] 6

	// -------------------------
	// 4. Call with slice using "..."
	// -------------------------
	nums := []int{1, 2, 3, 4}
	sum(nums...) // [1 2 3 4] 10
}
