package main

import (
	"errors"
	"fmt"
)

// - by convention errors are the last return value and have type error, a built in interface
func f1(arg int) (int, error) {
	if arg == 42 {

		// - errors.New constructs a basic error value with the given error message
		return -1, errors.New("can't work with 42")

	}

	// - nil in the rror position indicates that there was no error
	return arg + 3, nil
}

// - its possible to use custom types as errors by implementing the Error() method on them
// - custom type to explicitly represent an argument error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// - &argError builds a new struct suypplying values for the two fields arg and prob
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// In Go, its idiomatic to communicate errors via an explicit separate return value
// This is different with exceptions in Java, Ruby, other languages and the overloaded single
// result/error value sometimes used in C. Go's approach makes it easy to see which functions return errors and to handle
// them using the same language constructs employed for any other, non-error tasks
func errorsConstruct() {

	fmt.Println("ERRORS")

	// - test error returning functions
	for _, i := range []int{7, 42} {
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

	// - if you want to programatically use the data in a custom error you need to get the error as an instance of the custom error
	// type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
