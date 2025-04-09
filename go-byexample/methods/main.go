package main

import "fmt"

// Go support methods
/*
Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
*/

type Car struct {
	model  string
	number int
	length float64
	width  float64
}

func (c *Car) Area() float64 {
	c.length = 2
	return c.length * c.width
}

func (c Car) getLengthAndWitdh() (float64, float64) {
	fmt.Println("Inside this function")
	c.length = 3
	c.width = 4
	return c.length, c.width
}

func main() {
	c := Car{}
	c.model = "Audi"
	c.number = 2

	println(c.Area())

	c.length = 5.9
	c.width = 9.0093
	fmt.Println(c.Area())
	c.getLengthAndWitdh()
	fmt.Println(c.getLengthAndWitdh())
	fmt.Println(c.length, c.width)
}
