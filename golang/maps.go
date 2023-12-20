package main

import (
	"fmt"
	"maps"
)

func mapConstruct() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get value associated with key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if key does not exist, zero value of the value is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	// delete key/value pair from map
	delete(m, "k2")
	fmt.Println("map:", m)

	// remove all key/value pairs from map
	clear(m)
	fmt.Println("map:", m)

	// check if key was present
	// ignore value with _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
