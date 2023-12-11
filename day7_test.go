package aoc2023

import (
	"golang.org/x/exp/slices"
	"testing"
)

func TestSolveDay7ASample(t *testing.T) {
	input := ReadInputFile("day7_sample.txt")
	actual := SolveDay7A(input)
	AssertEquals(actual, 6440, t)
}

func TestSolveDay7ASampleMartin(t *testing.T) {
	input := []string{"84444 200", "29999 100"}
	actual := SolveDay7A(input)
	AssertEquals(actual, 500, t)
}

func TestSolveDay7A(t *testing.T) {
	input := ReadInputFile("day7.txt")
	actual := SolveDay7A(input)
	AssertEquals(actual, 251287184, t)
}

func TestSolveDay7BSample(t *testing.T) {
	input := ReadInputFile("day7_sample.txt")
	actual := SolveDay7B(input)
	AssertEquals(actual, 5905, t)
}

func TestSolveDay7B(t *testing.T) {
	input := ReadInputFile("day7.txt")
	actual := SolveDay7B(input)
	//AssertEquals(actual, 249573581, t) // too low
	AssertEquals(actual, 250608479, t) // too low
	//AssertEquals(actual, 251287184, t) // too high
	//AssertEquals(actual, 250660764, t) // incorrect
}

func TestGetValueWithJokers1(t *testing.T) {
	hand := Hand{cards: []int{13, 10, 11, 11, 10}}
	actual := hand.GetValueWithJokers(true)
	expected := "ZAZZZZMPYYP"
	if actual != expected {
		t.Fatalf("Value %s does not equal %s", actual, expected)
	}
}

func TestGetValueWithJokers2(t *testing.T) {
	hand := Hand{cards: []int{13, 10, 9, 11, 10}}
	actual := hand.GetValueWithJokers(true)
	expected := "ZZZAZZMPQYP"
	if actual != expected {
		t.Fatalf("Value %s does not equal %s", actual, expected)
	}
}

func TestGetValueWithJokers3(t *testing.T) {
	hand := Hand{cards: []int{13, 10, 11, 11, 11}}
	actual := hand.GetValueWithJokers(true)
	expected := "ZAZZZZMPYYY"
	if actual != expected {
		t.Fatalf("Value %s does not equal %s", actual, expected)
	}
}

func TestHandSort(t *testing.T) {
	hands := []Hand{
		{cards: []int{11, 2, 2, 3, 3}},
		{cards: []int{3, 2, 3, 2, 3}},
		{cards: []int{2, 2, 4, 5, 6}},
		{cards: []int{2, 2, 3, 4, 2}},
		{cards: []int{11, 2, 3, 4, 5}},
	}
	sortedHands := SortHands(hands, true)

	if !slices.Equal(sortedHands[0].cards, []int{11, 2, 3, 4, 5}) {
		t.Fatalf("0 Not equal!!")
	}
	if !slices.Equal(sortedHands[1].cards, []int{2, 2, 4, 5, 6}) {
		t.Fatalf("1 Not equal!!")
	}
	if !slices.Equal(sortedHands[2].cards, []int{2, 2, 3, 4, 2}) {
		t.Fatalf("1 Not equal!!")
	}
	if !slices.Equal(sortedHands[3].cards, []int{11, 2, 2, 3, 3}) {
		t.Fatalf("1 Not equal!!")
	}
	if !slices.Equal(sortedHands[4].cards, []int{3, 2, 3, 2, 3}) {
		t.Fatalf("1 Not equal!!")
	}
}
