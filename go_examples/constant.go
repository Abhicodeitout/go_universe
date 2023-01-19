package main

import "fmt"

const PI = 3.14

func main()
{
	const ABH = "Abhicodeitout"
	fmt.Println("Hello", ABH)

	fmt.Println("Happy", PI, "Day")

	const Correct= true
	fmt.Println("Go rules?", Correct)
}

----------------------------------------------------------------------------------------------

package main

import "fmt"

func main()
{
	const A = "ABH"
	var B = "abhicodeitout"
	
	// Concat strings.
	var helloWorld = A+ " " + B
	helloWorld += "!"
	fmt.Println(helloWorld)
	
	// Compare strings.
	fmt.Println(A == "ABH")
	fmt.Println(B < A)
}


--------------------------------------------------------------------------------------------------

package main

import "fmt"

const Pi = 3.14

func main()
{
	const trueConst = true
	
	// Type definition using type keyword
	type myBool bool
	var defaultBool = trueConst // allowed
	var customBool myBool = trueConst // allowed
	
	// defaultBool = customBool // not allowed
	fmt.Println(defaultBool)
	fmt.Println(customBool)
}

---------------------------------------------------------------------------------------------------

// Const demonstration using go.
package main

import (
	"fmt"
	"math"
)

const s string = "GeeksForGeeks"

func main() {
	fmt.Println(s)

	const n = 5

	const d = 3e10 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

----------------------------------------------------------------------------------------------------

