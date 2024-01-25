package main

import "fmt"

// range iterates over elements in a variety of data structures
func rangeConstruct() {

	fmt.Println("RANGE")
	// - range can be used for slices and arrays
	// - compute total sum
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// - iterate over key, value pairs in a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// - iterate over keys in a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// - range on strings iterates over Unicode code points
	// - first value is the starting byte index of the rune
	// - second value is the reune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
