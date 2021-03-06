package main

import (
	"fmt"
)

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int
	var total int
	// goodScores has a score >= 90

	for _, score := range scores {
		if (score >= 90) {
			goodScores = append(goodScores, score)
		}

		total = total + score
	}

	average := float64(total) / float64(len(scores))
	fmt.Println("Average :", average)
	fmt.Println("Good Scores :", goodScores)
}
