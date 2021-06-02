package main

import (
	"fmt"
	"sync"
)

// channels are used to sync goroutines
// send message into channel: ch <- val
// receive message from channel: val := <-ch
// a channel can have multiple senders and receivers

// by default, channels are bidirectional,
// can be specified when passing the channel:
// send-only: chan<- type
// receive-only: <-chan type

// buffered channels are useful when the sender and receiver operate
// at different frequencies
// buffered channels block sender side till receiver is available;
// block receiver side till message is available

// the select statement allows goroutine to monitor several channels at once;
// blocks if all channels block;
// if multiple channels receive value simultaneously, behaviour is undefined

var wg = sync.WaitGroup{}

func main() {
	// create int channel
	ch := make(chan int)

	// signal only channel; occupies no memory
	// var doneCh = make(chan struct{})

	for j := 0; j < 5; j++ {
		wg.Add(2)

		// receiving data
		go func(ch <-chan int) {
			for {
				if i, ok := <-ch; ok {
					fmt.Println(i)
				} else {
					break
				}
			}

			wg.Done()
		}(ch)

		// channel with buffer
		/*
			// store up to 50 elements
			go func(ch <-chan int, 50) {
				i := <-ch
				fmt.Println(i)
				wg.Done()
			}(ch)
		*/

		// sending data
		go func(ch chan<- int) {
			ch <- 42
			ch <- 27
			close(ch) // close channel to avoid deadlock

			wg.Done()
		}(ch)
	}

	wg.Wait()
}
