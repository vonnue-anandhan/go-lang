package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main18() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping(ch1, "hello world")
	go pong(ch1, ch2)

	fmt.Println(<-ch2)
}
