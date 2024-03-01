package main

import "fmt"

// Go has various value types including strings, integers, floats booleans etc.
func values() {
	fmt.Println("VALUES")

	// Strings
	fmt.Println("go" + "lang")
	// Integers and floats
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	// Booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
