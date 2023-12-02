package part2_test

import (
	"day1/part2"
	"testing"
)

func TestGetFirstDigit(t *testing.T) {
	testCases := []struct {
		arg1     string
		expected int
	}{
		{"two1nine", 2},
		{"eightwothree", 8},
		{"abcone2threexyz", 1},
		{"xtwone3four", 2},
		{"4nineeightseven2", 4},
		{"zoneight234", 1},
		{"7pqrstsixteen", 7},
	}

	for _, testCase := range testCases {
		t.Run(testCase.arg1, func(t *testing.T) {
			result := part2.GetFirstDigit(testCase.arg1)
			if result != testCase.expected {
				t.Errorf(
					"GetFirstDigit(%s) expected %d but got %d",
					testCase.arg1,
					testCase.expected,
					result,
				)
			}
		})
	}
}

func TestGetLastDigit(t *testing.T) {
	testCases := []struct {
		arg1     string
		expected int
	}{
		{"two1nine", 9},
		{"eightwothree", 3},
		{"abcone2threexyz", 3},
		{"xtwone3four", 4},
		{"4nineeightseven2", 2},
		{"zoneight234", 4},
		{"7pqrstsixteen", 6},
	}

	for _, testCase := range testCases {
		t.Run(testCase.arg1, func(t *testing.T) {
			result := part2.GetLastDigit(testCase.arg1)
			if result != testCase.expected {
				t.Errorf(
					"GetLastDigit(%s) expected %d but got %d",
					testCase.arg1,
					testCase.expected,
					result,
				)
			}
		})
	}
}
func TestLineToDigit(t *testing.T) {
	testCases := []struct {
		arg1     string
		expected int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, testCase := range testCases {
		t.Run(testCase.arg1, func(t *testing.T) {
			result := part2.LineToDigit(testCase.arg1)
			if result != testCase.expected {
				t.Errorf(
					"LineToDigit(%s) expected %d but got %d",
					testCase.arg1,
					testCase.expected,
					result,
				)
			}
		})
	}
}
