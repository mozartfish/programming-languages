package main

import (
	"fmt"
	"time"
)

// function to run in a goroutine
// the done channel will notify another goroutine that this functions work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// We can use channels to synchronize execution across goroutines
// When waiting for multiple goroutines to finish, use a waitGroup
// - this program uses a blocking receive to wait for a goroutine to finish
func channelSynchronization() {
	fmt.Println("CHANNEL SYNCHRONIZATION")

	done := make(chan bool, 1)

	// - start a worker goroutine giving it the channel to notify on
	go worker(done)

	// - block until we receive a notification from the worker on the channel
	<-done
}
