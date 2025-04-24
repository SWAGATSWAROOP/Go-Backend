package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

func worker(id int, jobs <-chan int, result chan<- int) {
	for range jobs {
		fmt.Println("Worker ID ", id, " Started")
		time.Sleep(2 * time.Millisecond)
		result <- id * 2
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int)
	result := make(chan int)
	// Start Workers
	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, result)
		}(i)
	}

	const numJobs int = 10
	for i := range numJobs {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
		wg.Wait()
		close(result) //After Wait Group because one should wait all data is sent to this then only close
	}()
	// Collect Results
	for res := range result {
		fmt.Println("Result: ", res)
	}
}

// Making Buffred Channel Works, But Also we can Make a separate Go Routine For Collecting Results

// Important

/*
This WaitGroup is used to wait for all the goroutines launched here to finish. Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.

Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the errgroup package.
*/
