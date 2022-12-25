// Golang program to show the compiler
// throws an error if a variable is
// declared but not used

package main

import "fmt"

// Main function
func main() {

	// calling the function
	// function returns two values which are
	// assigned to mul and div identifier
	mul, div := mul_div(105, 7)

	// only using the mul variable
	// compiler will give an error
	fmt.Println("105 x 7 = ", mul)
}

// function returning two
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2
}

-------------------------------------------------------------------------------------------------------

// Golang program to the use of Blank identifier

package main

import "fmt"

// Main function
func main() {

	// calling the function
	// function returns two values which are
	// assigned to mul and blank identifier
	mul, _ := mul_div(105, 7)

	// only using the mul variable
	fmt.Println("105 x 7 = ", mul)
}

// function returning two
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2
}


------------------------------------------------------------------------------------------------

