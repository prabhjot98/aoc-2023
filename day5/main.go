package main

import (
	"bufio"
	"day5/util"
	"fmt"
	"math"
	"os"
)

func makeSourceDestinationMap(sc *bufio.Scanner) util.DestinationSourceMap {
	var maps []util.MapLine

	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		maps = append(maps, util.ParseLineToMap(sc.Text()))
	}

	return util.DestinationSourceMap{MapLines: maps}

}

func part_one() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line which have all the seed info
	scanner.Scan()
	seedLine := scanner.Text()

	seeds := util.ParseLineToSeeds(seedLine)
	scanner.Scan() // skip the first whitespace line

	scanner.Scan()
	soilSeedMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	fertilizerSoilMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	waterFertilizerMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	lightWaterMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	tempLightMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	humidTempMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	locationHumidMap := makeSourceDestinationMap(scanner)

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := locationHumidMap.GetDestination(
			humidTempMap.GetDestination(
				tempLightMap.GetDestination(
					lightWaterMap.GetDestination(
						waterFertilizerMap.GetDestination(
							fertilizerSoilMap.GetDestination(soilSeedMap.GetDestination(seed)),
						),
					),
				),
			),
		) // oh yeah baby, this is REAL programming

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println(lowestLocation)
}

func part_two() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line which have all the seed info
	scanner.Scan()
	seedLine := scanner.Text()

	seeds := util.ParseRangeLineToSeeds(seedLine)
	scanner.Scan() // skip the first whitespace line

	scanner.Scan()
	soilSeedMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	fertilizerSoilMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	waterFertilizerMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	lightWaterMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	tempLightMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	humidTempMap := makeSourceDestinationMap(scanner)

	scanner.Scan()
	locationHumidMap := makeSourceDestinationMap(scanner)

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := locationHumidMap.GetDestination(
			humidTempMap.GetDestination(
				tempLightMap.GetDestination(
					lightWaterMap.GetDestination(
						waterFertilizerMap.GetDestination(
							fertilizerSoilMap.GetDestination(soilSeedMap.GetDestination(seed)),
						),
					),
				),
			),
		) // oh yeah baby, this is REAL programming

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println(lowestLocation)
}

func main() {
	part_one()
	part_two()
}
