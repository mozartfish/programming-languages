package main

import "fmt"

// In Go, variables are explicitly declared and used by the compiler
// to check correctness of function calls

func variables() {

	fmt.Println("VARIABLES")
	// - var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// - can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// - go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// - variables declared without a corresponding initialization
	// are zero-valued ie. zero value for int is 0

	var e int
	fmt.Println(e)

	// - The := syntax is shorthand for declaring and initializing a variable.
	// This syntax is only available inside functions
	f := "apple"
	fmt.Println(f)
}
