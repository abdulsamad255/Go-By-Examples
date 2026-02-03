package main

import "fmt"

// -------------------------
// 1. Function returning two values
// -------------------------
func vals() (int, int) {
	return 3, 7 // returns two integers
}

func main() {

	// -------------------------
	// 2. Receive both return values
	// -------------------------
	a, b := vals() // a = 3, b = 7
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	// -------------------------
	// 3. Ignore first value
	// -------------------------
	_, c := vals() // ignore first, c = 7
	fmt.Println(c) // 7
}
