package main

import (
	"fmt"
	"time"
)

// This function is run in a gortouine. The done channel will be used to notify
// another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//send value to notify that we're done
	done <- true
}

// channels can be used to synchornize execution across
// goroutines. This example is using a blocking receive to wait for
// a goroutine to finish. When waiting for multiple goroutines to finish
// use a waitGroup
func channelSynchronization() {
	// start worker goroutine giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// block until we receive a notification from the worker on the channel
	// removing the <- done line from this program would cause the program
	// to exit before the worker even started
	<-done
}
