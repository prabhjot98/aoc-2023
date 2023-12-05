package util_test

import (
	"day4/util"
	"reflect"
	"testing"
)

func TestPartGetLength(t *testing.T) {
	line1 := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	line2 := "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"

	e1 := util.ScratchCard{
		1,
		[]util.Number{41, 48, 83, 86, 17},
		[]util.Number{83, 86, 6, 31, 17, 9, 48, 53},
	}
	e2 := util.ScratchCard{
		4,
		[]util.Number{41, 92, 73, 84, 69},
		[]util.Number{59, 84, 76, 51, 58, 5, 54, 83},
	}

	if !reflect.DeepEqual(util.ParseLineToScratchCard(line1), e1) {
		t.Errorf("TestParseLineToScratchCard: %+v %+v", line1, e1)
	}

	if !reflect.DeepEqual(util.ParseLineToScratchCard(line2), e2) {
		t.Errorf("TestParseLineToScratchCard: %+v %+v", line2, e2)
	}
}

func TestGetMatches(t *testing.T) {
	card1 := util.ScratchCard{
		1,
		[]util.Number{41, 48, 83, 86, 17},
		[]util.Number{83, 86, 6, 31, 17, 9, 48, 53},
	}
	e1 := 4

	if card1.GetMatches() != e1 {
		t.Errorf("TestGetMatches: expected %d got %d", 4, card1.GetMatches())
	}
}
