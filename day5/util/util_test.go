package util_test

import (
	"day5/util"
	"testing"
)

func TestGetDestination(t *testing.T) {
	type testCase struct {
		source   int
		expected int
	}

	testCases := []testCase{
		{0, 0},
		{1, 1},
		{48, 48},
		{49, 49},
		{50, 52},
		{51, 53},
		{96, 98},
		{97, 99},
		{98, 50},
		{99, 51},
	}

	seedToSoilMap := util.DestinationSourceMap{
		[]util.MapLine{{50, 98, 2}, {52, 50, 48}},
	}

	for _, testCase := range testCases {
		if seedToSoilMap.GetDestination(testCase.source) != testCase.expected {
			t.Errorf(
				"GetDestination(%d) expected %d got %d",
				testCase.source,
				testCase.expected,
				seedToSoilMap.GetDestination(testCase.source),
			)
		}
	}
}
