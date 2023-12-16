package aoc2023

import "testing"

func TestSolveDay13ASample(t *testing.T) {
	input := ReadInputFile("day13_sample.txt")
	actual := SolveDay13A(input)
	AssertEquals(actual, 405, t)
}

func TestSolveDay13A(t *testing.T) {
	input := ReadInputFile("day13.txt")
	actual := SolveDay13A(input)
	AssertEquals(actual, 33735, t)
}

func TestSolveDay13BSample(t *testing.T) {
	input := ReadInputFile("day13_sample.txt")
	actual := SolveDay13B(input)
	AssertEquals(actual, 400, t)
}

func TestSolveDay13B(t *testing.T) {
	input := ReadInputFile("day13.txt")
	actual := SolveDay13B(input)
	AssertEquals(actual, 38063, t)
}
