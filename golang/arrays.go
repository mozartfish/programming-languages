package main

import "fmt"

// In Go, ab array is a number sequence of elements of a specific length.
// In Go, slices are much more common, arrays are useful in some special scenarios
func arrays() {
	fmt.Println("ARRAYS")

	// Create an array that can hold 5 elements 
	// By default, an array is zero-valued - for ints means 0s
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// Get length of array
	fmt.Println("len:", len(a))

	// Declare array in 1 line 
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Two-dimensional array structure 
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
