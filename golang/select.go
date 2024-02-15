package main

import (
	"fmt"
	"time"
)

// Go's select lets you wait on multiple channel operations
// Combining goroutines and channels with select is a powerful feature in go
// - note that the total execution time is only  ~2 seconds since both the 1 and 2 second Sleeps execute concurrently
func selectConstruct() {
	fmt.Println("SELECT")

	// - channels for selecting across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// - each channel will receive a value after some amount of time
	// - time simulates blocking rpc operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// - select awaits both of the values simultaneously
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
