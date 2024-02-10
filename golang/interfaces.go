package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures
type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// - to implement an interface in go we need to implement all the methods
// in the interface

// geometry on rects
func (r rect1) area() float64 {
	return r.width * r.height
}

func (r rect1) perim() float64 {
	return 2*r.width + 2*r.height
}

// geometry on circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// - if a variable has an interface type then we can call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("AREA:", g.area())
	fmt.Println("PERIMETER:", g.perim())
}

func interfaces() {

	fmt.Println("INTERFACES")
	r := rect1{width: 3, height: 4}
	c := circle{radius: 5}

	// - circle and rect both implement geometry interface so we can use instances
	// of these structs as arguments to measure
	measure(r)
	measure(c)
}
