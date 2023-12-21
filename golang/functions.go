package main

import "fmt"

func plus(a int, b int) int {
	// go requires explicit returns
	return a + b
}

// can omit the type name for like typed parameters up to the final
// parameter for multiple consecutive parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func functionConstruct() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
