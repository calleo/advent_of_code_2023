package aoc2023

import (
	"testing"
)

func TestSolveDay7ASample(t *testing.T) {
	input := ReadInputFile("day7_sample.txt")
	actual := SolveDay7A(input)
	AssertEquals(actual, 6440, t)
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
	AssertEquals(actual, 250757288, t)
}
