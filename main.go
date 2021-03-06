package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string{
		"satu": "one",
		"dua": "two",
		"tiga": "three",
		"empat": "four",
	}

	for key, value := range myMap {
		fmt.Println("Key :", key, " Value :", value)
	}

	fmt.Println("===")

	delete(myMap, "dua")
	// will delete myMap with key "lima"

	for key, value := range myMap {
		fmt.Println("Key :", key, " Value :", value)
	}

	tigaValue := myMap["tiga"]
	fmt.Println(tigaValue)

	limaValue, isAvailable := myMap["lima"]
	fmt.Println(isAvailable)
	fmt.Println(limaValue)
	// isAvailable being a checker whether myMap has "lima" key
}
