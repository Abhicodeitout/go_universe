// Go program to illustrate the
// use of if statement
package main

import "fmt"

func main() {
	
// taking a local variable
var v int = 700

// using if statement for
// checking the condition
if v < 1000 {
		
	// print the following if
	// condition evaluates to true
	fmt.Printf("v is less than 1000\n")
}
	
fmt.Printf("Value of v is : %d\n", v)
	
}

-----------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of if...else statement
package main

import "fmt"

func main() {
	
// taking a local variable
var v int = 1200

// using if statement for
// checking the condition
if v < 1000 {
		
	// print the following if
	// condition evaluates to true
	fmt.Printf("v is less than 1000\n")
	
} else {
		
	// print the following if
	// condition evaluates to true
	fmt.Printf("v is greater than 1000\n")
}
	
}

-------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of nested if statement
package main
import "fmt"

func main() {
	
// taking two local variable
var v1 int = 400
var v2 int = 700

// using if statement
if( v1 == 400 ) {
		
	// if condition is true then
	// check the following
	if( v2 == 700 ) {
		
		// if condition is true
		// then display the following
		fmt.Printf("Value of v1 is 400 and v2 is 700\n" );
	}
}

}


----------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of if..else..if ladder
package main
import "fmt"

func main() {
	
// taking a local variable
var v1 int = 700

// checking the condition
if v1 == 100 {
		
	// if condition is true then
	// display the following */
	fmt.Printf("Value of v1 is 100\n")
	
} else if v1 == 200 {
		
	fmt.Printf("Value of a is 20\n")
	
} else if v1 == 300 {
		
	fmt.Printf("Value of a is 300\n")
	
} else {
		
	// if none of the conditions is true
	fmt.Printf("None of the values is matching\n")
}
}

------------------------------------------------------------------------------------------------------

