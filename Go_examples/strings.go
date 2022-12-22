// Go program to illustrate
// how to create strings
package main

import "fmt"

func main() {

	// Creating and initializing a
	// variable with a string
	// Using shorthand declaration
	My_value_1 := "Welcome to GeeksforGeeks"

	// Using var keyword
	var My_value_2 string
	My_value_2 = "GeeksforGeeks"

	// Displaying strings
	fmt.Println("String 1: ", My_value_1)
	fmt.Println("String 2: ", My_value_2)
}
------------------------------------------------------------------------------------------------------------------------------------

// Go program to illustrate string literals
package main

import "fmt"

func main() {

	// Creating and initializing a
	// variable with a string literal
	// Using double-quote
	My_value_1 := "Welcome to GeeksforGeeks"

	// Adding escape character
	My_value_2 := "Welcome!\nGeeksforGeeks"

	// Using backticks
	My_value_3 := `Hello!GeeksforGeeks`

	// Adding escape character
	// in raw literals
	My_value_4 := `Hello!\nGeeksforGeeks`

	// Displaying strings
	fmt.Println("String 1: ", My_value_1)
	fmt.Println("String 2: ", My_value_2)
	fmt.Println("String 3: ", My_value_3)
	fmt.Println("String 4: ", My_value_4)
}

------------------------------------------------------------------------------------------------------------

// Go program to illustrate
// string are immutable
package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a string
	// using shorthand declaration
	mystr := "Welcome to GeeksforGeeks"

	fmt.Println("String:", mystr)

	/* if you trying to change
		the value of the string
		then the compiler will
		throw an error, i.e,
		cannot assign to mystr[1]

	mystr[1]= 'G'
	fmt.Println("String:", mystr)
	*/

}


--------------------------------------------------------------------------------------------------------------------------

// Go program to illustrate how
// to iterate over the string
// using for range loop
package main

import "fmt"

// Main function
func main() {

	// String as a range in the for loop
	for index, s := range "GeeksForGeeKs" {
	
		fmt.Printf("The index number of %c is %d\n", s, index)
	}
}


--------------------------------------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// access the bytes of the string
package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a string
	str := "Welcome to GeeksforGeeks"

	// Accessing the bytes of the given string
	for c := 0; c < len(str); c++ {

		fmt.Printf("\nCharacter = %c Bytes = %v", str, str)
	}
}

---------------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// create a string from the slice

package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a slice of byte
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}

	// Creating a string from the slice
	mystring1 := string(myslice1)

	// Displaying the string
	fmt.Println("String 1: ", mystring1)

	// Creating and initializing a slice of rune
	myslice2 := []rune{0x0047, 0x0065, 0x0065,0x006b, 0x0073}

	// Creating a string from the slice
	mystring2 := string(myslice2)

	// Displaying the string
	fmt.Println("String 2: ", mystring2)
}

---------------------------------------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// find the length of the string

package main

import (
	"fmt"
	"unicode/utf8"
)

// Main function
func main() {

	// Creating and initializing a string
	// using shorthand declaration
	mystr := "Welcome to GeeksforGeeks ??????"

	// Finding the length of the string
	// Using len() function
	length1 := len(mystr)

	// Using RuneCountInString() function
	length2 := utf8.RuneCountInString(mystr)

	// Displaying the length of the string
	fmt.Println("string:", mystr)
	fmt.Println("Length 1:", length1)
	fmt.Println("Length 2:", length2)

}

------------------------------------------------------------------------------------------------------------------

