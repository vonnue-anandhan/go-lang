package main

import "fmt"

// *************************** Structs & custom types ***************************

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 2.99},
		tip:   0,
	}

	return b
}

// *************************** Receiver functions (like methods) ***************************

// Format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for key, price := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", price)
		total += price
	}

	// Add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// Update tip
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip - dereferencing 
	b.tip = tip
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
