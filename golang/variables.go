package main

import "fmt"

// variables are explicity declared and used by the compiler to check type-correctness
// of function calls
func variables() {

	// declare 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go infers type of initialized variables
	var d = true
	fmt.Println(d)

	// variables declared without initialization are zero-valued
	var e int
	fmt.Println(e)

	// shorthand for declaring and initializing variables
	// only available inside functions
	f := "apple"
	fmt.Println(f)
}
