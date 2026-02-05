package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer 1: fires after 2 seconds
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C // wait until timer fires
	fmt.Println("Timer 1 fired")

	// Timer 2: fire after 1 second, but we stop it before it fires
	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired") // this won't run
	}()

	stop2 := timer2.Stop() // stop the timer
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// give some time to see if timer2 fires
	time.Sleep(2 * time.Second)
}
