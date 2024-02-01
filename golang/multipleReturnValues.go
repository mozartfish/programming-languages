package main

import "fmt"

// returns two ints
func vals() (int, int) {
	return 3, 7
}

// Go has built in support for multiple return values
// - used in idiomatic go -> return both result and error values from a function
func multiReturn() {
	fmt.Println("MULTIPLE RETURN VALUES")
	// get two different values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// - return subset of returned values
	_, c := vals()
	fmt.Println(c)
}
