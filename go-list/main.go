// go list -f '{{ .Name }}' will show the package name
// go list -f '{{ .Name }}: {{ .Doc }}' will show the package name and documentation
// go list -f '{{ .Imports }}' will show the dependencies
// go list -f '{{ .Imports }}' <package> will show the dependencies of <package>
// go doc <package> <function> will show documentations of <function> in <package>
// godoc -http :<port> will show documentations of installed packages using localhost:<port>
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
}
