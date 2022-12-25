// Go program to illustrate
// how to trim a string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "!!Welcome to GeeksforGeeks !!"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using Trim() function
	res1 := strings.Trim(str1, "!")
	res2 := strings.Trim(str2, "@$")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}

--------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// trim left-hand side elements
// from the string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "!!Welcome to GeeksforGeeks **"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using TrimLeft() function
	res1 := strings.TrimLeft(str1, "!*")
	res2 := strings.TrimLeft(str2, "@")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}

-------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// trim right-hand side elements
// from the string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing the
	// string using shorthand declaration
	str1 := "!!Welcome to GeeksforGeeks **"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using TrimRight() function
	res1 := strings.TrimRight(str1, "!*")
	res2 := strings.TrimRight(str2, "$")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}

--------------------------------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// trim white space from the string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := " **Welcome to GeeksforGeeks** "
	str2 := " ##This is the tutorial of Golang## "

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println(str1, str2)

	// Trimming white space from the given strings
	// Using TrimSpace() function
	res1 := strings.TrimSpace(str1)
	res2 := strings.TrimSpace(str2)

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println(res1, res2)
}

------------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// trim a suffix string from the
// given string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "Welcome, GeeksforGeeks"
	str2 := "This is the, tutorial of Golang"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming suffix string from the given strings
	// Using TrimSuffix() function
	res1 := strings.TrimSuffix(str1, "GeeksforGeeks")
	res2 := strings.TrimSuffix(str2, "Hello")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}


-------------------------------------------------------------------------------------------------



// Go program to illustrate how to
// trim the prefix string from the
// given string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "Welcome, GeeksforGeeks"
	str2 := "This is the, tutorial of Golang"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)

	// Trimming prefix string from the given strings
	// Using TrimPrefix() function
	res1 := strings.TrimPrefix(str1, "Welcome")
	res2 := strings.TrimPrefix(str2, "Hello")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
}

--------------------------------------------------------------------------------------------------