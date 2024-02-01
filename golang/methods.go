package main

import "fmt"

type rect struct {
	width, height int
}

// method with receiver type of * rect
func (r *rect) area() int {
	return r.width * r.height
}

// method with value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Go supports methods defined on struct types
// - methods can be defined for either pointer or value receiver types
func methods() {
	fmt.Println("METHODS")

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// - go automatically handles conversion between values and pointers for method calls
	// - you may want to use a pointer receiver type to avoid copying on method calls or to allow
	// the method to mutate the receiving struct
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
