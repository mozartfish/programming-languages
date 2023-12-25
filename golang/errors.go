package main

import (
	"errors"
	"fmt"
)

// In Go, it's idiomatic to communciate errors via an explicit
// separate return value which contrasts with exceptions in other languages
// like java, ruby, and overloaded error in C. Go's approach makes it easy to see
// which functions return errors and to handle them using the same language
// constructs employed for any other non-error tasks

// By convention errors are the last return value and have type error
// a built in interface
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New constructs a basic error value with the given error message
		return -1, errors.New("can't work with 42")

	}

	// nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// Custom types as errors by implementing the Error() method on them
// Variant on the example above that uses a custom type to explicitly represent
// an argument error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// use &argError syntax to build a new struct supplying values for the
		// two fields arg and prob
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func errorConstruct() {
	// the two loops test out each of the error returning functions
	for _, i := range []int{7, 42} {
		// inline error check is a common idiom in golang code
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// programatically use the data in a custom error, get the error
	// as an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
