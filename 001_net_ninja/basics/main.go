package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(names []string, f func(string)) {

	for _, name := range names {
		f(name)
	}

}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, name[:1]) // instead of name[0] in JS
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

var score = 99.5

func updateName(x string) string {
	x = "Wedge"

	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func updateAge(x *int) {
	*x = 3	
}

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

	// while loop
	i := 0
	for i < 5 {
		fmt.Printf("Value of x is : %v \n", i)
		i++
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Value of i is : %v \n", i)
	}

	loopNames := []string{"Hitman", "Ipman", "Spiderman", "Superman"}
	for i := 0; i < len(loopNames); i++ {
		fmt.Println("Value is: ", loopNames[i])
	}

	for index, value := range loopNames {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	// *************************** Booleans & conditionals ***************************

	personAge := 22

	fmt.Println(personAge >= 20)

	if personAge > 30 {
		fmt.Println("Age is greater than 30")
	} else if personAge < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Nothing")
	}

	for index, value := range newNames {
		if index == 1 {
			fmt.Println("Continuing at pos:", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos:", index)
			break

		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}

	// *************************** Using functions ***************************

	sayGreeting("Mario")
	sayBye("Mario")

	cycleNames(loopNames, sayGreeting)

	areaOne := circleArea(10)
	fmt.Printf("Area: %v \n", areaOne)

	fn1, sn1 := getInitials("john doe")
	fmt.Println(fn1, sn1)

	// *************************** Package scope ***************************

	sayHello("John Doe")

	for _, point := range points {
		fmt.Println(point)
	}

	showScore()

	// *************************** Maps (like objects in JS) ***************************

	menu := map[string]float64{"soup": 4.99, "pie": 7.99, "salad": 6.99, "toffee pudding": 3.55}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, item := range menu {
		fmt.Println(k, "-", item)
	}

	// ints as key type
	phoneBook := map[int]string{255244: "Mario", 255308: "Yoshi"}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[255244])

	phoneBook[255244] = "Claw"

	fmt.Println(phoneBook)

	// *************************** Pass by value ***************************

	// Group A types => strings, ints, floats, bools, arrays, structs => makes a copy
	newName := "Tifa"

	newName = updateName(newName)

	fmt.Println(newName)

	// Group B types => slices, maps, functions => still makes a copy - but its copying the pointer not the actual data

	foodMenu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(foodMenu)

	fmt.Println(foodMenu)

	// *************************** Pointers ***************************

	m := &newName // Storing memory address of another variable
	fmt.Println("Memory address: ", m)
	fmt.Println("Value at memory address: ", *m)

	newAge := 10

	fmt.Println("Old newAge: ", newAge)

	n := &newAge
	updateAge(n)

	fmt.Println("new newAge: ", newAge)

}
