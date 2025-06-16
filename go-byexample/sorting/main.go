package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	strs := []string{"a", "d", "c"}
	sort.Strings(strs)
	//slices.Sort(strs)
	fmt.Println(strs)

	arrayOfint := []int{2, 4, 5, -1, 32, 32, 246, -19}
	slices.Sort(arrayOfint)
	fmt.Println(arrayOfint)

	fmt.Println(slices.IsSorted(arrayOfint))
}
