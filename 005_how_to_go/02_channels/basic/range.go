package basic

// import "fmt"

// func main() {
// 	ch := make(chan int)

// 	go send(ch)

// 	// receive

// 	for v := range ch { // because of range close ch
// 		fmt.Println(v)
// 	}
// }

// func send(ch chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		ch <- i
// 	}

// 	close(ch)
// }
