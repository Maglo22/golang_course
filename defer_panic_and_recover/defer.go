package main

import (
	"fmt"
)

// defer calls are executed in LIFO order,
// the last defer in the code is called first

func main() {
	fmt.Println("start")
	// call executed after main function finishes,
	// but before it returns
	defer fmt.Println("middle")
	fmt.Println("end")

	// defer takes the current values of the arguments
	// passed when called, ignores if they are changed later
	a := "start"
	defer fmt.Println(a) // prints start
	a = "end"
}
