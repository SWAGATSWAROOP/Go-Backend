package main

import "fmt"

// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

func increment() func() int {
	var a int = 0
	return func() int {
		a++
		return a
	}
}

func main() {
	inc := increment()
	fmt.Println(inc())

	// inc()
	println(inc())
}
