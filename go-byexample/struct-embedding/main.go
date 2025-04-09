package main

import "fmt"

type base struct {
	num int16
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num=%d", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "hello",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)
	// Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("Describer: ", d.describe())
}

/*
Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base
*/
