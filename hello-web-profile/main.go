// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
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
	// regex used to match string starts with anything ends with gmail.com
	re := regexp.MustCompile("^(.+)@gmail.com$")
	// Path of the request
	path := req.URL.Path
	match := re.FindAllStringSubmatch(path[1:], -1)
	if match != nil {
		// We have a match
		_, err := fmt.Fprintf(w, "Hello, gmail user %s\n", match[0][1])
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
