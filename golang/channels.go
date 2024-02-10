package main

import "fmt"

// channels are the pipes that connect concurrent goroutines
// you can send values into channels from one goroutine and receive those values into another goroutine
// - when the program runs successfull, the ping message is passed from the goroutine to another via a channel
// - by default sends and receives block until both the sender and receiver are ready
// - this allows us to wait at the end of our program for the ping message without having to use any other synchronization
func channels() {

	fmt.Println("CHANNELS")

	// - create new channel with make(chan val-type)
	// - channels are typed by the values they convey
	messages := make(chan string)
	go func() {
		// - send a value into a channel using the channel <- syntax
		// - this sends ping to the messages channel from a new goroutine
		messages <- "ping"
	}()

	// - <-channel syntax receives a value from the channel
	// - receive the ping message from goroutine and print it
	msg := <-messages
	fmt.Println(msg)
}
