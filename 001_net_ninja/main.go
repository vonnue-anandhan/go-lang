package main

import "fmt"

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

	// Printf (%_) | %v is the default format for var
	fmt.Printf("My age is %v and my score is %v", age, score)
}
