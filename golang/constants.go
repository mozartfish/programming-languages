package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean and numeric values

// const declares a constant value
const s string = "constant"

func constants() {
	
	fmt.Println("CONSTANTS")
	fmt.Println(s)

	// - a const statement can appear anywhere a var statement can
	const n = 500000000

	// - constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// - a numeric constant has no type until it's given one, such as by an explicit conversion
	fmt.Println(int64(d))

	// - a number can be given a type by using it in a context that requires one such as a variable
	// assignment or function cal.
	fmt.Println(math.Sin(n))
}
