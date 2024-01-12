package main

import (
	"fmt"
	"maps"
)

func main() {

	// Create an empty map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn’t exist, the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove all key/value pairs from a map, use the clear builtin
	clear(m)
	fmt.Println("map:", m)

	// If the key is present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
