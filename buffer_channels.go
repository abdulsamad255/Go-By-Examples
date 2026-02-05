package main

import "fmt"

func main() {
	// 1. Create a buffered channel that can hold 2 strings
	messages := make(chan string, 2)

	// 2. Send values into the buffered channel without a receiver yet
	messages <- "buffered"
	messages <- "channel"

	// 3. Receive and print the values from the channel
	fmt.Println(<-messages) // Output: buffered
	fmt.Println(<-messages) // Output: channel
}
