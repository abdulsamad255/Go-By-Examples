package main

import "fmt"

func main() {
	jobs := make(chan int, 5) // channel to send jobs
	done := make(chan bool)   // channel to signal completion

	// Worker goroutine
	go func() {
		for {
			j, more := <-jobs // receive job and check if channel is closed
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // notify main that all jobs are done
				return
			}
		}
	}()

	// Send 3 jobs to the worker
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs) // indicate no more jobs
	fmt.Println("sent all jobs")

	<-done // wait for worker to finish

	// Reading from a closed channel
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok) // ok is false since channel is closed
}
