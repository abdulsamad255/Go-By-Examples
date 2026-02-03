package main

import (
	"fmt"
	"slices" // Go 1.21+ for slice utilities
)

func main() {

	// -------------------------
	// 1. Uninitialized slice
	// -------------------------
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// Output: [] true true

	// -------------------------
	// 2. Create slice with make
	// -------------------------
	s = make([]string, 3) // length 3, capacity 3
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// -------------------------
	// 3. Assign values
	// -------------------------
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// -------------------------
	// 4. Append elements dynamically
	// -------------------------
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// -------------------------
	// 5. Copy slices
	// -------------------------
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// -------------------------
	// 6. Slicing (sub-slices)
	// -------------------------
	l := s[2:5] // from index 2 to 4
	fmt.Println("sl1:", l)

	l = s[:5] // first 5 elements
	fmt.Println("sl2:", l)

	l = s[2:] // from index 2 to end
	fmt.Println("sl3:", l)

	// -------------------------
	// 7. Declare and compare slices
	// -------------------------
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// -------------------------
	// 8. 2D slice (dynamic rows & columns)
	// -------------------------
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
