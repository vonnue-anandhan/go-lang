package basic

import "fmt"

func buffered() {
	ch := make(chan int, 1)

	// go func() {
		ch <- 42
	// }()

	fmt.Println(<-ch)
}
