package main

import "fmt"

// -------------------------
// 1. Factorial using recursion
// -------------------------
// fact(n) returns n! = n * (n-1) * ... * 1
func fact(n int) int {
	if n == 0 { // base case
		return 1
	}
	return n * fact(n-1) // recursive call
}

func main() {

	// -------------------------
	// 2. Call factorial
	// -------------------------
	fmt.Println(fact(7)) // Output: 5040

	// -------------------------
	// 3. Fibonacci using recursion
	// -------------------------
	var fib func(n int) int // declare function variable

	fib = func(n int) int { // assign anonymous function
		if n < 2 { // base cases: fib(0)=0, fib(1)=1
			return n
		}
		return fib(n-1) + fib(n-2) // recursive call
	}

	fmt.Println(fib(7)) // Output: 13
}
