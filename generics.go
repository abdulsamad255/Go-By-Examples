package main

import "fmt"

// -------------------------------
// Generic function to find index in any slice
// S = slice type, E = element type
// ~[]E means S must be a slice of E
// comparable = allows == operator
// -------------------------------
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1 // not found
}

// -------------------------------
// Generic Linked List
// -------------------------------
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push a value to the list
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Convert all elements to a slice
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	// Generic function example with slice of strings
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// Explicit type parameters (optional)
	_ = SlicesIndex[[]string, string](s, "zoo")

	// Generic linked list example
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Get all elements as slice
	fmt.Println("list:", lst.AllElements())
}
