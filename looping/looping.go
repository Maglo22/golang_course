package main

import (
	"fmt"
)

func main() {
	// index scoped to the loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// with two index
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// index scoped to the function
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// incrementer inside loop block
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite for loop
	for {
		if i%2 == 0 {
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println("oh no")
		i++
	}

	// nested loop with label
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j >= 3 {
				// get out of both loops, thanks to the label
				break Loop
			}

			fmt.Println(i * j)
		}
	}

	// looping a collection
	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Println(key, value)
	}

}
