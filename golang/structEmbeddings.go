package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// - a container embeds a base
// - an embedding looks like a field without a name
type container struct {
	base
	str string
}

// Go supports embeddings of structs and interfaces to express a more seamless composition of types
func structEmbeddings() {

	fmt.Println("STRUCT EMBEDDING")
	// - when creating structs with literals, we have to initialize the embedding explicitly
	// - the embedding type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	// - since the container embeds base, the methods of base also become methods of a container
	// - invoke a method that was embedded from base directly on co
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// - embedding structs may be used to bestow interface implementations onto other structs 
	// - container now implements the describer interface because it embeds base 
	var d describer = co
	fmt.Println("describer:", d.describe())

}
