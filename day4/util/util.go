package util

import (
	"slices"
	"strconv"
	"strings"
)

type CardNumber int
type Number int

type ScratchCard struct {
	Card_number     CardNumber
	Winning_numbers []Number
	Scratch_numbers []Number
}

func (s ScratchCard) GetMatches() int {
	match := 0
	for _, winning_number := range s.Winning_numbers {
		if slices.Contains(s.Scratch_numbers, winning_number) {
			match++
		}
	}
	return match
}

func ParseLineToScratchCard(line string) ScratchCard {
	left, right, _ := strings.Cut(line, ":")

	card_number, _ := strconv.Atoi(strings.Fields(left)[1])

	winning_string, scratch_string, _ := strings.Cut(right, "|")

	winning_number_strings := strings.Fields(winning_string)
	scratch_number_string := strings.Fields(scratch_string)

	var winning_numbers []Number
	for _, str := range winning_number_strings {
		num, _ := strconv.Atoi(str)
		winning_numbers = append(winning_numbers, Number(num))
	}

	var scratch_numbers []Number
	for _, str := range scratch_number_string {
		num, _ := strconv.Atoi(str)
		scratch_numbers = append(scratch_numbers, Number(num))
	}

	return ScratchCard{CardNumber(card_number), winning_numbers, scratch_numbers}
}
