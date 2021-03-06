package main

import (
	"fmt"
)

func main() {
	fmt.Println("I'm learning Go")
	printMessage("This is a good programming language")
	sentence := printResult("I'm learning Go")
	fmt.Println(sentence)
	fmt.Println(printResult("I'm learning Go"))
}

func printMessage(message string) {
	fmt.Println(message)
}

func printResult(message string) string {
	newMessage := message + " for daily learn"
	return newMessage
}
