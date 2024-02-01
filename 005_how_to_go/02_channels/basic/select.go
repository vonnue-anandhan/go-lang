package basic

// import "fmt"

// func main() {
// 	evenChan := make(chan int)
// 	oddChan := make(chan int)
// 	quitChan := make(chan int)

// 	go send(evenChan, oddChan, quitChan)

// 	receive(evenChan, oddChan, quitChan)

// 	fmt.Println("about to exit")
// }

// func send(evenChan chan int, oddChan chan int, quitChan chan int) {
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			evenChan <- i
// 		} else {
// 			oddChan <- i
// 		}
// 	}

// 	close(evenChan)
// 	close(oddChan)

// 	quitChan <- 0
// }

// func receive(evenChan chan int, oddChan chan int, quitChan <-chan int) {
// 	for {
// 		select {
// 		case n := <-evenChan:
// 			fmt.Println("From even channel", n)

// 		case n := <-oddChan:
// 			fmt.Println("From odd channel", n)

// 		case q := <-quitChan:
// 			fmt.Println("From even channel", q)
// 			return
// 		}
// 	}

// }
