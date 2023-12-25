package main

import "fmt"

// Takes a map of any type and returns a slice of its keys
// This function has two type parameters - K and V; K has the comparable constraint,
// meaning that we can compare values of this type with the == and != operators.
// This is required for map keys in Go. V has the any constraint,
// meaning that it’s not restricted in any way (any is an alias for interface{}).
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Define a link list object for any type T
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Methods on generic types. Have to keep the type parameters in place
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// Go has started supporting generics (also known as type parameters)
// since version 1.18
func generics() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// Can rely on type inference when invoking generic functions
	fmt.Println("keys:", MapKeys(m))

	// specify types explicitly
	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
