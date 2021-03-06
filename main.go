package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("I'm learning Go! :", i)
	}

	n := 11
	for n <= 20 {
		fmt.Println("I'm learning Go! :", n)
		n++
	}

	title := "Golang Programming Language"
	for index, letter := range title {
		fmt.Println("Index :", index, " Letter :", letter, " Letter String: ", string(letter))
	}

	newTitle := "Golang Best For Learning"
	for _, letter := range newTitle {
		fmt.Println("Letter :", letter, " Letter String: ", string(letter))
	}
}
