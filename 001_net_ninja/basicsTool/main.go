package main

import "fmt"

func main() {
	mybill := newBill("Mario's bill")

	mybill.format()

	fmt.Println(mybill.format())
}
