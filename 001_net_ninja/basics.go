package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// String
	var name string = "John Doe"
	var nickName string

	// Style
	nickName = "sdsd"

	fmt.Println(name)
	fmt.Println(nickName)

	lastName := "Hitman" // This style is only allowed inside functions

	fmt.Println(lastName)

	// Int
	var age int = 20
	var speed = 30
	price := 3000

	fmt.Println(age, speed, price)

	// Bits and memory
	// var numOne int8 = 2523 ERROR!!
	var numOne int32 = 2523
	// var positiveNumber uint = -124 ERROR!!

	var score float32 = 123.23

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")

	// Println
	fmt.Println("Hello ninjas!")
	fmt.Println("My age is :", numOne, score)

	// Printf (%_) | %v is the default format specifier for var
	fmt.Printf("My age is %v and my score is %v", age, score)

	// type format specifier
	fmt.Printf("Age is of type %T", age)

	fmt.Printf("You scrored %0.1f", 20.55)

	// save formatted strings
	var str = fmt.Sprintf("My age is %v and my score is %v", 10, 20000)
	fmt.Println(str)

	// *************************** Arrays and slices ***************************

	// Arrays

	// var ages [3]int = [3]int{10, 20, 30}

	var ages = [3]int{10, 20, 30}

	names := []string{"Hitman", "Ipman", "Spiderman", "Superman"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)

	var scores = []int{100, 20, 1}
	scores = append(scores, 90)

	fmt.Println(scores)

	// slice ranges
	slicedScores := scores[0:1]
	fmt.Println("sliced scores", slicedScores)

	greeting := "Hello world"

	// *************************** Standard library ***************************

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	newAges := []int{100, 20, 200, 3, 5, 7, 0}
	sort.Ints(newAges)

	fmt.Println(newAges)

	index := sort.SearchInts(newAges, 30)
	fmt.Println(index)

	newNames := []string{"ATest 1", "CTest 2", "ZTest 3", "BTest 4", "YTest 5"}
	sort.Strings(newNames)

	fmt.Println(sort.SearchStrings(names, "Hitman"))

	// *************************** Loops ***************************
}
