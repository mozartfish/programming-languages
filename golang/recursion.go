package main

import "fmt"

// function that calls i9tself until it reaches the base case of
// fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Go supports recursive functions
func recursion() {
	fmt.Println("RECURSION")
	fmt.Println(fact(7))

	// - Closures can also be recursive but this requires the closure
	// to be declared with a type var explicitly before its defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
