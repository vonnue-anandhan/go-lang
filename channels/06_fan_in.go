package channels

import (
	"fmt"
	"sync"
)

func TestFanIn() {
	const LIMIT = 10

	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go sendFanIn(even, odd, LIMIT)

	go receiveFanIn(even, odd, fanIn)

	for number := range fanIn {
		fmt.Println(number)
	}
}

func sendFanIn(e, o chan<- int, limit int) {
	for i := 1; i < limit+1; i++ {
		isEven := i%2 == 0

		if isEven {
			e <- i
		} else {
			o <- i
		}
	}

	close(e)
	close(o)
}

func receiveFanIn(e, o <-chan int, fanIn chan<- int) {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		for number := range e {
			fanIn <- number
		}

		wg.Done()
	}()

	go func() {
		for number := range o {
			fanIn <- number
		}

		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}
