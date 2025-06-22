package main

import (
	"fmt"
	"os"
)

//Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

func main() {
	f := createFile("/tmp/test.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("creating", path)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing", f.Name())
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing", f.Name())
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
