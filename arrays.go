// Package main defines an executable program
package main

// Import fmt package for printing
import "fmt"

func main() {

	// -------------------------
	// 1. Declare an empty array
	// -------------------------
	var a [5]int
	fmt.Println("emp:", a) // [0 0 0 0 0]

	// -------------------------
	// 2. Assign value to an element
	// -------------------------
	a[4] = 100
	fmt.Println("set:", a)    // [0 0 0 0 100]
	fmt.Println("get:", a[4]) // 100

	// -------------------------
	// 3. Array length
	// -------------------------
	fmt.Println("len:", len(a)) // 5

	// -------------------------
	// 4. Declare and initialize
	// -------------------------
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// -------------------------
	// 5. Let Go calculate length automatically
	// -------------------------
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// -------------------------
	// 6. Indexed initialization
	// -------------------------
	b = [...]int{100, 3: 400, 500} // index 0=100, index 3=400, next=500
	fmt.Println("idx:", b)         // [100 0 0 400 500]

	// -------------------------
	// 7. 2D array using loops
	// -------------------------
	var twoD [2][3]int
	for i := range 2 { // rows
		for j := range 3 { // columns
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [[0 1 2] [1 2 3]]

	// -------------------------
	// 8. 2D array direct initialization
	// -------------------------
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD) // [[1 2 3] [1 2 3]]
}
