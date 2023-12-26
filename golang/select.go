package main

import (
	"fmt"
	"time"
)

// Go select lets you wait on multiple channel operations
// Combining goroutines and channels with select is a powerful feature in go
func selectConstruct() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time
	// (to simulate blocking RPC operations executing in concurrent goroutines)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Use select to await both of these values simultaneoulsy printing each one as it arrives
	// Total execution time is ~2 seconds since both the 1 and 2 second SLeeps execute concurrently
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
