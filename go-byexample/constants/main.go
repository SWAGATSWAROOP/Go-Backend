package main

import (
	"fmt"
	"math"
)

const s string = "Hello"

func main() {
	fmt.Println(s)
	const n = 1000
	fmt.Println(n)
	// const z := 20 Shorthand operator cannot be used in const
	const d = 3e20 / n
	fmt.Println(int64(d))
	//A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(math.Sin(n))
	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
}
