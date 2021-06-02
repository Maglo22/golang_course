package main

import (
	"fmt"
)

// maps are a key - value data structure
// every value needs to be from the same type in one map
// the key type has to be able to be tested for equality
// maps are referenced when copied

// the return order of a map is not guaranteed

func main() {
	// map[type_of_key]type_of_value
	// declaration with make -> make(map[string]int)
	statesPopulation := map[string]int{
		"California": 321654,
		"Texas":      123132,
		"Florida":    985131,
		"New York":   7945120,
	}

	fmt.Println(statesPopulation)

	// pulling a value
	// a mispell in the key will return 0, not an error
	fmt.Println(statesPopulation["Texas"])

	// to check if the key exists in the map
	pop, ok := statesPopulation["Lmao"]
	// or to ignore value -> _, ok := statesPopulation["Lmao"]
	fmt.Println(pop, ok) // 0 false

	// adding a value
	statesPopulation["Georgia"] = 101234

	// deleting a value
	delete(statesPopulation, "Georgia")
}
