package main

import "fmt"

func main() {
	// 1. Create a channel of type string
	messages := make(chan string)

	// 2. Start a new goroutine that sends "ping" into the channel
	go func() {
		messages <- "ping" // send "ping" into the channel
	}()

	// 3. Receive the message from the channel
	msg := <-messages // wait and get the value from the channel

	// 4. Print the received message
	fmt.Println(msg) // Output: ping
}
