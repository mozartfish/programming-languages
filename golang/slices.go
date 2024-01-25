package main

import (
	"fmt"
	"slices"
)

// Slices are an import data type in Go, giving a more powerful interface
// to sequences than arrays
// - slices of slices modify the original slice (shallow copy)

func sliceConstruct() {

	fmt.Println("SLICES")
	// - slices are typed only by the elements they contain (not the number of elements)
	// - an uninitialized slice equals nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// - slice of strings of length 3 (zero-valued)
	// - new slice capacity is equal to its length
	// - can set capacity explicitly if we know its going to grow ahead
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// - slices can be set like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// - len returns the length of a slice
	fmt.Println("len:", len(s))

	// - builtin append returns a slice containing one or more new values
	// need to accept a return value from append to get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// - builtin copy creates a slice of the same length and copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// - slices support a slice operatation with slice[low:high] similar to python
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// - declare and initialize a slice on 1 line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// - slice package contains a bunch of utility functions
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// - multi-dimensional slices
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
