package main

import (
	"bufio"
	"day6/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part_one() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	durationStr := strings.Split(scanner.Text(), ":")[1]
	durationStrs := strings.Fields(durationStr)

	scanner.Scan()
	distanceStr := strings.Split(scanner.Text(), ":")[1]
	distanceStrs := strings.Fields(distanceStr)

	var races []util.Race

	for i := range durationStrs {
		duration, _ := strconv.Atoi(durationStrs[i])
		distance, _ := strconv.Atoi(distanceStrs[i])

		races = append(races, util.Race{Duration: duration, Distance: distance})
	}

	total := 1
	for _, race := range races {
		sum := 0
		for i := 0; i <= race.Duration; i++ {
			if race.BeatsRecord(i) {
				sum++
			}
		}
		total *= sum
	}
	fmt.Println(total)
}

func part_two() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	durationStr := strings.Split(scanner.Text(), ":")[1]
	durationStrs := strings.Join(strings.Fields(durationStr), "")

	scanner.Scan()
	distanceStr := strings.Split(scanner.Text(), ":")[1]
	distanceStrs := strings.Join(strings.Fields(distanceStr), "")

	duration, _ := strconv.Atoi(durationStrs)
	distance, _ := strconv.Atoi(distanceStrs)

	race := util.Race{Duration: duration, Distance: distance}

	sum := 0
	for i := 0; i <= race.Duration; i++ {
		if race.BeatsRecord(i) {
			sum++
		}
	}
	fmt.Println(sum)
}

func main() {
	part_one()
	part_two()

}
