package main

import "os"

// A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

func main() {

	//panic("A problem Occured")

	_, err := os.Create("/tmp/test.txt")
	//Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
	if err != nil {
		panic(err)
	}
	// Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status. When first panic in main fires, the program exits without reaching the rest of the code. If you’d like to see the program try to create a temp file, comment the first panic out.
}
