package aoc2023

import "testing"

func TestSolveDay15Sample(t *testing.T) {
	input := ReadInputFile("day15_sample.txt")
	actual := SolveDay15(input)
	AssertEquals(actual, 1320, t)
}

func TestSolveDay15(t *testing.T) {
	input := ReadInputFile("day15.txt")
	actual := SolveDay15(input)
	AssertEquals(actual, 504449, t)
}

func TestHash(t *testing.T) {
	actual := Hash("HASH")
	AssertEquals(actual, 52, t)
}
