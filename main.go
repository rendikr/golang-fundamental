package main

import (
	"fmt"
	"golang-fundamental/calculation"
)

func main() {
	fmt.Println("Hello World!")

	// sentence := TestString()
	// fmt.Println(sentence)

	addResult := calculation.Add(2, 3)
	fmt.Println(addResult)

	multiplyResult := calculation.Multiply(5, 3)
	fmt.Println(multiplyResult)
}
