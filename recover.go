package main

import "fmt"

func mayPanic() {
	panic("something went wrong")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	mayPanic()

	fmt.Println("This line will NOT run")
}
