package main

import (
	"fmt"
	"structs/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {

	var greeting str = "hello world!!"
	greeting.log()

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")
	admin.User.OutputUserDetails() // *** IMP ***
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
