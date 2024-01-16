package main

import "fmt"

func main() {
	age := 32 // Regular variable

	agePointer := &age

	// fmt.Println("Age:", age)

	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Age:", *agePointer)

	modifyAdultYears(agePointer)

	fmt.Println(age)
}

func modifyAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
