package util_test

import (
	"day6/util"
	"testing"
)

func TestBeatsRecord(t *testing.T) {
	type TestCase struct {
		util.Race
		TimeButtonHeld int
		Expected       bool
	}

	testCases := []TestCase{
		{util.Race{7, 9}, 0, false},
		{util.Race{7, 9}, 1, false},
		{util.Race{7, 9}, 2, true},
		{util.Race{7, 9}, 3, true},
		{util.Race{7, 9}, 4, true},
		{util.Race{7, 9}, 5, true},
		{util.Race{7, 9}, 6, false},
		{util.Race{7, 9}, 7, false},
		{util.Race{7, 9}, 8, false},
		{util.Race{7, 9}, 9, false},
	}

	for _, testCase := range testCases {
		if testCase.Race.BeatsRecord(testCase.TimeButtonHeld) != testCase.Expected {
			t.Errorf(
				"BeatsRecord(%d) expected %t got %t",
				testCase.TimeButtonHeld,
				testCase.Expected,
				testCase.Race.BeatsRecord(testCase.TimeButtonHeld),
			)

		}

	}
}
