// Go program to illustrate the concept
// of == and != operator with strings
package main

import "fmt"

// Main function
func main() {

	// Creating and initializing strings
	// using shorthand declaration
	str1 := "abhis"
	str2 := "abhish"
	str3 := "abhisforabhis"
	str4 := "abhis"

	// Checking the string are equal
	// or not using == operator
	result1 := str1 == str2
	result2 := str2 == str3
	result3 := str3 == str4
	result4 := str1 == str4
	
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
	fmt.Println("Result 4: ", result4)

	// Checking the string are not equal
	// using != operator
	result5 := str1 != str2
	result6 := str2 != str3
	result7 := str3 != str4
	result8 := str1 != str4
	
	fmt.Println("\nResult 5: ", result5)
	fmt.Println("Result 6: ", result6)
	fmt.Println("Result 7: ", result7)
	fmt.Println("Result 8: ", result8)

}

----------------------------------------------------------------------------------------------------

// Go program to illustrate the concept
// of comparison operator with strings
package main

import "fmt"

// Main function
func main() {

	// Creating and initializing
	// slice of string using the
	// shorthand declaration
	myslice := []string{"abhis", "abhis",
					"abhi", "abhi", "for"}
	
	fmt.Println("Slice: ", myslice)

	// Using comparison operator
	result1 := "abhi" > "abhis"
	fmt.Println("Result 1: ", result1)

	result2 := "abhi" < "abhis"
	fmt.Println("Result 2: ", result2)

	result3 := "abhis" >= "for"
	fmt.Println("Result 3: ", result3)

	result4 := "abhis" <= "for"
	fmt.Println("Result 4: ", result4)

	result5 := "abhis" == "abhis"
	fmt.Println("Result 5: ", result5)

	result6 := "abhis" != "for"
	fmt.Println("Result 6: ", result6)

}

-----------------------------------------------------------------------------------------------------

// Go program to illustrate how to compare
// string using compare() function
package main

import (
	"fmt"
	"strings"
)

func main() {

	// Comparing string using Compare function
	fmt.Println(strings.Compare("abh", "abhishek"))
	
	fmt.Println(strings.Compare("abhisheks","abhishek"))
	
	fmt.Println(strings.Compare("abhishek", "ABH"))
	
	fmt.Println(strings.Compare("abhis", "abhis"))

}


------------------------------------------------------------------------------------------------------