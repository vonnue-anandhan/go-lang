package main

import (
	"fmt"
	"math"
)


// Go supports constants of character, string, boolean, and numeric values.

const s string = "constant"

func main4() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
