package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

func main() {
	//Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() {
		messages <- "ping"
	}()

	// messages <- "Swagat" //We cannot do this as The messages <- "Swagat" line blocks because there’s no goroutine running concurrently to receive from the channel.
	// As a result, the program hits a deadlock and crashes.

	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
	fmt.Println(msg)
	// By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
}
