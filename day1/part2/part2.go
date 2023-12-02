package part2

import (
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	One   = "one"
	Two   = "two"
	Three = "three"
	Four  = "four"
	Five  = "five"
	Six   = "six"
	Seven = "seven"
	Eight = "eight"
	Nine  = "nine"
)

func fileToString(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Could not find file name with name: %s", fileName)
	}
	return string(file)
}

func RuneToInt(r rune) int {
	return int(r) - '0'
}

func GetFirstDigit(line string) int {
	match := ""
	for _, rune := range line {
		if unicode.IsDigit(rune) {
			return RuneToInt(rune)
		}
		match = match + string(rune)
		if len(match) > 4 {
			last5 := match[len(match)-5:]
			switch last5 {
			case Three:
				return 3
			case Seven:
				return 7
			case Eight:
				return 8
			}
		}

		if len(match) > 3 {
			last4 := match[len(match)-4:]
			switch last4 {
			case Four:
				return 4
			case Five:
				return 5
			case Nine:
				return 9
			}
		}

		if len(match) > 2 {
			last3 := match[len(match)-3:]
			switch last3 {
			case One:
				return 1
			case Two:
				return 2
			case Six:
				return 6
			}
		}
	}
	return 0
}

func GetLastDigit(line string) int {
	match := ""
	for i := len(line) - 1; i >= 0; i-- {
		rune := rune(line[i])
		if unicode.IsDigit(rune) {
			return RuneToInt(rune)
		}
		match = string(line[i]) + match
		if len(match) > 4 {
			last5 := match[:5]
			switch last5 {
			case Three:
				return 3
			case Seven:
				return 7
			case Eight:
				return 8
			}
		}

		if len(match) > 3 {
			last4 := match[:4]
			switch last4 {
			case Four:
				return 4
			case Five:
				return 5
			case Nine:
				return 9
			}
		}

		if len(match) > 2 {
			last3 := match[:3]
			switch last3 {
			case One:
				return 1
			case Two:
				return 2
			case Six:
				return 6
			}
		}
	}
	return 0
}

func LineToDigit(line string) int {
	firstDigit := GetFirstDigit(line)
	lastDigit := GetLastDigit(line)
	return (firstDigit * 10) + lastDigit
}

func LameWay() int {
	list := fileToString("input.txt")
	lines := strings.Split(list, "\n")

	totalCalibrationValues := 0

	for _, line := range lines {
		totalCalibrationValues += LineToDigit(line)
	}

	return totalCalibrationValues
}
