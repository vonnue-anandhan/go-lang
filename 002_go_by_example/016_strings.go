package main

import (
	"fmt"
	"unicode/utf8"
)

// In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point

func main() {
	const s = "สวัสดี"

	// Go string literals are UTF-8 encoded text

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		// Generates the hex values of all the bytes
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
