package main

import "fmt"

// takes two ints and returns an int sum
func plus(a int, b int) int {
	return a + b
}

// - When you have multiple consecutive parameters of the same type,
// you can omit the type name for the like-typed parameters yp to the final parameter
// that declares the type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// - go requires explicit returns, won't automatically return the value of the last expression
func functions() {
	fmt.Println("FUNCTIONS")
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
