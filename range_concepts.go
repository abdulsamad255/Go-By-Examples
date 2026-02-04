package main

import "fmt"

func main() {

	// -------------------------------
	// Range over a slice of integers
	// -------------------------------
	nums := []int{2, 3, 4}
	sum := 0

	// "_" ignores the index, "num" holds the value
	for _, num := range nums {
		sum += num // add each number to sum
	}
	fmt.Println("sum:", sum)

	// -------------------------------
	// Get index when a specific value is found
	// -------------------------------
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// -------------------------------
	// Range over a map (key-value pairs)
	// -------------------------------
	kvs := map[string]string{"a": "apple", "b": "banana"}

	// "k" is the key, "v" is the value
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// -------------------------------
	// Range over only map keys
	// -------------------------------
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// -------------------------------
	// Range over a string
	// -------------------------------
	// "i" is the byte index
	// "c" is the Unicode value (rune)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
