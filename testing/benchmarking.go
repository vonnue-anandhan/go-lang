package testing

import "fmt"

func Greet(s string) string {
	return fmt.Sprintf("Welcome my dear %v", s)
}

func RunBenchmarking() {
	fmt.Println(Greet("Janice"))
}
