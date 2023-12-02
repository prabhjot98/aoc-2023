package day1_test

import (
	"day1"
	"testing"
)

func TestLineToDigit(t *testing.T) {
	testCases := []struct {
		arg1     string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, testCase := range testCases {
		t.Run(testCase.arg1, func(t *testing.T) {
			result := day1.LineToDigit(testCase.arg1)
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
