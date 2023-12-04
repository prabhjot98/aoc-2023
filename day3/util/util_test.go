package util_test

import (
	"day3/util"
	"fmt"
	"slices"
	"testing"
)

func TestParseRound(t *testing.T) {
	testEngineSchema := `
  467..114..
  ...*......
  ..35..633.
  ......#...
  617*......
  .....+.58.
  ..592.....
  ......755.
  ...$.*....
  .664.598..
  `
	fmt.Println(testEngineSchema)
}

func TestPartGetLength(t *testing.T) {
	pn := util.PartNumberId(1234)
	pnLen := pn.GetLength()
	expected := 4
	if expected != pnLen {
		t.Errorf("GetLength(%d) Expected %d but got %d", pn, expected, pnLen)
	}

	pn = util.PartNumberId(1)
	pnLen = pn.GetLength()
	expected = 1
	if expected != pnLen {
		t.Errorf("GetLength(%d) Expected %d but got %d", pn, expected, pnLen)
	}
}

func TestParsePartNumbers(t *testing.T) {
	line := "12..@.*.35.+12.&.11"
	expected := []util.PartNumberId{12, 35, 12, 11}
	expectedIndices := []int{0, 8, 12, 17}
	pns, indices := util.ParsePartNumbers(line)
	if !slices.Equal(pns, expected) {
		t.Errorf("ParsePartNumbers(%s) Expected %d but got %d", line, expected, pns)
	}
	if !slices.Equal(expectedIndices, indices) {
		t.Errorf("ParsePartNumbers(%s) Expected %d but got %d", line, expectedIndices, indices)
	}
}

func TestParseSymbols(t *testing.T) {
	line := "12..@.*.35.+12.&.11"
	expected := []string{"@", "*", "+", "&"}
	expectedIndices := []int{4, 6, 11, 15}
	pns, indices := util.ParseSymbols(line)
	if !slices.Equal(pns, expected) {
		t.Errorf("ParseSymbols(%s) Expected %s but got %s", line, expected, pns)
	}
	if !slices.Equal(expectedIndices, indices) {
		t.Errorf("ParseSymbols(%s) Expected %d but got %d", line, expectedIndices, indices)
	}
}

func TestIsAdjacentTo(t *testing.T) {
	s1 := util.Symbol{util.Coord{1, 1}}
	s2 := util.Symbol{util.Coord{4, 3}}
	s3 := util.Symbol{util.Coord{3, 2}}
	p1 := util.Part{util.Coord{0, 0}, 1}
	p2 := util.Part{util.Coord{0, 3}, 100}

	//   01234
	// 0 1....
	// 1 .$...
	// 2 ...!.
	// 3 100.@

	if !s1.IsAdjacentTo(p1) {
		t.Errorf("IsAdjacentTo() Expected to be adjacent")
	}
	if s1.IsAdjacentTo(p2) {
		t.Errorf("IsAdjacentTo() Expected not to be adjacent")
	}
	if s2.IsAdjacentTo(p2) {
		t.Errorf("IsAdjacentTo() Expected not to be adjacent")
	}
	if !s3.IsAdjacentTo(p2) {
		t.Errorf("IsAdjacentTo() Expected s3 to be adjacent")
	}
}
