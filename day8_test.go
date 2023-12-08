package aoc2023

import (
	"testing"
)

func TestSolveDay8ASample(t *testing.T) {
	input := ReadInputFile("day8_sample.txt")
	actual := SolveDay8A(input)
	AssertEquals(actual, 2, t)
}

func TestSolveDay8A(t *testing.T) {
	input := ReadInputFile("day8.txt")
	actual := SolveDay8A(input)
	AssertEquals(actual, 17141, t)
}

/*func TestSolveDay8BSample(t *testing.T) {
	input := ReadInputFile("day8b_sample.txt")
	actual := SolveDay8B(input, t)
	AssertEquals(actual, 6, t)
}*/
