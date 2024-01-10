package main

import "fmt"

func main() {
	mybill := newBill("Mario's bill")

	mybill.format()

	mybill.updateTip(10)
	mybill.addItem("onion soup", 4.50)

	fmt.Println(mybill.format())

}
