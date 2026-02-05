package main

import "fmt"

// send-only channel: can only send messages
func ping(pings chan<- string, msg string) {
	pings <- msg // send the message into the channel
}

// receive from one channel, send to another
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // receive from pings
	pongs <- msg   // send to pongs
}

func main() {
	pings := make(chan string, 1) // buffered send/receive channel
	pongs := make(chan string, 1) // buffered send/receive channel

	ping(pings, "passed message") // send a message
	pong(pings, pongs)            // move message from pings -> pongs

	fmt.Println(<-pongs) // receive message from pongs
}
