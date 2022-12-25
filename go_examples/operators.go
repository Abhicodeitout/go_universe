// Go program to illustrate the
// use of arithmetic operators
package main

import "fmt"

func main() {
p:= 34
q:= 20
	
// Addition
result1:= p + q
fmt.Printf("Result of p + q = %d", result1)
	
// Subtraction
result2:= p - q
fmt.Printf("\nResult of p - q = %d", result2)
	
// Multiplication
result3:= p * q
fmt.Printf("\nResult of p * q = %d", result3)
	
// Division
result4:= p / q
fmt.Printf("\nResult of p / q = %d", result4)
	
// Modulus
result5:= p % q
fmt.Printf("\nResult of p %% q = %d", result5)
}

--------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of relational operators
package main

import "fmt"

func main() {
p:= 34
q:= 20
	
// ‘=='(Equal To)
result1:= p == q
fmt.Println(result1)
	
// ‘!='(Not Equal To)
result2:= p != q
fmt.Println(result2)
	
// ‘<‘(Less Than)
result3:= p < q
fmt.Println(result3)
	
// ‘>'(Greater Than)
result4:= p > q
fmt.Println(result4)
	
// ‘>='(Greater Than Equal To)
result5:= p >= q
fmt.Println(result5)
	
// ‘<='(Less Than Equal To)
result6:= p <= q
fmt.Println(result6)
	
	
}

---------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of logical operators
package main
import "fmt"
func main() {
	var p int = 23
	var q int = 60
	
	if(p!=q && p<=q){
		fmt.Println("True")
	}
	
	if(p!=q || p<=q){
		fmt.Println("True")
	}
	
	if(!(p==q)){
		fmt.Println("True")
	}
	
}

------------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of bitwise operators
package main

import "fmt"

func main() {
p:= 34
q:= 20
	
// & (bitwise AND)
result1:= p & q
fmt.Printf("Result of p & q = %d", result1)
	
// | (bitwise OR)
result2:= p | q
fmt.Printf("\nResult of p | q = %d", result2)
	
// ^ (bitwise XOR)
result3:= p ^ q
fmt.Printf("\nResult of p ^ q = %d", result3)
	
// << (left shift)
result4:= p << 1
fmt.Printf("\nResult of p << 1 = %d", result4)
	
// >> (right shift)
result5:= p >> 1
fmt.Printf("\nResult of p >> 1 = %d", result5)
	
// &^ (AND NOT)
result6:= p &^ q
fmt.Printf("\nResult of p &^ q = %d", result6)
	
	
}

--------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of assignment operators
package main

import "fmt"

func main() {
var p int = 45
	var q int = 50
	
// “=”(Simple Assignment)
p = q
fmt.Println(p)
	
// “+=”(Add Assignment)
	p += q
fmt.Println(p)
	
//“-=”(Subtract Assignment)
p-=q
fmt.Println(p)
	
// “*=”(Multiply Assignment)
p*= q
fmt.Println(p)
	
// “/=”(Division Assignment)
	p /= q
fmt.Println(p)
	
	// “%=”(Modulus Assignment)
	p %= q
fmt.Println(p)
	
}

---------------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of Misc Operators
package main

import "fmt"

func main() {
a := 4
	
// Using address of operator(&) and
// pointer indirection(*) operator
b := &a
fmt.Println(*b)
*b = 7
fmt.Println(a)
}

---------------------------------------------------------------------------------------------------

