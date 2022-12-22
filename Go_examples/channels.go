// Go program to illustrate
// how to create a channel
package main

import "fmt"

func main() {

	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)
}

-------------------------------------------------------------------------------------------------------------

// Go program to illustrate send
// and receive operation
package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}

---------------------------------------------------------------------------------------------------------------

// Go program to illustrate how
// to close a channel using for
// range loop and close function
package main

import "fmt"

// Function
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

// Main function
func main() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}


---------------------------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// use for loop in the channel

package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string)

	// Anonymous goroutine
	go func() {
		mychnl <- "ABH"
		mychnl <- "abh"
		mychnl <- "abhis"
		mychnl <- "Abhicodeits"
		close(mychnl)
	}()

	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
	}
}


----------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// find the length of the channel

package main

import "fmt"
a
// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 4)
	mychnl <- "ABHI"
	mychnl <- "ACIO"
	mychnl <- "Abhis"
	mychnl <- "Abhicodeitout"

	// Finding the length of the channel
	// Using len() function
	fmt.Println("Length of the channel is: ", len(mychnl))
}

----------------------------------------------------------------------------------------------------

// Go program to illustrate how to
// find the capacity of the channel

package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 5)
	mychnl <- "ABH"
	mychnl <- "abhi"
	mychnl <- "abhis"
	mychnl <- "Abhicodeitout"

	// Finding the capacity of the channel
	// Using cap() function
	fmt.Println("Capacity of the channel is: ", cap(mychnl))
}

----------------------------------------------------------------------------------------------------

// Go program to illustrate the concept
// of the unidirectional channel
package main

import "fmt"

// Main function
func main() {

	// Only for receiving
	mychanl1 := make(<-chan string)

	// Only for sending
	mychanl2 := make(chan<- string)

	// Display the types of channels
	fmt.Printf("%T", mychanl1)
	fmt.Printf("\n%T", mychanl2)
}


-----------------------------------------------------------------------------------------------------------

// Go program to illustrate how to convert
// bidirectional channel into the
// unidirectional channel
package main

import "fmt"

func sending(s chan<- string) {
	s <- "GeeksforGeeks"
}

func main() {

	// Creating a bidirectional channel
	mychanl := make(chan string)

	// Here, sending() function convert
	// the bidirectional channel
	// into send only channel
	go sending(mychanl)

	// Here, the channel is sent
	// only inside the goroutine
	// outside the goroutine the
	// channel is bidirectional
	// So, it print GeeksforGeeks
	fmt.Println(<-mychanl)
}
