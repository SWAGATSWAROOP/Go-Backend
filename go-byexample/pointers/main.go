package main

import "fmt"

func makeValueZero(num int) {
	fmt.Println(num)
	num = 0
}

func makeZeroByPointer(ptr *int) {
	fmt.Println(*ptr)
	fmt.Println("Pointer: ", ptr)
	*ptr = 0
}

func main() {
	fmt.Println("Printing the values")

	var num int = 5
	makeValueZero(num)

	fmt.Println("Value: ", num)

	makeZeroByPointer(&num)
	println("Now after Pointer Value: ", num)

}
