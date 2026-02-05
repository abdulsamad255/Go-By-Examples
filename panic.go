package main

import "fmt"

func main() {
	fmt.Println("Program started")

	panic("something went wrong")

	fmt.Println("This will NEVER run")
}
