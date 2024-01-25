package main

import "fmt"

// for is Go's only looping construct
func forConstruct() {

	fmt.Println("FOR")
	// - while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// - classic for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// - infinite loop
	for {
		fmt.Println("loop")
		break
	}

	// - continue statement
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
