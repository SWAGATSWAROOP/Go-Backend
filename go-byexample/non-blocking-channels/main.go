package main

import "fmt"

//Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.

func main() {
	messages := make(chan string) //Make buffered channel to sent message and receive
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received Message:", msg)

	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent Message", msg)
	default:
		fmt.Println("No message was sent")
	}
	// Here, you try to send a message to the messages channel. But since no goroutine is concurrently receiving from messages, and this is inside a select with a default, the send is non-blocking and skips to the default
	select {
	case msg := <-messages:
		fmt.Println("Received Message", msg)

	case sig := <-signals:
		fmt.Println("received signal", sig)

	default:
		fmt.Println("No Activity")
	}
	// We can use multiple cases above the default clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both messages and signals.
}
