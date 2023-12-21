package main

import "fmt"

// function that takes an arbitrary number of integer arguments
// within function nums is equivalent to []int
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Variadic functions can be called with any number
// of trailing arguments. fmt.Println is a common variadic function
func variadicFunctions() {
	sum(1, 2)
	sum(1, 2, 3)

	// apply multiple args in a slice to a variadic function
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
