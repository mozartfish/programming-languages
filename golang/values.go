package main

import "fmt"

// Go has various value types including strings, integers, floats
// booleans etc.
func values() {

	fmt.Println("VALUES")
	// strings (concatentation)
	fmt.Println("go" + "lang")

	// ints, floats
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
