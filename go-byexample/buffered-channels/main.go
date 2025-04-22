package main

import "fmt"

// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	// Here we make a channel of strings buffering up to 2 values.
	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	// Also in same go routine or thread can have this channel not in previous unbufferred channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
