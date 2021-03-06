package main

import (
	"fmt"
)

func main() {
	number := 5
	fmt.Println("Old value :", number)
	fmt.Println("Memory: ", &number)

	change(&number, 100)

	fmt.Println("New value :", number)
	fmt.Println("Memory: ", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Value inside function :", *old)
	fmt.Println("Memory inside function: ", old)
}
