package main

import (
	"fmt"
	"maps"
)

//Hashes and Dictionary in other languages
//To make map : - make(map[key-value] value-type)

func main() {
	m := make(map[string]int64)

	m["k1"] = 7
	m["k2"] = 8

	//Shows key value pairs
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v2 := m["k3"]
	fmt.Println(v2) //default value is provided

	//Deleting a key-value pair
	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)
	//The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.

	n := map[string]int{"foo": 1}
	fmt.Println(n)

	n2 := map[string]int{"foo": 1}

	if maps.Equal(n, n2) {
		fmt.Println("n == m")
	}
}
