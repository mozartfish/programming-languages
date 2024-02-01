package main

import "fmt"

// returns another function
// - returned function closes over the variable i to form closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Go supports anonymous functions, which can form closures
// - Anonymous functions are useful when you want to define a function
// inline without having to name it
func closure() {
	fmt.Println("CLOSURES")

	// - assign result to next int 
	// - function captures its own i value which will be updated each time 
	// nextInt is called 
	
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// - state is unique to the function call and variable 
	newInts := intSeq()
	fmt.Println(newInts())

}
