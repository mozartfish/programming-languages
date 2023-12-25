package main

import (
	"fmt"
	"math"
)

// interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}
type square struct {
	side float64
}
type circle struct {
	radius float64
}

// implement this interface for squares
func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perim() float64 {
	return 4 * s.side
}

// implement this interface for circles
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// If a variable has an interface type, then we can
// call methods that are named in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Interfaces are a named collection of method signatures in golang
func interfaces() {
	s := square{side: 4}
	c := circle{radius: 5}

	// since rectangles and circles both implement geometry
	// we can use these objects as arguments for measure function
	measure(s)
	measure(c)
}
