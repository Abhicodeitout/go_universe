// Go program to illustrate how to create
// an array using shorthand declaration
// and accessing the elements of the
// array using for loop
package main

import "fmt"

func main() {

// Shorthand declaration of array
arr:= [4]string{"abhi", "aco", "Abhicodeitout1231", "abhicodeitouts"}

// Accessing the elements of
// the array Using for loop
fmt.Println("Elements of the array:")

for i:= 0; i < 3; i++{
fmt.Println(arr[i])
}

}

-----------------------------------------------------------------------------------------------

// Go program to illustrate the
// concept of multi-dimension array
package main

import "fmt"

func main() {

	// Creating and initializing
	// 2-dimensional array
	// Using shorthand declaration
	// Here the (,) Comma is necessary
	arr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}

	// Accessing the values of the
	// array Using for loop
	fmt.Println("Elements of Array 1")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {

			fmt.Println(arr[x][y])
		}
	}

	// Creating a 2-dimensional
	// array using var keyword
	// and initializing a multi
	// -dimensional array using index
	var arr1 [2][2]int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	// Accessing the values of the array
	fmt.Println("Elements of array 2")
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr1[p][q])
		}
	}
}


--------------------------------------------------------------------------------------------------

// Go program to illustrate an array
package main

	import "fmt"

func main()
{

	// Creating an array of int type
	// which stores, two elements
	// Here, we do not initialize the
	// array so the value of the array
	// is zero
	var myarr[2] int fmt.Println("Elements of the Array: ", myarr)
}


-------------------------------------------------------------------------------------------------------

// Go program to illustrate how to find
// the length of the array
package main

import "fmt"

func main() {

// Creating array
// Using shorthand declaration	
arr1:= [3]int{9,7,6}
arr2:= [...]int{9,7,6,4,5,3,2,4}
arr3:= [3]int{9,3,5}

// Finding the length of the
// array using len method
fmt.Println("Length of the array 1 is:", len(arr1))
fmt.Println("Length of the array 2 is:", len(arr2))
fmt.Println("Length of the array 3 is:", len(arr3))
}

-----------------------------------------------------------------------------------------

// Go program to illustrate the
// concept of ellipsis in an array
package main

import "fmt"

func main() {
	
// Creating an array whose size is determined
// by the number of elements present in it
// Using ellipsis
myarray:= [...]string{"ABHI", "acio", "abhi","abhicodeitout", "ABHICODEITOUT"}

fmt.Println("Elements of the array: ", myarray)

// Length of the array
// is determine by
// Using len() method
fmt.Println("Length of the array is:", len(myarray))
}

---------------------------------------------------------------------------------------------------

// Go program to illustrate how to pass an
// array as an argument in the function
package main

import "fmt"

// This function accept
// an array as an argument
func myfun(a [6]int, size int) int {
	var k, val, r int

	for k = 0; k < size; k++ {
		val += a[k]
	}

	r = val / size
	return r
}

// Main function
func main() {

	// Creating and initializing an array
	var arr = [6]int{67, 59, 29, 35, 4, 34}
	var res int

	// Passing an array as an argument
	res = myfun(arr, 6)
	fmt.Printf("Final result is: %d ", res)
}

