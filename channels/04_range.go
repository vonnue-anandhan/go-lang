package channels

import "fmt"

// Range blocks until the channel is closed

func TestRange() {
	c := make(chan int)

	go tic(c)

	tac(c)
}

func tic(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

func tac(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
