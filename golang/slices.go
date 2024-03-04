package main

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go giving a more powerful interface to sequences than arrays
func sliceConstruct() {
	fmt.Println("SLICES")

	// Slices are typed only by the elements they contain (not the number of elements)
	// An uninitalized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// make is used to create slices
	// Capacity default: length of slice
	// If we know slice is going to grow ahead of time, can define a capacity explicitly as a third parameter
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// append - returns a slice containing one ore more values
	// need to accept a return value from append as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy built in function 
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice - slice[low:high] (excluding high element)
	// slice from index 2(inclusive) to index 5(exclusive)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slice up to index 5(everything but element at index 5)
	l = s[:5]
	fmt.Println("sl2:", l)

	// slice from index 2(inclusive) to the end
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize a slice in asingle line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slice utility library 
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Multi-dimensional slice structure
	// Inner slice can vary unlike multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
