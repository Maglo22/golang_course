package main

import (
	"fmt"
)

// a slice is a projection of an array; they are backed by an array
// almost all the same functions of arrays are available to slices
// slices don't have to have a fixed size

func main() {
	// slice, not an array
	a := []int{1, 2, 3}

	fmt.Println(a) // [1 2 3]

	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// slices are naturally referenced values
	b := a
	// this changes the value in a
	b[1] = 5
	fmt.Println(a) // [1 5 3]

	// other ways to create slices
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice of all elements
	all := numbers[:]
	// slice from 4th element to end
	fourth := numbers[3:]
	// slice of first 6 elements
	six := numbers[:6]
	// slice from the 4th to the 6th
	// first number -> inclusive; second number -> exclusive
	fs := numbers[3:6]

	fmt.Println(all)
	fmt.Println(fourth)
	fmt.Println(six)
	fmt.Println(fs)

	// using make to create a slice
	// type, length, capacity
	slice := make([]int, 3, 100)

	fmt.Println(slice)
	fmt.Printf("Length: %v\n", len(slice))
	fmt.Printf("Capacity: %v\n", cap(slice))

	// concatenating slices with spreading
	s1 := []int{}
	s1 = append(s1, []int{1, 2, 3, 4}...)
}
