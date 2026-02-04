package main

import "fmt"

// -------------------------------
// zeroval receives a VALUE (copy)
// Changing ival does NOT affect the original variable
// -------------------------------
func zeroval(ival int) {
	ival = 0
}

// -------------------------------
// zeroptr receives a POINTER
// *iptr accesses the original value using its memory address
// This WILL change the original variable
// -------------------------------
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	// Declare integer variable i
	i := 1
	fmt.Println("initial:", i)

	// Call function with value (copy is passed)
	zeroval(i)
	// Original value remains unchanged
	fmt.Println("zeroval:", i)

	// Call function with address of i
	zeroptr(&i)
	// Original value is changed using pointer
	fmt.Println("zeroptr:", i)

	// Print memory address of i
	fmt.Println("pointer:", &i)
}
