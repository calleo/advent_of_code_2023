package aoc2023

import "testing"

func TestSolveDay6ASample(t *testing.T) {
	input := ReadInputFile("day6_sample.txt")
	actual := SolveDay6A(input)
	AssertEquals(actual, 288, t)
}

func TestSolveDay6A(t *testing.T) {
	input := ReadInputFile("day6.txt")
	actual := SolveDay6A(input)
	AssertEquals(actual, 3316275, t)
}

func TestSolveDay6BSample(t *testing.T) {
	input := ReadInputFile("day6_sample.txt")
	actual := SolveDay6B(input)
	AssertEquals(actual, 71503, t)
}

func TestSolveDay6B(t *testing.T) {
	input := ReadInputFile("day6.txt")
	actual := SolveDay6B(input)
	AssertEquals(actual, 27102791, t)
}
