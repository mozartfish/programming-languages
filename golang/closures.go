package main

import "fmt"

// returns another function.
// returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Go supports anonymous function which can form closures
// anonymous functions are useful when wanting to define functions
// with no names inline
func closures() {
	nextInt := intSeq()

	fmt.Println(nextInt())

	fmt.Println(nextInt())

	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
