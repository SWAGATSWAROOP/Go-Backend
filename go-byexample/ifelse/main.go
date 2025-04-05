package main

import "fmt"

func main() {

	if num := 9; num > 0 {
		fmt.Println(num, "is postive")
	}
	//println(n) //n cannot be accessed
}

// A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
