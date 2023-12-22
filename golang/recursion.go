package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func recursion() {
	fmt.Println(fact(7))
	var fib func(n int) int

	// Closures can also be recursive
	// Requires the closure to be declared with a typed var explicitly
	// before its defined
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
