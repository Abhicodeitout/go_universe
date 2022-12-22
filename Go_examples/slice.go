// Golang program to illustrate
// the working of the slice components
package main

import "fmt"

func main() {

	// Creating an array
	arr := [7]string{"This", "is", "the", "tutorial",
						"of", "Go", "language"}

	// Display array
	fmt.Println("Array:", arr)

	// Creating a slice
	myslice := arr[1:6]

	// Display slice
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}

-----------------------------------------------------------------------------------------------------

// Golang program to illustrate how
// to create a slice using a slice
// literal
package main

import "fmt"

func main() {

	// Creating a slice
	// using the var keyword
	var my_slice_1 = []string{"Geeks", "for", "Geeks"}

	fmt.Println("My Slice 1:", my_slice_1)

	// Creating a slice
	//using shorthand declaration
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
}

---------------------------------------------------------------------------------------------------------

// Golang program to illustrate how to
// create slices from the array
package main

import "fmt"

func main() {

	// Creating an array
	arr := [4]string{"abhishek", "abh", "abhish", "ABH"}

	// Creating slices from the given array
	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	// Display the result
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}

--------------------------------------------------------------------------------------------------------

// Golang program to illustrate how to
// create slices from the slice
package main

import "fmt"

func main() {

	// Creating s slice
	oRignAl_slice := []int{90, 60, 40, 50,
		34, 49, 30}

	// Creating slices from the given slice
	var my_slice_1 = oRignAl_slice[1:5]
	my_slice_2 := oRignAl_slice[0:]
	my_slice_3 := oRignAl_slice[:6]
	my_slice_4 := oRignAl_slice[:]
	my_slice_5 := my_slice_3[2:4]

	// Display the result
	fmt.Println("Original Slice:", oRignAl_slice)
	fmt.Println("New Slice 1:", my_slice_1)
	fmt.Println("New Slice 2:", my_slice_2)
	fmt.Println("New Slice 3:", my_slice_3)
	fmt.Println("New Slice 4:", my_slice_4)
	fmt.Println("New Slice 5:", my_slice_5)
}

---------------------------------------------------------------------------------------------------------

// Go program to illustrate how to create slices
// Using make function
package main

import "fmt"

func main() {

	// Creating an array of size 7
	// and slice this array till 4
	// and return the reference of the slice
	// Using make function
	var my_slice_1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
	my_slice_1, len(my_slice_1), cap(my_slice_1))

	// Creating another array of size 7
	// and return the reference of the slice
	// Using make function
	var my_slice_2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
	my_slice_2, len(my_slice_2), cap(my_slice_2))
	
}

--------------------------------------------------------------------------------------------------------

// Golang program to illustrate the
// iterating over a slice using
// for loop
package main

import "fmt"

func main() {

	// Creating a slice
	myslice := []string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}

	// Iterate using for loop
	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}
}

------------------------------------------------------------------------------------------------------------

// Golang program to illustrate the iterating
// over a slice using range in for loop
package main

import "fmt"

func main() {

	// Creating a slice
	myslice := []string{"This", "is", "the", "tutorial",
								"of", "Go", "language"}

	// Iterate slice
	// using range in for loop
	for index, ele := range myslice {
		fmt.Printf("Index = %d and element = %s\n", index+3, ele)
	}
}

-------------------------------------------------------------------------------------------------------------

// Golang program to illustrate the iterating over
// a slice using range in for loop without an index
package main

import "fmt"

func main() {

	// Creating a slice
	myslice := []string{"This", "is", "the",
		"tutorial", "of", "Go", "language"}

	// Iterate slice
	// using range in for loop
	// without index
	for _, ele := range myslice {
		fmt.Printf("Element = %s\n", ele)
	}
}

-----------------------------------------------------------------------------------------------------------------

// Golang program to illustrate
// how to modify the slice
package main

import "fmt"

func main() {

	// Creating a zero value slice
	arr := [6]int{55, 66, 77, 88, 99, 22}
	slc := arr[0:4]

	// Before modifying

	fmt.Println("Original_Array: ", arr)
	fmt.Println("Original_Slice: ", slc)

	// After modification
	slc[0] = 100
	slc[1] = 1000
	slc[2] = 1000

	fmt.Println("\nNew_Array: ", arr)
	fmt.Println("New_Slice: ", slc)
}


-----------------------------------------------------------------------------------------------------------------

// Go program to illustrate the multi-dimensional slice
package main

import "fmt"

func main() {

	// Creating multi-dimensional slice
	s1 := [][]int{{12, 34},
		{56, 47},
		{29, 40},
		{46, 78},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", s1)

	// Creating multi-dimensional slice
	s2 := [][]string{
		[]string{"abhish", "for"},
		[]string{"abhish", "ABH"},
		[]string{"abh", "abhish"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", s2)

}
-----------------------------------------------------------------------------------------------------------------------

// Go program to show the slice
// - composite literal
package main

import "fmt"

func main() {

	// Slice with composite literal
	// Slice allows you to group together
	// the values of the same type
	// here type of values is int
	s1 := []int{23, 56, 89, 34}
	
	// displaying the values
	fmt.Println(s1)
}

