package main

import "fmt"

// By default channels are unbuffered - they will only accept sends (chan <-)
// if there is a corresponding receive(<- chan) ready to receive the sent value
// Buffered channels accept a limited number of values without a corresponding
// receiver for those values
func channelBuffering() {
	// make a channel of strings buffering up to two values
	messages := make(chan string, 2)

	// since the channel is buffered we can send these values
	// into the channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// receive the two messages
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
