package main

import "fmt"

// Go structs are typed collections of fields
// They're useful for grouping data together to form records

type person struct {
	name string
	age  int
}

// return a person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func structConstruct() {

	fmt.Println("STRUCTS")
	// - new struct
	fmt.Println(person{"Bob", 20})

	// - name fields explicitly
	fmt.Println(person{name: "Alice", age: 30})

	// - omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// - pointer to a person struct
	fmt.Println(&person{name: "Ann", age: 40})

	// - it is idomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("John"))

	// - access struct fields with a dot
	s := person{name: "Harold", age: 50}
	fmt.Println(s.name)

	// - access struct pointer fields -> the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// - structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// - if a struct type is only used for a single value, we don't have to give it a name.
	// - the value can have an anonymous struct type
	// - this technique is commonly used for table-driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
