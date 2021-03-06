package main

import (
	"fmt"
)

type User struct {
	ID int
	FirstName string
	LastName string
	Username string
	IsActive bool
}

func main() {
	firstUser := User{}
	firstUser.ID = 1
	firstUser.FirstName = "Rendi"
	firstUser.LastName = "Kurnia"
	firstUser.Username = "rendikr"
	firstUser.IsActive = true

	secondUser := User{
		ID: 2,
		FirstName: "Asep",
		LastName: "Jaelani",
		Username: "ajae",
		IsActive: true,
	}

	fmt.Println(displayUser(firstUser))
	fmt.Println(displayUser(secondUser))
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s Username : %s", user.FirstName, user.LastName, user.Username)
}
