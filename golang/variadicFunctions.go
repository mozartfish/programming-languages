package main

import "fmt"

// takes and arbitrary number of ints and returns sum
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Variadic functions can be called with any number of trailing arguments
// - println is a common variadic function
func variadicFunctions() {
	fmt.Println("VARIADIC FUNCTIONS")
	sum(1, 2)
	sum(1, 2, 3)

	// - variadic function call on slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
