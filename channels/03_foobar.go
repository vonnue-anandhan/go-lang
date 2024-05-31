package channels

import "fmt"

func foo(ch chan<- int) {
	ch <- 42
	close(ch)
}

func bar(ch <-chan int) {
	fmt.Println(<-ch)
}

func TestFoobar() {
	ch := make(chan int)

	go foo(ch)
	bar(ch)
}
