package main

import "fmt"

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("Here is the error : %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran - ", e, "\n")
}

func main() {
	// pass

	c1 := customError{info: "need more coffee"}

	foo(c1)
}
