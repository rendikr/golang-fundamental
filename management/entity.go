package management

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
	Username string
	IsActive bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Username : %s", user.FirstName, user.LastName, user.Username)
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func (group Group) Display() {
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
