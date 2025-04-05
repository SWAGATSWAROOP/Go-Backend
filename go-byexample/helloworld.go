package main

import "fmt"

func main() {
	a, err := fmt.Println("Hello World") //Returns integer and error
	if err == nil {
		// fmt.Print("Hello")
		fmt.Print(a)
	}
}
