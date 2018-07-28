// go list -f '{{ .Name }}' will show the package name
// go list -f '{{ .Name }}: {{ .Doc }}' will show the package name and documentation
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
}
