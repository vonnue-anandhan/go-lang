package main

import (
	"fmt"
	"sync"
)

// ******* Channels *******

func main() {
	fmt.Println("Channels in go lang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// pass value into channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer (*wg).Done()

		ch <- 5
		// ch <- 10

		close(ch)
		// ch <- 6
	}(myCh, wg)

	// receive value from channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(ch) receive only channel
		defer (*wg).Done()

		val, isChannelOpen := <-ch

		if isChannelOpen {
			fmt.Println(<-ch)
			fmt.Println(val)
		}
	}(myCh, wg)

	wg.Wait()
}
