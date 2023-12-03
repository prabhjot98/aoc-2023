package util

import (
	"strconv"
	"strings"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

type Game struct {
	number int
	rounds []Round
}

type Round struct {
	Red_cubes   int
	Green_cubes int
	Blue_cubes  int
}

const (
	Red   = "red"
	Green = "green"
	Blue  = "blue"
)

func parseGameNumber(line string) (int, string) {
	tokens := strings.Split(line, ":")
	gameToken := strings.Split(tokens[0], " ")
	gameNumber, _ := strconv.Atoi(gameToken[1]) // look man, it's never gonna error out okay
	otherTokens := tokens[1]
	return gameNumber, otherTokens
}

func parseRounds(strs string) []Round {
	roundStrs := strings.Split(strs, ";")

	var rounds []Round
	for _, roundStr := range roundStrs {
		rounds = append(rounds, ParseRound(roundStr))
	}
	return rounds
}

func ParseRound(str string) Round {
	tokens := strings.Split(str, ",") // [ 1 green , " 2 blue"]
	red_cubes := 0
	green_cubes := 0
	blue_cubes := 0

	for _, token := range tokens {
		count, color := ParseToken(token)
		switch color {
		case Red:
			red_cubes = count
		case Green:
			green_cubes = count
		case Blue:
			blue_cubes = count
		}
	}
	return Round{red_cubes, green_cubes, blue_cubes}
}

func ParseToken(str string) (int, string) {
	tokens := strings.Split(str, " ")[1:]
	count, _ := strconv.Atoi(tokens[0]) // this wont error either ;)
	color := tokens[1]
	return count, color
}

func LineToGame(line string) Game {
	gameNumber, roundTokens := parseGameNumber(line)
	rounds := parseRounds(roundTokens)
	return Game{gameNumber, rounds}
}

func (g Game) GetHighestColors() (int, int, int) {
	red, green, blue := 0, 0, 0
	for _, round := range g.rounds {

		if round.Red_cubes > red {
			red = round.Red_cubes
		}
		if round.Green_cubes > green {
			green = round.Green_cubes
		}
		if round.Blue_cubes > blue {
			blue = round.Blue_cubes
		}
	}

	return red, green, blue
}

func IsValidGame(g Game, redCubes int, greenCubes int, blueCubes int) (bool, int) {
	red, green, blue := g.GetHighestColors()
	if red > redCubes || green > greenCubes || blue > blueCubes {
		return false, 0
	}
	return true, g.number
}
