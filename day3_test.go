package aoc2023

import "testing"

func TestSolveDay3ASample(t *testing.T) {
	input := ReadInputFile("day3_sample.txt")
	m := InputAsMatrix(input)
	actual := SolveDay3A(m)
	AssertEquals(actual, 4361, t)
}

func TestSolveDay3A(t *testing.T) {
	input := ReadInputFile("day3.txt")
	m := InputAsMatrix(input)
	actual := SolveDay3A(m)
	AssertEquals(actual, 498559, t)
}

func TestSolveDay3BSample(t *testing.T) {
	input := ReadInputFile("day3_sample.txt")
	m := InputAsMatrix(input)
	actual := SolveDay3B(m)
	AssertEquals(actual, 467835, t)
}

func TestSolveDay3B(t *testing.T) {
	input := ReadInputFile("day3.txt")
	m := InputAsMatrix(input)
	actual := SolveDay3B(m)
	AssertEquals(actual, 72246648, t)
}
