package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{1, 2, 3}
	fmt.Println(arr)

	for index, value := range arr {
		fmt.Printf("%d this is index and value is %d\n", index, value)
	}

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details.
	name := "Swagat"
	for i, m := range name {
		fmt.Println(i, " ", m)
	}

	kvs := map[string]int{"Swagat": 1}

	for i, j := range kvs {
		println("key: ", i, " value : ", j)
	}
}
