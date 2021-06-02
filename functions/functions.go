package main

import (
	"fmt"
)

// functions can return addresses of local variables
// functions can assign functions to variables or use as
// arguments and return values in functions

// type signature is like function signature, with no parameter names
// var f func(string, string, int) (int, error)

// methods are functions that execute in context of a type

func main() {
	greeting := "Hello"
	name := "Boi"

	// values as parameters
	sayGreeting(greeting, name)
	// pointers as parameters
	pointerGreeting(&greeting, &name)

	// with variadic parameter
	sum("Sum: ", 1, 2, 3, 4, 5)

	// single return value
	something := returnSomething()
	fmt.Println(something)

	// multiple returned values
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// anonymous function
	// immediately invoked
	func() {
		fmt.Println("Anonymous")
	}()

	// assigned to a variable
	// declare, then call
	var f func() = func() {
		fmt.Println("F function")
	}
	f()

	// call a method
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

// two parameters with the same type
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

// function with pointers as parameters
func pointerGreeting(greeting, name *string) {
	*name = "Ted"
	fmt.Println(*greeting, *name)
}

// function with variadic parameter
// the variadic parameter has to be the last one,
// it converts the values passed into a slice
func sum(msg string, values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println(msg, result)
}

// function that returns something (a string)
func returnSomething() string {
	return "lmao"
}

// function with named return value
func namedReturn() (result int) {
	result = 10
	return
}

// function with multiple returned values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// method declaration
// g greeter -> value receiver, g *greeter -> pointer receiver
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
