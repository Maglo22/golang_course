package main

import (
	"fmt"
)

// zero value of a pointer is <nil>
// pointer arithmetic is considered unsafe,
// to do it, import unsafe package

func main() {
	var a int = 42
	// b points to memory address of a
	var b *int = &a
	fmt.Println(a, b)
	// dereference to get value in memory address
	*b = 14
	fmt.Println(a, *b)

	arr := [3]int{1, 2, 3}
	e1 := &arr[0]
	e2 := &arr[1]
	fmt.Printf("%v %p %p", arr, e1, e2)

	// pointer to a struct
	var ms *myStruct // <nil>
	ms = new(myStruct)
	// dereference is done by the compiler
	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
