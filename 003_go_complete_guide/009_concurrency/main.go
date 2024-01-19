package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	doneChan <- true
	close(doneChan)
}

func main() {
	done := make(chan bool)

	// dispatching go-routine
	go greet("Nice to meet you!", done)

	// dones[1] = make(chan bool)
	go greet("How are you?", done)

	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)

	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	for range done {
		// fmt.Println(doneChan)
	}

}
