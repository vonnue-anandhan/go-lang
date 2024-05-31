package channels

import "fmt"

// Code will BLOCK unless send & receive happens at the same time / use a buffered channel
// Receive only without send from a buffered channel is also BLOCKING
// Receive only channels cannot be closed

func blocking() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

func nonBlocking() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bufferedChan() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

func closedChan() {
	c := make(chan int)
	close(c)

	value, ok := <-c
	fmt.Println("Value:", value, ", Is channel open: ", ok)
}

func onlySend() {
	c := make(chan int)

	c <- 100
}

func onlySendBuffered() {
	c := make(chan int, 1)

	c <- 100
}

func bufferedWithoutSend() {
	c := make(chan int, 1)

	fmt.Println(<-c)
}

func TestBasics() {
	// blocking() ERROR
	nonBlocking()
	bufferedChan()
	closedChan()
	// onlySend() ERROR
	onlySendBuffered()
	// bufferedWithoutSend() ERROR
}
