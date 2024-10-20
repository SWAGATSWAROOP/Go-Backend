package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Scan(&name)
	fmt.Println("Your Name is: ", name)
	// For taking a line input and os.Stdin :- to get any type of input from os
	reader := bufio.NewReader(os.Stdin)
	//To keep track of where to stop
	name2, _ := reader.ReadString('\n')
	fmt.Println(name2)
}
