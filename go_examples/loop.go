// Go program to illustrate the
// use of simple for loop
package main

import "fmt"

// Main function
func main() {
	
	// for loop
	// This loop starts when i = 0
	// executes till i<4 condition is true
	// post statement is i++
	for i := 0; i < 4; i++{
	fmt.Printf("GeeksforGeeks\n")
	}
	
}

------------------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of an infinite loop
package main

import "fmt"

// Main function
func main() {
	
	// Infinite loop
	for {
	fmt.Printf("GeeksforGeeks\n")
	}
	
}

---------------------------------------------------------------------------------------------------

// Go program to illustrate the
// the for loop as while Loop
package main

import "fmt"

// Main function
func main() {
	
	// while loop
	// for loop executes till
	// i < 3 condition is true
	i:= 0
	for i < 3 {
	i += 2
	}
fmt.Println(i)
}

-----------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of simple range loop
package main

import "fmt"

// Main function
func main() {
	
	// Here rvariable is a array
	rvariable:= []string{"ABH", "abhis", "abhicodeitout"}
	
	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j:= range rvariable {
	fmt.Println(i, j)
	}
	
}

-------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use for loop using string
package main

import "fmt"

// Main function
func main() {
	
	// String as a range in the for loop
	for i, j:= range "XabCd" {
	fmt.Printf("The index number of %U is %d\n", j, i)
	}
	
}

-------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use for loop using maps
package main

import "fmt"

// Main function
func main() {
	
	// using maps
	mmap := map[int]string{
		22:"abhis",
		33:"ABH",
		44:"Abhicodeitout",
	}
	for key, value:= range mmap {
	fmt.Println(key, value)
	}
	
}

---------------------------------------------------------------------------------------------

// Go program to illustrate the
// use for loop using channel
package main

import "fmt"

// Main function
func main() {
	
	// using channel
	chnl := make(chan int)
	go func(){
		chnl <- 100
		chnl <- 1000
	chnl <- 10000
	chnl <- 100000
	close(chnl)
	}()
	for i:= range chnl {
	fmt.Println(i)
	}
	
}

----------------------------------------------------------------------------------------------

