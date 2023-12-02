package part1

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
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

func getFirstDigit(line string) int {
	for _, rune := range line {
		if unicode.IsDigit(rune) {
			return RuneToInt(rune)
		}
	}
	return 0
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		rune := rune(line[i])
		if unicode.IsDigit(rune) {
			return RuneToInt(rune)
		}
	}
	return 0
}

func LineToDigit(line string) int {
	firstDigit := getFirstDigit(line)
	secondDigit := getLastDigit(line)
	return (firstDigit * 10) + secondDigit
}

func CoolWay() int {
	list := fileToString("input.txt")
	lines := strings.Split(list, "\n")

	totalCalibrationValues := 0
	totalLines := 0

	digitChan := make(chan int)

	for _, line := range lines {
		go func(line string) {
			digitChan <- LineToDigit(line)
		}(line)
		totalLines++
	}

	for i := 0; i < totalLines; i++ {
		totalCalibrationValues += <-digitChan
	}

	return totalCalibrationValues
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

func MeasureTime(fn func() int) int {
	start := time.Now()
	rt := fn()
	duration := time.Since(start)
	fmt.Printf("Took %f seconds to perform the function\n", duration.Seconds())
	return rt
}
