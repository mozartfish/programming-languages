package main

import "fmt"

// range iterates over elements in a variety of data structures
func rangeConstruct() {

	// range works on both arrays and slices

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// iterate over key-value pairs in map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate over map keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// iterating over unicode code points
	// i - starting byte index
	// c - rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
