// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays
package main

import (
	"fmt"
)

func main() {

	//To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
	var s []string
	fmt.Println(s, len(s), cap(s))

	//
	s = make([]string, 3)
	fmt.Println("length: ", len(s), "Capacity: ", cap(s))
}
