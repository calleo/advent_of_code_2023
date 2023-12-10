package aoc2023

import "testing"

func TestSolveDay9ASample(t *testing.T) {
	input := ReadInputFile("day9_sample.txt")
	actual := SolveDay9A(input, false)
	AssertEquals(actual, 114, t)
}

func TestSolveDay9A(t *testing.T) {
	input := ReadInputFile("day9.txt")
	actual := SolveDay9A(input, false)
	AssertEquals(actual, 2098530125, t)
}

func TestSolveDay9BSample(t *testing.T) {
	input := ReadInputFile("day9_sample.txt")
	actual := SolveDay9B(input)
	AssertEquals(actual, 2, t)
}

func TestSolveDay9B(t *testing.T) {
	input := ReadInputFile("day9.txt")
	actual := SolveDay9B(input)
	AssertEquals(actual, 1016, t)
}
