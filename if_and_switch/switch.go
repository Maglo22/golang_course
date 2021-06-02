package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("three, four or five")
	default:
		fmt.Println("another")
	}

	// tagless switch
	i := 10

	switch {
	case i <= 10:
		fmt.Println("less or equial to ten")
		fallthrough // force execute the next case
	case i <= 20:
		fmt.Println("less or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// typed switch
	var j interface{} = 1

	// compare the type of the variable
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
		break // force leave switch
		// fmt.Println("This wont print")
	case float64:
		fmt.Println("j is a float64")
	case string:
		fmt.Println("j is a string")
	default:
		fmt.Print("j is another type")
	}

}
