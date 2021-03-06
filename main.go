package main

import (
	"fmt"
)

func main() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Playstation 4")
	gamingConsoles = append(gamingConsoles, "Playstation 5")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")

	fmt.Println(gamingConsoles)

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
