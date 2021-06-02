package main

import (
	"fmt"
)

// a constant value can't be defined by something that resolves
// at runtime (functions).

// constants can be shadowed
const myConst int = 11

// enumerated constants
// allows operations that can be determined at compile time
// arithmetic, bitwise, bitshifting
// iota is scoped to this block
const (
	a = iota
	b
	c
)

// ignore the first value of iota (0)
const (
	_ = iota
	innitStatus
	runStatus
	endStatus
)

func main() {
	// naming convention:
	// camelCase, upper case first letter to export it

	// typed constant
	const myConst int = 12
	fmt.Printf("%v, %T\n", myConst, myConst)

	// inferred constant
	// allows to do operations mixing (similar) types
	const con = 42
	var v int16 = 10
	fmt.Printf("%v, %T\n", con+v, con+v)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
