package main

import "fmt"

// Go supports pointers allowing you to pass references to values
// and records within your program

// - takes an int parameter, arguments are passed to it by value
// - gets a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// - takes a *int parameter(int pointer)
// - *iptr code in the function body dereferences the pointer from its memory
// address to the current value at that address
// - assigning a value to the dereferenced pointer changes the value at the referenced
// address
// - zeroval doesn't change the i in in main, but zeroptr does because it references
// the memory address for the variable i
func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers() {

	fmt.Println("POINTERS")
	i := 1
	fmt.Println("i initial value:", i)

	zeroval(i)
	fmt.Println("zeroval call: ", i)

	// - &i gives the memory address of i; a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr call: ", i)

	// - print pointer (i memory address)
	fmt.Println("pointer(i memory address):", &i)

}
