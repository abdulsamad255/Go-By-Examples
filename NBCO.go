// Non-blocking channel operations

package main

import "fmt"

func main() {
	messages := make(chan string) // unbuffered channel
	signals := make(chan bool)    // another unbuffered channel

	// Non-blocking receive
	select {
	case msg := <-messages: // will execute if a message is ready
		fmt.Println("received message", msg)
	default: // executes immediately if no message
		fmt.Println("no message received")
	}

	// Non-blocking send
	msg := "hi"
	select {
	case messages <- msg: // will execute if a receiver is ready
		fmt.Println("sent message", msg)
	default: // executes immediately if no receiver
		fmt.Println("no message sent")
	}

	// Multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity") // executes if neither channel is ready
	}
}
