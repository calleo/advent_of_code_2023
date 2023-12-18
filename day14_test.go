package aoc2023

import "testing"

func TestSolveDay14ASample(t *testing.T) {
	input := ReadInputFile("day14_sample.txt")
	actual := SolveDay14A(input)
	AssertEquals(actual, 136, t)
}

func TestSolveDay14A(t *testing.T) {
	input := ReadInputFile("day14.txt")
	actual := SolveDay14A(input)
	AssertEquals(actual, 108826, t)
}

func TestSolveDay14BSample(t *testing.T) {
	input := ReadInputFile("day14_sample.txt")
	actual := SolveDay14B(input)
	AssertEquals(actual, 64, t)
}

func TestSolveDay14B(t *testing.T) {
	input := ReadInputFile("day14.txt")
	actual := SolveDay14B(input)
	AssertEquals(actual, 99291, t)
}
