package basic

import "fmt"

func directional() {
	// ch1 := make(chan<- int, 2)
	ch2 := make(<-chan int, 2)


	// fmt.Println("%T\n", <-ch1) ERROR!!
	fmt.Println("%T\n", <-ch2)

}
