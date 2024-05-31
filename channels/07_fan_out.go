package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func TestFanOut() {
	const LIMIT = 10
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1, LIMIT)

	go fanOutIn(c1, c2)

	for value := range c2 {
		fmt.Println(value)
	}
}

func populate(c chan<- int, limit int) {
	for i := 0; i < limit; i++ {
		c <- i
	}

	close(c)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	wg := &sync.WaitGroup{}

	for value := range c1 {
		wg.Add(1)
		go func(value int) {
			c2 <- timeConsumingWork(value)
			wg.Done()
		}(value)
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))

	return n + rand.Intn(1000)
}
