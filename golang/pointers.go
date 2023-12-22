package main

import "fmt"

// Zeroval has an int parameter so its arguments will be passed to it
// by value. Zeroval will get a copy of ival distinict from the one
// in the calling function
func zeroval(ival int) {
	ival = 0
}

// Zeroptr has an *int parameter which means it takes an int pointer
// The *iptr code in the function dereferences the pointer from its memory address
// to the current value at the address. Assigning a value to a dereferenced pointer changes the value
// at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

// Go supports pointers allowing you to pass references to values and records
// within your program
func pointers() {
	// zeroval doesn't change but zeroptr changes
	// because it has a reference to the memory address for the variable
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gives the memory address of i -> pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
