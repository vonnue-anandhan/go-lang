package testing

import "fmt"

func MySum(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func RunMySum() {
	fmt.Println("1 + 1 = ", MySum(1, 1))
	fmt.Println("2 + 2 = ", MySum(2, 2))
	fmt.Println("3 + 3 = ", MySum(3, 3))
}
