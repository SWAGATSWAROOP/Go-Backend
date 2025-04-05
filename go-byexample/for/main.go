package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	//Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for i := range 3 {
		fmt.Println("Range ", i)
	}

	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
