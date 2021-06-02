package main

import (
	"fmt"
)

// arrays allocate their elements in continous segments of memory
// they have fixed size

func main() {
	// array of 3 int
	// ... -> large enough for the data initialized with
	grades := [...]int{97, 85, 93}

	// array with no elements
	var students [3]string

	// assign elements to array
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Lmao"

	// get array length
	arrLen := len(students)

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", arrLen)

	// matrix
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	// arrays are considered values in go
	// copying an array -> b is a literal copy, not a reference
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5

	fmt.Print(a) // [1 2 3]
	fmt.Print(b) // [1 5 3]

	// use pointers to reference the original instead of copying
	c := &a
	fmt.Print(c) // &[1 2 3]
}
