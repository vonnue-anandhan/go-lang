package channels

import "fmt"

// The primary use case for select is to wait on multiple channels
// The select statement blocks until one of its cases can proceed
// If multiple cases are ready, one is chosen at random
// If none of the channels are ready and there is no default case, the select statement will block indefinitely until at least one of the channels is ready
// If a default case is present, it will execute immediately if none of the other cases can proceed, making the select statement non-blocking in that scenario

func TestSelect() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 1; i < 11; i++ {
		isEven := i%2 == 0

		if isEven {
			e <- i
		} else {
			o <- i
		}
	}

	q <- true

	close(o)
	close(e)
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case value := <-e:
			fmt.Println("Even: ", value)

		case value := <-o:
			fmt.Println("Odd: ", value)

		case _, ok := <-q:
			if !ok {
				fmt.Println("Got all even and odd numbers. Exiting.....")
				return
			} else {
				fmt.Println("Not yet done")
			}

		default:
			fmt.Println("No number")
		}
	}
}
