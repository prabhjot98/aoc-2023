package util

import (
	"strconv"
	"strings"
)

// destination-to-source map :
type MapLine struct {
	Destination int
	Source      int
	Range       int
}

type DestinationSourceMap struct {
	MapLines []MapLine
}

func (d DestinationSourceMap) GetDestination(source int) int {
	for _, mapLine := range d.MapLines {
		if source >= mapLine.Source && source < mapLine.Source+mapLine.Range {
			return mapLine.Destination + (source - mapLine.Source)
		}
	}
	return source
}

func ParseLineToMap(line string) MapLine {
	strs := strings.Fields(line)

	dest, _ := strconv.Atoi(strs[0])
	src, _ := strconv.Atoi(strs[1])
	rng, _ := strconv.Atoi(strs[2])

	return MapLine{dest, src, rng}
}

func ParseLineToSeeds(line string) []int {
	_, right, _ := strings.Cut(line, ": ")
	strs := strings.Fields(right)

	var seeds []int
	for _, str := range strs {
		seed, _ := strconv.Atoi(str)
		seeds = append(seeds, seed)
	}

	return seeds
}

func ParseRangeLineToSeeds(line string) []int {
	_, right, _ := strings.Cut(line, ": ")
	strs := strings.Fields(right)

	var seeds []int
	for i := 0; i < len(strs); i += 2 {
		seed, _ := strconv.Atoi(strs[i])
		rng, _ := strconv.Atoi(strs[i+1])

		for j := 0; j < rng; j++ {
			seeds = append(seeds, seed+j)
		}
	}

	return seeds
}
