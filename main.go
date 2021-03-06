package main

import (
	"fmt"
)

func main() {
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["Go"] = 9
	myMap["PHP"] = 8
	myMap["NodeJS"] = 7

	fmt.Println(myMap)

	newMap := map[string]string{
		"satu": "one",
		"dua": "two",
		"tiga": "three",
	}

	fmt.Println(newMap)
}
