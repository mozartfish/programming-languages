package main

import (
	"fmt"
	"maps"
)

// maps are Go's built in associate data type (hashMap, dicts in other languages)
// - maps print in the form map[k:v k:v]
func mapConstruct() {

	fmt.Println("MAPS")
	// - string-integer map
	m := make(map[string]int)

	// - set key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	// - print map
	fmt.Println("map:", m)

	// - get value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// - if key does not exist, zero value is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// - builtin len returns the number of key-value pairs in a map
	fmt.Println("len:", len(m))

	// - builtin delete removes key-value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// - builtin clear removes all key-value pairs from a map
	clear(m)
	fmt.Println("map:", m)

	// - optional second return value when getting a value from a map
	// indicates whether if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// - declare and initialize a map in 1 line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// - map package contains a number of utility functions
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
