package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Sometimes we’ll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort strings by their length instead of alphabetically. Here’s an example of custom sorts in Go.

func main() {
	fruits := []string{"grape", "orange", "apple"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)

	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	person := []Person{
		Person{"Bob", 22},
		Person{"Alice", 23},
		Person{"Alice", 20},
	}

	slices.SortFunc(person, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	//Note: if the Person struct is large, you may want the slice to contain *Person instead and adjust the sorting function accordingly. If in doubt, benchmark!

	fmt.Println(person)
}
