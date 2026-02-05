package main

import "fmt"

func main() {
	queue := make(chan string, 2) // create a buffered channel

	// send values into the channel
	queue <- "one"
	queue <- "two"

	close(queue) // close the channel to signal no more values

	// iterate over values in the channel until it's empty
	for elem := range queue {
		fmt.Println(elem)
	}
}
