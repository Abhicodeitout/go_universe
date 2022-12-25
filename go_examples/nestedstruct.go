// Golang program to illustrate
// the nested structure
package main

import "fmt"

// Creating structure
type Student struct {
	name   string
	branch string
	year   int
}

// Creating nested structure
type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

func main() {

	// Initializing the fields
	// of the structure
	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{"Bongo", "CSE", 2},
	}

	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", result.details.name)
	fmt.Println("Student's branch name: ", result.details.branch)
	fmt.Println("Year: ", result.details.year)
}
