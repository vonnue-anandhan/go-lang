package main

import "fmt"

type multiplier func(number int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubler := getMultiplier(2)
	tripler := getMultiplier(3)

	doubled := getMultipliedValues(&numbers, doubler)
	tripled := getMultipliedValues(&numbers, tripler)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func getMultipliedValues(numbers *[]int, multiplier multiplier) []int {
	dNumbers := []int{}

	for _, number := range *numbers {
		dNumbers = append(dNumbers, multiplier(number))
	}

	return dNumbers
}

func getMultiplier(factor int) multiplier {
	return func(number int) int {
		return number * factor
	}
}
