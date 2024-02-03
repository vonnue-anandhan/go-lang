package main

import "fmt"

func main() {
	ch := gen()

	for value := range ch {
		fmt.Println(value)
	}

}

func gen() <-chan int {
	ch := make(chan int)

	// go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}

		close(ch)
	// }()

	return ch
}
