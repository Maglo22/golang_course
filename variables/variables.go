package main

import (
	"fmt"
	"strconv"
)

// Scopes:
// lower case first letter -> scoped to only the package
// upper case first letter -> exported from package

// Naming conventions:
// Pascal or camelCase
// small name -> short variable usage (~immediately)
// longer name -> long variable usage (throughout the code)
// acronyms in upper case (Url -> URL, Http -> HTTP, etc.)

// variable at package level
var i int = 41

// declare block of variables
var (
	firstName string = "Boi"
	lastName  string = "Aso"
	age       int    = 21
)

// variables not being used are considered
// errors to the compiler

func main() {
	// declare -> var name type
	// same name, different scope
	var i int
	// assign value
	i = 42

	// initialize and assign
	var j int = 43

	// let the compiler figure out the type
	k := 44

	// i = 42 because of shadowing
	fmt.Println(i + j + k)

	// type conversions are explicit only
	var flt float32
	flt = float32(i)
	fmt.Printf("%v, %T\n", flt, flt)

	// convert int to string
	// string(i) would result in the character with
	// the UNICODE value of i (42)
	var str string
	str = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", str, str)
}
