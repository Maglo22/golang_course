package main

import (
	"fmt"
	"log"
)

// panics executes after defer calls

// when recover is used, the current function will not
// continue, but higher functions in the call stack will
// only useful in deferred functions

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	// recover from the panic
	defer func() {
		if err := recover(); err != nil {
			log.Panicln("Error:", err)
			// panic(err) -> if desired, can be repanicked
		}
	}()
	panic("something happened")
}
