package main

import "fmt"

// Gp structs are typed collections of fields
// They're useful for grouping data together to form records

// Person struct
type person struct {
	name string
	age  int
}

// constructs a new person struct with a given name
// returns a pointer to local variable as a local
// variable will survive the scope of the function
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func structs() {
	// create new struct
	fmt.Println(person{"Bob", 20})

	// name fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero valued
	fmt.Println(person{name: "Fred"})

	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// idiomatic to encapsulate new struct creation in constructor
	// function
	fmt.Println(newPerson("Jon"))

	// access struct with dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// access struct pointers which are automatically derferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// anonymous struct type. Used a lot in testing
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
