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
// - output the blocking call first 
// - output the two goroutines 
// - the goroutines output may be interleaved because go routines run concurrently by the Go runtime
func goRoutines() {

	fmt.Println("GOROUTINES")
	
	// synchronous function 
	f("direct")

	// execute concurrently with the calling function  
	go f("goroutine")

	// goroutine for an anonymous function call 
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// - the two functions calls are running asynchronously in a separate goroutines
	// - wait for them to finish (a more robust approach is waitGroups)
	time.Sleep(time.Second)
	fmt.Println("done")
}
