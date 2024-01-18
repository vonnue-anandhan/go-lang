package main

import "fmt"

// type aliases
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"google":              "https://google.com",
		"amazon web services": "https://aws.com",
	}

	fmt.Println(websites["google"])

	websites["linkedin"] = "https://linkedin.com"
	delete(websites, "linkedin")

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 5.0
	courseRatings["js"] = 3.5
	courseRatings["node"] = 4.5
	courseRatings["react"] = 2.5

	courseRatings.output()

}
