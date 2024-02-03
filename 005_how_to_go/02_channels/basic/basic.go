package basic

import "fmt"

func basic() {
	ch := make(chan int)

	// go func() {
		ch <- 42
	// }()

	fmt.Println(<-ch)
}
