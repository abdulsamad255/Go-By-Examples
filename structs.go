package main

import "fmt"

// Defining a struct named person
// It groups name and age together
type person struct {
	name string
	age  int
}

// Function that creates a new person
// It returns a pointer to person
func newPerson(name string) *person {

	// Create struct with name value
	p := person{name: name}

	// Assign default age
	p.age = 42

	// Return address (pointer) of struct
	return &p
}

func main() {

	// Creating struct using positional values
	fmt.Println(person{"Bob", 20})

	// Creating struct using named fields
	fmt.Println(person{name: "Alice", age: 30})

	// Creating struct with only name
	// Age will be default (0)
	fmt.Println(person{name: "Fred"})

	// Creating struct and returning pointer
	fmt.Println(&person{name: "Ann", age: 40})

	// Creating struct using function
	fmt.Println(newPerson("Jon"))

	// Creating struct and storing in variable
	s := person{name: "Sean", age: 50}

	// Accessing struct field
	fmt.Println(s.name)

	// Creating pointer to struct
	sp := &s

	// Accessing struct field using pointer
	// Go automatically dereferences pointer
	fmt.Println(sp.age)

	// Modifying struct value using pointer
	sp.age = 51
	fmt.Println(sp.age)

	// Anonymous struct (struct without name)
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	// Printing anonymous struct
	fmt.Println(dog)
}
