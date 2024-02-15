package main

import "fmt"

// Accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts one channel for receives (pings) and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// When using channels as function parameters you
// can specify if a channel is mean to only send or receive values
// This increases the type-safety of the program
func channelDirections() {
	fmt.Println("CHANNEL DIRECTIONS")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
