package util

import (
	"regexp"
	"strconv"
)

type PartNumberId int

type Coord struct {
	X, Y int
}

func (c1 Coord) Equals(c2 Coord) bool {
	if c1.X == c1.X && c2.Y == c2.Y {
		return true
	}
	return false
}

type Part struct {
	Coord      Coord
	PartNumber PartNumberId
}

func (p1 Part) Equals(p2 Part) bool {
	if p1.Coord.Equals(p2.Coord) && p1.PartNumber == p2.PartNumber {
		return true
	}
	return false
}

type Symbol struct {
	Coord Coord
}

func (s Symbol) IsAdjacentTo(p Part) bool {
	partLength := p.PartNumber.GetLength()

	if s.Coord.X >= p.Coord.X-1 && s.Coord.X <= p.Coord.X+partLength {
		if s.Coord.Y >= p.Coord.Y-1 && s.Coord.Y <= p.Coord.Y+1 {
			return true
		}
	}
	return false
}

func (p PartNumberId) GetLength() int {
	digits := 1
	for p/10 > 0 {
		digits++
		p /= 10
	}
	return digits
}

type EngineSchematic struct {
}

func flatten[T any](list [][]T) []T {
	var flat []T
	for _, item := range list {
		flat = append(flat, item...)
	}
	return flat
}

func ParsePartNumbers(line string) ([]PartNumberId, []int) {
	re := regexp.MustCompile("\\d+")
	matches := re.FindAll([]byte(line), -1)
	ranges := flatten(re.FindAllIndex([]byte(line), -1))

	var indices []int
	throwAway := false

	for _, r := range ranges {
		if !throwAway {
			indices = append(indices, r)
		}
		throwAway = !throwAway
	}

	var partNumbers []PartNumberId
	for _, match := range matches {
		partNumber, _ := strconv.Atoi(string(match))
		partNumbers = append(partNumbers, PartNumberId(partNumber))

	}
	return partNumbers, indices
}

func ParseSymbols(line string) ([]string, []int) {
	re := regexp.MustCompile("[^.\\d]")
	matches := re.FindAll([]byte(line), -1)
	ranges := flatten(re.FindAllIndex([]byte(line), -1))

	var indices []int
	throwAway := false

	for _, r := range ranges {
		if !throwAway {
			indices = append(indices, r)
		}
		throwAway = !throwAway
	}

	var partNumbers []string
	for _, match := range matches {
		partNumber := (string(match))
		partNumbers = append(partNumbers, (partNumber))

	}
	return partNumbers, indices
}

func MakeThings(lines []string) ([]Part, []Symbol) {
	var parts []Part
	var symbols []Symbol

	for i, line := range lines {
		partNumbers, partNumberIndices := ParsePartNumbers(line)
		symbls, symbolIndices := ParseSymbols(line)

		for j, partNumber := range partNumbers {
			newThing := Part{Coord{partNumberIndices[j], i}, partNumber}
			parts = append(parts, newThing)
		}

		for j := range symbls {
			newThing := Symbol{Coord{symbolIndices[j], i}}
			symbols = append(symbols, newThing)
		}
	}
	return parts, symbols
}
