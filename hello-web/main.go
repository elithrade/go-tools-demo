// Package main provides ...
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	// Nothing will happen, no errors will show
	// we can run errcheck to show errors
	fmt.Fprintln(w, "Hello web")
}
