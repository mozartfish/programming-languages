package main

import "fmt"

func arrays() {
	var a [5]int
	fmt.Println("empty", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a)
	fmt.Println("len:", len(a))

	// declare and initialize array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println("2d: ", twoD)

}
