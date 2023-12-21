package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// multiple return values are used to return both result and error
// values from a function. Common in idiomatic go
func multipleReturnValues() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// return asubset of the returned values
	_, c := vals()
	fmt.Println(c)
}
