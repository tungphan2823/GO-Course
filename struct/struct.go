package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time
}

func (User *user) outputUserDetails() {
	fmt.Printf("First Name: %s\n", User.firstName)
	fmt.Printf("Last Name: %s\n", User.lastName)
	fmt.Printf("Birth Date: %s\n", User.birthDate)
	fmt.Printf("Created Date: %s\n", User.createdDate.Format("2006-01-02"))
	fmt.Println("---------------------------------------------")
	fmt.Println()
}
func (User *user) clearOut() {
	User.firstName = ""
}

func newUser(userFirstName, userLastName, userBirthDate string) (*user, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		fmt.Println("Invalid input. Please try again.")
		return nil, errors.New("Invalid input. Please try again.")
	}
	return &user{
		firstName:   userFirstName,
		lastName:    userLastName,
		birthDate:   userBirthDate,
		createdDate: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Print(err)
		return
	}
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
