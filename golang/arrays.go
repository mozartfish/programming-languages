package main

import "fmt"

// In Go, an array is a numbered sequence of elements of a specific length.
// In typical Go code, slices are more common. Arrays are useful in some
// special scenario.
func arrayConstruct() {

	fmt.Println("ARRAYS")
	// - array that holds exactly 5 ints
	// - by default an array is zero-valued
	var a [5]int
	fmt.Println("emp:", a)

	//- set array values
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// - builtin len returns the length of an array
	fmt.Println("len:", len(a))

	// - declare and initialize array in 1 line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// - multi-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
