package main

import (
	"fmt"
)

func main() {
	// create a slice with value is a map
	students := []map[string]string{
		{"name": "Rendi K.", "score": "A"},
		{"name": "Asep Jaelani", "score": "B"},
		{"name": "Maman Jawa", "score": "C"},
	}

	for _, student := range students {
		fmt.Println("Student Name :", student["name"], " Score :", student["score"])
	}
}
