package main

import "fmt"

func main() {
	sentence := "Golang the best language"

	for index, letter := range sentence {
		if index % 2 == 0 {
			fmt.Println("Index :", index, " Letter :", string(letter))
		}
	}

	fmt.Println("---")

	for index, letter := range sentence {
		switch string(letter) {
			case "a", "i", "u", "e", "o":
				fmt.Println("Index :", index, " Letter :", string(letter))
		}
	}
}
