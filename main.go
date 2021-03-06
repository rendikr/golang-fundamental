package main

import "fmt"

func main() {
	score := 80
	var grade string

	if score > 90 {
		grade = "A"
	} else if (score > 70) {
		grade = "B"
	} else if (score > 50) {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Println("Your Grade:")
	fmt.Println(grade)
}
