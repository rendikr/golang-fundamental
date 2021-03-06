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

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Username : %s", user.FirstName, user.LastName, user.Username)
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func (group Group) display() {
	fmt.Printf("Group Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Admin : %s %s", group.Admin.FirstName, group.Admin.LastName)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")

	for index, user := range group.Users {
		fmt.Printf("User %d : %s %s", (index + 1), user.FirstName, user.LastName)
		fmt.Println("")
	}
}

func main() {
	firstUser := User{}
	firstUser.ID = 1
	firstUser.FirstName = "Rendi"
	firstUser.LastName = "Kurnia"
	firstUser.Username = "rendikr"
	firstUser.IsActive = true

	displayUser := firstUser.display()
	fmt.Println(displayUser)

	secondUser := User{
		ID: 2,
		FirstName: "Asep",
		LastName: "Jaelani",
		Username: "ajae",
		IsActive: true,
	}

	thirdUser := User{
		ID: 3,
		FirstName: "Maman",
		LastName: "Jawa",
		Username: "Mamawa",
		IsActive: true,
	}

	users := []User{firstUser, secondUser, thirdUser}

	group := Group{"Gamer", secondUser, users, true}

	group.display()
}
