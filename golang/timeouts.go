package main

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time
// Implementing timouts in go is easy and elegant thanks to channels and select
func timeOut() {
	fmt.Println("TIMEOUTS")

	// The code simulates executing an external call that returns its result on a channel c1
	// after 2 seconds. Since the channel is buffered, the send in the goroutine is nonblocking
	// This is a common pattern to prevent goroutine leakes in case the channel is never read
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// The select is implementing a timeout. res :-= <- c1 awaits the result and <-time.After
	// awaits a value to be sent out after the timeout of 1 second
	// Since select proceeds with the first receive that's ready, we take the timeout case if the operation takes
	// more than 1 second
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If a longer timeout of 3 ecounds is allowed, then the receive from c2 will succeed and get printed 
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
