package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
	if c, ok := g.(*circle); ok {
		fmt.Println("Circle with radius", c.radius)
	} else {
		fmt.Println("Its a reactangle")
	}
}

func main() {
	rect := rect{}
	rect.height = 2
	rect.width = 5

	fmt.Println(rect.area())
	measure(&rect)
	circle := circle{}
	detectCircle(&circle)

	detectCircle(&rect)
}

/*

 */
