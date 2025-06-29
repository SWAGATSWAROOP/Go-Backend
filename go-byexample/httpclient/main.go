package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	//Issue an HTTP GET request to a server. http.Get is a convenient shortcut around creating an http.Client object and calling its Get method; it uses the http.DefaultClient object which has useful default settings.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	/*
		    is used to ensure the HTTP response body is properly closed once you are done reading from it, even if an error or early return occurs later in the function.
		    Why is it needed?
		    When you make an HTTP request using http.Get, the server sends back a response, and the Body is a stream that needs to be closed manually to:
		    Free up resources (like file descriptors and memory)
		    Avoid connection leaks — Go reuses HTTP connections by default (keep-alive).
		    If you don't close the response body, the connection won't be reused,
			and the underlying TCP connection will remain open, eventually exhausting available connections.
		    Allow the HTTP client's connection pool to work efficiently
	*/

	fmt.Println("Response Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
