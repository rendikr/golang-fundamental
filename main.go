package main

import (
	"fmt"
)

func main() {
	sum := add(5, 9)
	fmt.Println(sum)
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}
