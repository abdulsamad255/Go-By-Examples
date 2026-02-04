package main

import "fmt"

// Defining a struct for rectangle
type rect struct {
	width, height int
}

// Method with pointer receiver
// Calculates area of rectangle
func (r *rect) area() int {
	return r.width * r.height
}

// Method with value receiver
// Calculates perimeter of rectangle
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	// Creating a rectangle object
	r := rect{width: 10, height: 5}

	// Calling methods using struct value
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Creating pointer to rectangle
	rp := &r

	// Calling methods using struct pointer
	// Go automatically dereferences pointer
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
