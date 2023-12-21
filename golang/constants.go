package main

import (
	"fmt"
	"math"
)

// declare a constant value
const s string = "constant"

func constants() {
	fmt.Println(s)

	// constants can appear anywhere a var statement can
	const n = 500000000

	// constant expressions perform arithmetic with arbitrary position
	const d = 3e20 / n
	fmt.Println(d)

	// numeric constant has no type until it's given one such as by an explicit conversion
	fmt.Println(int64(d))

	// a number can be given a type by using it in a context
	// that requires one, such as a variable assignment or function call
	fmt.Println(math.Sin(n))
}
