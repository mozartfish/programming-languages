package main

import "fmt"

// no ternary statements in go
func ifElseConstruct() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// If statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators in a conditional
	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// a statement can precede conditionals;
	// any variables declared in this statement are available in the current
	// and subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
