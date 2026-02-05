package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Timeout occurs (channel sends after 2s, timeout is 1s)
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // Simulate long operation
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res) // won't execute
	case <-time.After(1 * time.Second): // timeout after 1s
		fmt.Println("timeout 1")
	}

	// Example 2: Operation succeeds (channel sends after 2s, timeout is 3s)
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // Simulate operation
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res) // prints "result 2"
	case <-time.After(3 * time.Second): // timeout after 3s
		fmt.Println("timeout 2")
	}
}
