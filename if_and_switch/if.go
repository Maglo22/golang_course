package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("True")
	}

	statesPopulation := map[string]int{
		"California": 321654,
		"Texas":      123132,
		"Florida":    985131,
		"New York":   7945120,
	}

	if pop, ok := statesPopulation["Florida"]; ok {
		fmt.Println(pop)
	}

	number := 50

	if number > 1 && number < 51 {
		fmt.Println("In range")
	} else if number > 1 || number < 100 {
		fmt.Println("Almost in range")
	} else {
		fmt.Println("Not in range")
	}

}
