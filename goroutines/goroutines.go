package main

import (
	"fmt"
	"runtime"
	"sync"
)

// when using anonymous functions, pass data as local variables

// best practices
// don't creatre goroutines in libraries
// know how a goroutine will end; avoid memory leaks
// check for race conditions at compile time -> go run -race {file.go}

// sync multiple goroutines toguether
var wg = sync.WaitGroup{}

// lock for sync
var m = sync.RWMutex{}
var counter = 0

func main() {
	runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		// add goroutines
		wg.Add(2)
		m.RLock() // read lock
		go sayHello()
		m.Lock() // write lock
		go increment()
	}

	// wait for goroutines
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // release read lock

	wg.Done()
}

// increment counter
func increment() {
	counter++
	m.Unlock() // release write lock

	wg.Done()
}
