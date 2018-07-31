// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
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
	// Path of the request
	path := req.URL.Path
	if strings.HasSuffix(path, "@gmail.com") {
		name := strings.TrimSuffix(path, "@gmail.com")
		_, err := fmt.Fprintf(w, "Hello, gmail user %s\n", name[1:])
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	_, err := fmt.Fprintf(w, "Hello %s\n", path)
	if err != nil {
		log.Fatal(err)
	}
}
