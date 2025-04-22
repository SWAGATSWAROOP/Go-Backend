package main

import (
	"errors"
	"fmt"
)

// It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

// A custom error type usually has the suffix “Error”.

type argError struct {
	arg     int
	message string
}

// Adding this Error method makes argError implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
