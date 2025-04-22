package main

import (
	"errors"
	"fmt"
)

// In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = fmt.Errorf("No more tea available")
var ErrPower = fmt.Errorf("Can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

/*
We can wrap errors with higher-level errors to add context. The simplest way to do this is with the %w verb in fmt.Errorf. Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
*/

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}
	// errors.Is checks that a given error (or any error in its chain) matches a specific error value. This is especially useful with wrapped or nested errors, allowing you to identify specific error types or sentinel errors in a chain of errors.
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy a new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Println("Unknown error : ", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}
