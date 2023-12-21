package main

import (
	"fmt"
	"slices"
)

func slicesConstruct() {

	// uninitialized slice
	// always nil and has length of zero
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// empty slice with non-zero length
	// cap(capcity) = length by default
	// can pass a capacity to make if we know the capacity ahead of time
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// append operation
	// slice needs to accept a return value from append
	// as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy a slice (this is a deep copy)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", s)

	// this a slice of a slice
	// it is a shallow copy - modifying s will modify l
	// to avoid this, use the copy function for a deep copy
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// multi dimensional slice
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
