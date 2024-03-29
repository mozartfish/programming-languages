package main

import "fmt"

// In Go variables are explicitly declared and used by the compiler to check type-correctness 
// of function calls 
func variables() {
	fmt.Println("VARIABLES")

	// Declare 1 variable
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued
	// ie. int zero value is 0
	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing variable
	// := syntax is only available inside functions
	f := "apple"
	fmt.Println(f)
}
