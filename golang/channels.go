package main

import "fmt"

// channels are the pipes that connect concurrent goroutines
// send values into channels from one go routine and receive those
// values into another goroutine
func channels() {

	// create a new channel with make (chan val-type)
	// channels are typed by the values they convey
	messages := make(chan string)
	go func() {
		// send a value into a channel using the channel <- syntax
		// sending ping to the messages channel from a new go routine
		messages <- "ping"
	}()
	// <- channel syntax receives a value from the channel
	// receive ping and then print it
	msg := <-messages

	// by default sends and receives block until both the sender
	// and receiver are ready. This property allowed us to wait at the
	// end of our program for the ping message without having to use any other
	// synchronization
	fmt.Println(msg)
}
