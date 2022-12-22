// Golang program to demonstrate the variables
// storing the hexadecimal values
package main

import "fmt"

func main() {

	// storing the hexadecimal
	// values in variables
	x := 0xFF
	y := 0x9C
	
	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)
	
	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)
	
}

-------------------------------------------------------------------------------------------------

// Golang program to demonstrate the declaration
// and initialization of pointers
package main

import "fmt"

func main() {

	// taking a normal variable
	var x int = 5748
	
	// declaration of pointer
	var p *int
	
	// initialization of pointer
	p = &x

		// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
}

-----------------------------------------------------------------------------------------------------------

// Golang program to demonstrate
// the use of type inference in
// Pointer variables
package main

import "fmt"

func main() {

	// using var keyword
	// we are not defining
	// any type with variable
	var y = 458
	
	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}

----------------------------------------------------------------------------------------------------

// Golang program to demonstrate
// the use of shorthand syntax in
// Pointer variables
package main

import "fmt"

func main() {

	// using := operator to declare
	// and initialize the variable
	y := 458
	
	// taking a pointer variable using
	// := by assigning it with the
	// address of variable y
	p := &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}


------------------------------------------------------------------------------------------------------

// Golang program to illustrate the
// concept of dereferencing a pointer
package main

import "fmt"

func main() {

	// using var keyword
	// we are not defining
	// any type with variable
	var y = 458
	
	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in y(*p) = ", *p)

}

--------------------------------------------------------------------------------------------

// Golang program to illustrate the
// above mentioned concept
package main

import "fmt"

func main() {

	// using var keyword
	// we are not defining
	// any type with variable
	var y = 458
	
	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y before changing = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in y(*p) Before Changing = ", *p)

	// changing the value of y by assigning
	// the new value to the pointer
	*p = 500

	fmt.Println("Value stored in y(*p) after Changing = ",y)

}
