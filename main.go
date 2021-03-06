package main

import (
	"fmt"
)

type Student struct {
	ID int
	Name string
	GPA float32
}

func main() {
	student := Student{1, "Rendi K.", 3.32}

	fmt.Println(student.Name)

	graduate(&student)

	fmt.Println(student.Name)
}

func graduate(student *Student) {
	student.Name = student.Name + " S.T"
}
