package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func main() {
	// mybill := newBill("Mario's bill")

	// mybill.format()

	// mybill.updateTip(10)
	// mybill.addItem("onion soup", 4.50)

	// fmt.Println(mybill.format())

	mybill := createBill()

	fmt.Println(mybill)

}
