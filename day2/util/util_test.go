package util_test

import (
	"day2/util"
	"encoding/json"
	"testing"
)

func TestParseRound(t *testing.T) {
	expected := util.Round{4, 0, 3}
	roundStr := " 3 blue, 4 red"
	round := util.ParseRound(roundStr)

	if round != expected {
		e, _ := json.Marshal(expected)
		r, _ := json.Marshal(round)
		t.Errorf("ParseRound(%s) expected %s got %s ", roundStr, e, r)
	}
}

func TestParseToken(t *testing.T) {
	eCount := 3
	eColor := "blue"
	token := " 3 blue"
	count, color := util.ParseToken(token)

	if count != eCount || eColor != color {
		t.Errorf("ParseToken(%s) expected %d %s got %d %s ", token, eCount, eColor, count, color)
	}
}

func TestIsValidGame(t *testing.T) {
	testCases := []struct {
		game_state  string
		red_cubes   int
		green_cubes int
		blue_cubes  int
		expected    bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 12, 13, 14, true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12, 13, 14, true},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			12,
			13,
			14,
			false,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			12,
			13,
			14,
			false,
		},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 12, 13, 14, true},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.game_state,
			func(t *testing.T) {
				result, _ := util.IsValidGame(
					util.LineToGame(testCase.game_state),
					testCase.red_cubes,
					testCase.green_cubes,
					testCase.blue_cubes)

				if result != testCase.expected {
					t.Errorf(
						"LineToDigit(%s) expected %t but got %t",
						testCase.game_state,
						testCase.expected,
						result,
					)
				}
			},
		)
	}
}
