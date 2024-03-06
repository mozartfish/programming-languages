package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// Go has built in support for multiple return values
// This feature is used often in idomatic GO, for example to return both result and error values from a function
func multipleReturnValues() {
	fmt.Println("MULTIPLE RETURN VALUES")

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Return subset of the returned values -> blank identifier ignores first value
	_, c := vals()
	fmt.Println(c)

}
