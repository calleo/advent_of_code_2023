package aoc2023

import "testing"

func TestSolveDay11A(t *testing.T) {
	input := ReadInputFile("day11_sample.txt")
	actual := SolveDay11A(input)
	AssertEquals(actual, 1, t)
}
