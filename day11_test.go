package aoc2023

import "testing"

func TestSolveDay11ASample(t *testing.T) {
	input := ReadInputFile("day11_sample.txt")
	actual := SolveDay11AB(input, 2)
	AssertEquals(actual, 374, t)
}

func TestSolveDay11A(t *testing.T) {
	input := ReadInputFile("day11.txt")
	actual := SolveDay11AB(input, 2)
	AssertEquals(actual, 9418609, t)
}

func TestSolveDay11BSample(t *testing.T) {
	input := ReadInputFile("day11_sample.txt")
	actual := SolveDay11AB(input, 10)
	AssertEquals(actual, 1030, t)
}

func TestSolveDay11B(t *testing.T) {
	input := ReadInputFile("day11.txt")
	actual := SolveDay11AB(input, 1000000)
	AssertEquals(actual, 593821230983, t)
}
