package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// Make Buffered or unbuffered  channel it doesn't affect
	<-done
}
/*
This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
Start a worker goroutine, giving it the channel to notify on.
Block until we receive a notification from the worker on the channel.                
If you removed the <- done line from this program, the program would exit before the worker even started.
*/