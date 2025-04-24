package main

import (
	"fmt"
	"sync"
)

// In the previous example we saw how to manage simple counter state using atomic operations. For more complex state we can use a mutex to safely access data across multiple goroutines.

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Container holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access. Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.

func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	// This function increments a named counter in a loop.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
		wg.Done()
	}
	// Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 3000)

	wg.Wait()
	fmt.Println(c.counters)
}
