package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs. In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent Job ", j)
	}
	close(jobs)
	fmt.Println("Sent All Jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs: ", ok)
}

// Reading from a closed channel succeeds immediately, returning the zero value of the underlying type. The optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero value generated because the channel is closed and empty.
