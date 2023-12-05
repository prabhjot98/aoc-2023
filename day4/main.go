package main

import (
	"bufio"
	"day4/util"
	"fmt"
	"os"
)

func Part_two() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scratch_cards []util.ScratchCard
	for scanner.Scan() {
		sc := util.ParseLineToScratchCard(scanner.Text())
		scratch_cards = append(scratch_cards, sc)
	}

	sc_map := make(map[util.CardNumber]int)

	for i := len(scratch_cards) - 1; i >= 0; i-- {
		sc := scratch_cards[i]
		matches := sc.GetMatches()
		if matches == 0 {
			sc_map[sc.Card_number] = 1
		} else {
			match_sum := 1
			for j := 1; j <= matches; j++ {
				match_sum += sc_map[sc.Card_number+util.CardNumber(j)]
			}
			sc_map[sc.Card_number] = match_sum
		}
	}

	sum := 0

	for _, v := range sc_map {
		sum += v
	}

	fmt.Println(sum)
}

func main() {
	// part one was in rust
	Part_two()
}
