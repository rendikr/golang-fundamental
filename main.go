package main

import "fmt"

func main() {
	var languages [5]string

	languages[0] = "Go"
	languages[1] = "PHP"
	languages[2] = "NodeJS"
	languages[3] = "ReactJS"
	languages[4] = "VueJS"

	length := len(languages)
	fmt.Println(languages)
	fmt.Println(length)

	ages := [3]int{4, 2, 31}

	ageLength := len(ages)
	fmt.Println(ages)
	fmt.Println(ageLength)

	topics := [3]string{
		"General",
		"Finance",
		"Customer Service",
	}

	topicsLength := len(topics)
	fmt.Println(topics)
	fmt.Println(topicsLength)

	engineers := [...]string{
		"Frontend Engineer",
		"Backend Engineer",
		"DevOps Engineer",
	}

	engineersLength := len(engineers)
	fmt.Println(engineers)
	fmt.Println(engineersLength)

	for index, engineer := range engineers {
		fmt.Println("Index :", index, " Engineer :", engineer)
	}
}
