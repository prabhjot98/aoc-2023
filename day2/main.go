package main

import (
	"day2/util"
	"fmt"
	"log"
	"os"
	"strings"
)

func fileToString(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Could not find file name with name: %s", fileName)
	}
	return string(file)
}

func partOne() {
	gameStrs := strings.Split(fileToString("input.txt"), "\n")
	gameStrs = gameStrs[:len(gameStrs)-1]

	var games []util.Game
	for _, gameStr := range gameStrs {
		games = append(games, util.LineToGame(gameStr))
	}

	sum := 0
	for _, game := range games {
		_, id := util.IsValidGame(game, 12, 13, 14)
		sum += id
	}
	fmt.Println(sum)
}

func partTwo() {
	gameStrs := strings.Split(fileToString("input.txt"), "\n")
	gameStrs = gameStrs[:len(gameStrs)-1]

	var games []util.Game
	for _, gameStr := range gameStrs {
		games = append(games, util.LineToGame(gameStr))
	}

	totalPower := 0
	for _, game := range games {
		red, green, blue := game.GetHighestColors()
		power := red * green * blue
		totalPower += power
	}

	fmt.Println(totalPower)
}

func main() {
	partTwo()
}
