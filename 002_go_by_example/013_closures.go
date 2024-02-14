package main

import "fmt"

func incrementor() func() int {
	i := 0

	// Returning an anonymous function
	return func() int {
		i++
		return i
	}
}

func main13() {
	inc := incrementor()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	newInts := incrementor()
	fmt.Println(newInts())
}
