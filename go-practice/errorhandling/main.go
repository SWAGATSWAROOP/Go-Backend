package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return a / b, fmt.Errorf("divide by 0 error")
	}
	return a / b, nil
}

func main() {
	var a float64 = 2.00
	var b = 9.00
	c := 20.00
	d, err := divide(a, b)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(d)
	k, err := divide(c, 0)
	if err != nil {
		fmt.Println("Divide by 0 error")
	}
	println(k)
}
