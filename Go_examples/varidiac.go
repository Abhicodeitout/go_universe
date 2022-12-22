// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {

	// zero argument
	fmt.Println(joinstr())

	// multiple arguments
	fmt.Println(joinstr("abhis", "ABH"))
	fmt.Println(joinstr("abhiss", "for", "abhiss"))
	fmt.Println(joinstr("A", "B", "b", "H", "I"))

}

--------------------------------------------------------------------------------------------------------------

// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	// pass a slice in variadic function
	elements := []string{"abhiss", "FOR", "abhiss"}
	fmt.Println(joinstr(elements...))
}

