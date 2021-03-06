package main

import (
	"fmt"
	"golang-fundamental/calculation"
)

func main() {
	fmt.Println("Hello World!")

	// sentence := TestString()
	// fmt.Println(sentence)

	result := calculation.Add(2, 3)
	fmt.Println(result)
}
