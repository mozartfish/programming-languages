package main

import "fmt"

func forConstruct() {
	fmt.Println("FOR")

	// While loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// For Loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Iterate range times (up to but not including)
	// for i := range 3 {
	// 	fmt.Println("range", i)
	// }

	// Infinite Loop
	for {
		fmt.Println("loop")
		break
	}

	// Continue statement
	// for n := range 6 {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }
}
