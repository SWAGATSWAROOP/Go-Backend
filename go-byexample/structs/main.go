package main

import "fmt"

type person struct {
	name  string
	class int64
}

// Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.
func newPerson(name string, class int64) *person {
	p := person{name: name}
	p.class = class
	return &p
}

func main() {
	p := person{name: "Swagat"}
	fmt.Println(p)
	p.class = 4

	fmt.Println(p)
	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &p
	fmt.Println(sp)
	fmt.Println(sp.name)
	fmt.Println(newPerson("Swagat", 4))

	// If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
	dog := struct {
		name string
	}{
		"Tom",
	}
	fmt.Println(dog)
}
