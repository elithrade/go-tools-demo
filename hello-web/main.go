// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// If the return error is not handled
	// the program will fail silently
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "Hello web")
	if err != nil {
		log.Fatal(err)
	}
}
