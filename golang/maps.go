package main

import (
	"fmt"
	"maps"
)

// Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages)
// Maps appear in the form [k:v k:v] when printed to console
func mapConstruct() {
	fmt.Println("MAPS")

	// make is used to create maps 
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	// Print map - shows all key-value pairs
	fmt.Println("map:", m)

	// Get value from map 
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key does not exist, zero-value of the value type is returned 
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// len - returns the number of key-value pairs in map 
	fmt.Println("len:", len(m))

	// delete - remove key-value pairs from a map 
	delete(m, "k2")
	fmt.Println("map:", m)

	// clear - removes all key-value pairs from map
	clear(m)
	fmt.Println("map:", m)

	// Optional second return value when getting value from a map indicates
	// if the key was present in the map. _ -> blank identifier. used for ignoring value 
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize map on one line 
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	// Map utility library
	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
