package lists

import "fmt"

// func main() {
// 	numbers := [4]int{0, 1, 2, 3}

// 	featuredNumbers := numbers[1:] // 1, 2, 3
// 	featuredNumbers[0] = 19

// 	highlightedNumbers := featuredNumbers[:1] // 19

// 	highlightedNumbers = highlightedNumbers[:3]
// 	fmt.Println(highlightedNumbers)
// 	fmt.Println(len(highlightedNumbers), cap(highlightedNumbers))
// }

type Product struct {
	id    string
	title string
	price float64
}

func ss() {
	// 1
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3
	mainHobbies := hobbies[:2] // sports, cooking

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3] // cooking,reading
	fmt.Println(mainHobbies, hobbies)

	// 5
	courseGoals := []string{"Learn go", "Learn everything"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Learn details"
	courseGoals = append(courseGoals, "Learn all the basics")

	fmt.Println(courseGoals)

	// 7
	products := []Product{{id: "1", title: "Chair", price: 10.99}}
	fmt.Println(products)

	userNames := make([]int, 2, 3)

	userNames = append(userNames, 10, 20, 304, 50, 23230)

	fmt.Println(userNames)
	fmt.Println(cap(userNames))
}
