// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

// Variable declared and
// initialized without the
// explicit type
var myvariable1 = 20
var myvariable2 = "abhissforabhiss"
var myvariable3 = 34.80

// Display the value and the
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",myvariable1)
									
fmt.Printf("The type of myvariable1 is : %T\n",myvariable1)
	
fmt.Printf("The value of myvariable2 is : %s\n",myvariable2)
										
fmt.Printf("The type of myvariable2 is : %T\n",myvariable2)
	
fmt.Printf("The value of myvariable3 is : %f\n",myvariable3)
										
fmt.Printf("The type of myvariable3 is : %T\n",myvariable3)
	
}

-------------------------------------------------------------------------------------------------

// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

	// Variable declared and
	// initialized without expression
	var myvariable1 int
	var myvariable2 string
	var myvariable3 float64

	// Display the zero-value of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",myvariable2)

	fmt.Printf("The value of myvariable3 is : %f",myvariable3)
}

-----------------------------------------------------------------------------------------------------

// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

	// Multiple variables of the same type
	// are declared and initialized
	// in the single line
	var myvariable1, myvariable2, myvariable3 int = 2, 454, 67

// Display the values of the variables
fmt.Printf("The value of myvariable1 is : %d\n",myvariable1)

fmt.Printf("The value of myvariable2 is : %d\n",myvariable2)

fmt.Printf("The value of myvariable3 is : %d",myvariable3)
}

---------------------------------------------------------------------------------------------------

// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Multiple variables of different types
// are declared and initialized in the single line
var myvariable1, myvariable2, myvariable3 = 2, "abhi", 67.56

// Display the value and
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",myvariable1)

fmt.Printf("The type of myvariable1 is : %T\n",myvariable1)

fmt.Printf("\nThe value of myvariable2 is : %s\n",myvariable2)

fmt.Printf("The type of myvariable2 is : %T\n",myvariable2)

fmt.Printf("\nThe value of myvariable3 is : %f\n",myvariable3)

fmt.Printf("The type of myvariable3 is : %T\n",myvariable3)
}

---------------------------------------------------------------------------------------------------

// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
myvar1 := 39
myvar2 := "abhissforabhiss"
myvar3 := 34.67

// Display the value and type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}

------------------------------------------------------------------------------------------------------

// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
// Multiple variables of same types
// are declared and initialized in
// the single line
myvar1, myvar2, myvar3 := 800, 34, 56

// Display the value and
// type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %d\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %d\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}


------------------------------------------------------------------------------------------------------

// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
// Here, short variable declaration acts
// as an assignment for myvar2 variable
// because same variable present in the same block
// so the value of myvar2 is changed from 45 to 100
myvar1, myvar2 := 39, 45
myvar3, myvar2 := 45, 100

// If you try to run the commented lines,
// then compiler will gives error because
// these variables are already defined
// myvar1, myvar2 := 43, 47
// myvar2:= 200

// Display the values of the variables
fmt.Printf("The value of myvar1 and myvar2 is : %d %d\n",myvar1, myvar2)
											
fmt.Printf("The value of myvar3 and myvar2 is : %d %d\n",myvar3, myvar2)
}


