package aoc2023

import "testing"

func TestSolveDay4ASample(t *testing.T) {
	input := ReadInputFile("day4_sample.txt")
	actual := SolveDay4A(input)
	AssertEquals(actual, 13, t)
}

func TestSolveDay4A(t *testing.T) {
	input := ReadInputFile("day4.txt")
	actual := SolveDay4A(input)
	AssertEquals(actual, 22897, t)
}

func TestSolveDay4BSample(t *testing.T) {
	input := ReadInputFile("day4_sample.txt")
	actual := SolveDay4B(input)
	AssertEquals(actual, 30, t)
}

func TestSolveDay4B(t *testing.T) {
	input := ReadInputFile("day4.txt")
	actual := SolveDay4B(input)
	AssertEquals(actual, 5095824, t)
}
