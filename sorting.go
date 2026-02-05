package main

import (
	"fmt"
	"slices"
)

func main() {

	// 1️⃣ Sorting strings
	strs := []string{"c", "a", "b"}
	slices.Sort(strs) // sorts A → Z
	fmt.Println("Strings:", strs)

	// 2️⃣ Sorting integers
	ints := []int{7, 2, 4}
	slices.Sort(ints) // sorts small → big
	fmt.Println("Ints:   ", ints)

	// 3️⃣ Check if slice is sorted
	isSorted := slices.IsSorted(ints)
	fmt.Println("Sorted:", isSorted)
}
