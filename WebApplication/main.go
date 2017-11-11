// author: Jamie O'Sullivan
// ChatBot
//

package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"fmt"
	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	http.HandleFunc("/", handler)

	// After defining the server, "listen and serve" on port 8080
	http.ListenAndServe(":8080", nil)
}

// It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
