package main

import "fmt"

// Basic sends and receives on channels are blocking
// Select can be used with a default clause to implement non-blocking sends,
// receives and even non-blocking multi-way selects
func nonBlockingConstruct() {
	fmt.Println("NON-BLOCKING CHANNEL OPERATIONS")

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
