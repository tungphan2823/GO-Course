package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser user
	appUser = user{
		firstName:   userFirstName,
		lastName:    userLastName,
		birthDate:   userBirthDate,
		createdDate: time.Now(),
	}
	outputUserDetails(appUser)
}

func outputUserDetails(User user) {
	fmt.Printf("First Name: %s\n", User.firstName)
	fmt.Printf("Last Name: %s\n", User.lastName)
	fmt.Printf("Birth Date: %s\n", User.birthDate)
	fmt.Printf("Created Date: %s\n", User.createdDate.Format("2006-01-02"))
	fmt.Println("---------------------------------------------")
	fmt.Println()
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
