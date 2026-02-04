package main

import (
	"fmt"
	"iter"   // library for sequences
	"slices" // utility functions for slices
)

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

// Push adds a new value to the linked list
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returns a sequence for iteration
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// iterate linked list and yield each element
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) { // stop if yield returns false
				return
			}
		}
	}
}

// -------------------------------
// Fibonacci generator using iter.Seq
// -------------------------------
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) { // yield a number, stop if false
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	// -------------------------------
	// Linked list example
	// -------------------------------
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Iterate over list using range
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Collect all elements into a slice
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// -------------------------------
	// Fibonacci generator example
	// -------------------------------
	for n := range genFib() {
		if n >= 10 { // stop after 10
			break
		}
		fmt.Println(n)
	}
}
