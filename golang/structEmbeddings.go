package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds a base
type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

// go supports embedding of structs and interfaces to express
// a more seamless composition of types
func structEmbeddings() {
	// initialize the embedding explicitly
	// where the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	// since container embeds base, the methods of base also
	// become methods of a container. We invoke a method
	// that was embedded from base directly on co
	fmt.Println("describe:", co.describe())

	// embeddings structs with methods may be used to bestow interface implementations
	// onto other structs. The container now implements the describer interface because
	// because it embeds it
	var d describer = co
	fmt.Println("describer:", d.describe())

}
