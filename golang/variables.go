package main

import "fmt"

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// zero-value int 
	var e int
	fmt.Println(e)

	// available only inside functions
	f := "apple"
	fmt.Println(f)
}
