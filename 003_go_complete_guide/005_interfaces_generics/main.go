package main

import (
	"bufio"
	"example.com/note/note"
	"example.com/note/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface { // only 1 method ? then method + er for interface name
	// Save(int, string) error
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// Special "any value allowed" type
func printSomething(value interface{}) { // Type switches
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)

	// case float64:
	// 	fmt.Println("Float64: ", value)

	// case string:
	// 	fmt.Println("String: ", value)

	// }

	// Extracting type info from values

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
	}

	float64Val, ok := value.(float64)

	if ok {
		fmt.Println("Float64: ", float64Val)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("Float64: ", stringVal)
	}

}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	// todo

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	// note

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = saveData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)

	printSomething(todo)

	fmt.Println(superAdd(1, 2))
}

func outputData(data outputtable) error {
	data.Display()

	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for windows

	return text
}

// ************ GENERICS ************

func superAdd[T int | float64 | string](a, b T) T {
	// aInt, isAInt := a.(int)
	// bInt, isBInt := b.(int)

	// if isAInt && isBInt {
	// 	return aInt + bInt

	// }

	// return nil

	return a + b

}
