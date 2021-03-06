package main

import (
	"fmt"
)

type Student struct {
	ID int
	Name string
	GPA float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
}

func main() {
	student := Student{1, "Rendi K.", 3.32}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}
