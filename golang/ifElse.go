package main

import "fmt"

// Branching with if and else in Go is straight forward
// - no parentheses around conditions in go but curly braces are required
// - There is no ternary if in Go, so you have to use full if statement even for
// basic conditions
func ifElseConstruct() {

	fmt.Println("IF/ELSE")
	// - basics
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// - if statement without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// - logical operators
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// - statement preceding conditionals; any variables
	// declared in this statement are available in the current
	// and the subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
