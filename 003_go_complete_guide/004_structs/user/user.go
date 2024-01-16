package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// user User
	User // Directly
}

// Receiver functions / methods
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Constructor pattern

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		// Can return nil as a valid value for a pointer type
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User:     User{firstName: "Admin", lastName: "Admin", birthdate: "-", createdAt: time.Now()},
	}
}

// func newUser(firstName, lastName, birthDate string) (user, error) {
// 	if firstName == "" || lastName == "" || birthDate == "" {
// 		return user{}, errors.New("first name, last name and birth date are required")
// 	}

// 	return user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthdate: birthDate,
// 		createdAt: time.Now(),
// 	}, nil
// }
