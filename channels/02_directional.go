package channels

import "fmt"

func sendChan() {
	c := make(chan<- int, 1)
	c <- 22

	// fmt.Println(<-c) ERROR
}

func receiveChan() {
	c := make(<-chan int, 1)
	// close(c) ERROR: cannot close receive only channel
	// c <- 1000 ERROR

	fmt.Println(<-c)
}

func TestDirectional() {
	sendChan()
	receiveChan()
}
