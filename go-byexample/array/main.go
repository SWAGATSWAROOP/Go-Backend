package main

import "fmt"

//In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{2, 3, 20, 6, 4} //Compiler count the value
	fmt.Println(b)

	//If you specify the index with :, the elements in between will be zeroed.
	// b = [...]int{100, 2: 400} Cannot Do this wrong in documentation

	b = [5]int{1: 5} // If you specify the index with :, the elements in between will be zeroed. Works in between only
	fmt.Println(b)

	// c := [...]int{2, 3}
	// println(c) We cannot do this

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	// The problem is that you cannot use the const keyword with an array literal.
	threeD := [2][3][1]int{
		{{1}, {1}, {1}},
		{{1}, {1}, {1}},
	}
	fmt.Println(threeD)

}
