package main

import (
	"fmt"
	"maps" // Go 1.21+ for map utilities
)

func main() {

	// -------------------------
	// 1. Create an empty map
	// -------------------------
	m := make(map[string]int)

	// -------------------------
	// 2. Assign values
	// -------------------------
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m) // map[k1:7 k2:13]

	// -------------------------
	// 3. Access values
	// -------------------------
	v1 := m["k1"] // existing key
	fmt.Println("v1:", v1)

	v3 := m["k3"] // non-existing key → returns zero value
	fmt.Println("v3:", v3)

	// -------------------------
	// 4. Map length
	// -------------------------
	fmt.Println("len:", len(m)) // 2

	// -------------------------
	// 5. Delete a key
	// -------------------------
	delete(m, "k2")
	fmt.Println("map:", m) // map[k1:7]

	// -------------------------
	// 6. Clear the map
	// -------------------------
	clear(m)
	fmt.Println("map:", m) // map[]

	// -------------------------
	// 7. Check if key exists
	// -------------------------
	_, prs := m["k2"]
	fmt.Println("prs:", prs) // false

	// -------------------------
	// 8. Declare & compare maps
	// -------------------------
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
