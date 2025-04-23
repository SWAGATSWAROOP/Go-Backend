package main

import "fmt"

// Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs. In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.

func main() {
	fmt.Println("")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)
	// This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for ele := range queue {
		fmt.Println(ele)
	}
}

// This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
