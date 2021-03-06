package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	resultOne, errOne := calculate(10, 2, "+")
	fmt.Println(resultOne)
	if (errOne != nil) {
		fmt.Println(errOne)
	}
	resultTwo, errTwo := calculate(10, 2, "-")
	fmt.Println(resultTwo)
	if (errTwo != nil) {
		fmt.Println(errTwo)
	}
	resultThree, errThree := calculate(10, 2, "*")
	fmt.Println(resultThree)
	if (errThree != nil) {
		fmt.Println(errThree)
	}
	resultFour, errFour := calculate(10, 2, "/")
	fmt.Println(resultFour)
	if (errFour != nil) {
		fmt.Println(errFour)
	}
	resultFive, errFive := calculate(10, 2, "?")
	fmt.Println(resultFive)
	if (errFive != nil) {
		fmt.Println(errFive)
	}
}

func sum(scores []int) (total int) {
	for _, score := range scores {
		total = total + score
	}

	return
}

func calculate(number, numberTwo int, operator string) (result int, err error) {
	switch operator {
		case "+":
			result = number + numberTwo
		case "-":
			result = number - numberTwo
		case "*":
			result = number * numberTwo
		case "/":
			result = number / numberTwo
		default:
			err = errors.New("Error! Invalid operator")
	}

	return
}
