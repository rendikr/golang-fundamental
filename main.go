package main

import (
	"fmt"
	"golang-fundamental/management"
)

func main() {
	firstUser := management.User{}
	firstUser.ID = 1
	firstUser.FirstName = "Rendi"
	firstUser.LastName = "Kurnia"
	firstUser.Username = "rendikr"
	firstUser.IsActive = true

	displayUser := firstUser.Display()
	fmt.Println(displayUser)

	secondUser := management.User{
		ID: 2,
		FirstName: "Asep",
		LastName: "Jaelani",
		Username: "ajae",
		IsActive: true,
	}

	thirdUser := management.User{
		ID: 3,
		FirstName: "Maman",
		LastName: "Jawa",
		Username: "Mamawa",
		IsActive: true,
	}

	users := []management.User{firstUser, secondUser, thirdUser}

	group := management.Group{"Gamer", secondUser, users, true}

	group.Display()
}
