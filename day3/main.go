package main

import (
	"day3/util"
	"fmt"
	"log"
	"os"
	"slices"
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
	lines := strings.Split(fileToString("input.txt"), "\n")
	lines = lines[:len(lines)-1]

	parts, symbols := util.MakeThings(lines)

	var validParts []util.Part
	var sum int

	for _, symbol := range symbols {
		for _, part := range parts {
			if symbol.IsAdjacentTo(part) && !slices.Contains(validParts, part) {
				validParts = append(validParts, part)
				sum += int(part.PartNumber)
			}
		}
	}

	fmt.Println(sum)
}

func partTwo() {
	lines := strings.Split(fileToString("input.txt"), "\n")
	lines = lines[:len(lines)-1]

	parts, symbols := util.MakeThings(lines)

	var validParts []util.Part
	var sum int

	gears := make(map[util.Symbol][]util.Part)

	for _, symbol := range symbols {
		for _, part := range parts {
			if symbol.IsAdjacentTo(part) && !slices.Contains(validParts, part) {
				validParts = append(validParts, part)
				gears[symbol] = append(gears[symbol], part)
			}
		}
	}

	for _, parts := range gears {
		if len(parts) == 2 {
			sum += int(parts[0].PartNumber) * int(parts[1].PartNumber)
		}

	}

	fmt.Println(sum)
}

func main() {
	partTwo()
}
