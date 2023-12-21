package main

import "fmt"

// go's only looping construct
func forConstruct() {

	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)

	}

	// infinite loop
	for {
		fmt.Println("loop")
		break
	}

	// continue statement in a for loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
