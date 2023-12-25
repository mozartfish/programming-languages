package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// A goroutine is a lightweight thread of execution
func goroutines() {

	// call a function synchronously
	f("direct")

	// execute concurrently with the calling function - main function
	go f("goroutine")

	// anonmous function goroutine call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// two function calls are running asynchronously in separate goroutines
	// Wait 1 second for them to finish
	time.Sleep(time.Second)
	fmt.Println("done")

	// Order of Execution
	//  see the output of blocking call first
	//  then the output of the two goroutines which
	// may be interleaved because goroutines are being run concurrently
	// by the go runtime
}
