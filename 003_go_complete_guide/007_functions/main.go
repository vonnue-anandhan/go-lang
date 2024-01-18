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

	fact := factorial(5)
	fmt.Println(fact)

	fmt.Println(sumAll(1, 2, 3))
	fmt.Println(sumAll(numbers...))
}

func getMultipliedValues(numbers *[]int, multiplier multiplier) []int {
	dNumbers := []int{}

	for _, number := range *numbers {
		dNumbers = append(dNumbers, multiplier(number))
	}

	return dNumbers
}

// Closure
func getMultiplier(factor int) multiplier {
	return func(number int) int {
		return number * factor
	}
}

// Recursive functions
func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func sumAll(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
