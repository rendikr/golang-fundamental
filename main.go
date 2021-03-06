package main

import (
	"fmt"
)

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(gameName string) {
	gamer.Games = append(gamer.Games, gameName)
}

func main() {
	gamer := Gamer{Name: "Asep"}

	gamer.AddGame("BioShock 3")
	gamer.AddGame("Hunt: Showdown")
	gamer.AddGame("Battlefield 1")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
