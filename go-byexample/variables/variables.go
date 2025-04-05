package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b int
	fmt.Println(b)
}

//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
