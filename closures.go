package main

import "fmt"

// -------------------------
// 1. intSeq returns a closure
// -------------------------
// A closure captures its surrounding variable `i`
func intSeq() func() int {
	i := 0
	return func() int { // anonymous function returned
		i++ // increments captured variable
		return i
	}
}

func main() {

	// -------------------------
	// 2. Create first counter
	// -------------------------
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// -------------------------
	// 3. Create a new independent counter
	// -------------------------
	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
