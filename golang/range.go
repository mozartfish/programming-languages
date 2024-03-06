package main

import "fmt"

// range iterates over elements in a variety of data structures
func rangeConstruct() {
	fmt.Println("RANGE")

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range on arrays and slices provide both the index and value for each entry
	// blank identifier _ ignores the index/value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range iteration on map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range iteration on map keys
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range iteration on strings
	// range on strings iterates over Unicode code points
	// The first value is the starting by index of the rune, the second is the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
