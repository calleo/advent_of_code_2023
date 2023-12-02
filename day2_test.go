package aoc2023

import (
	"testing"
)

func TestSolveDay2ASample(t *testing.T) {
	expected := 8
	input := ReadInputFile("day2_sample.txt")
	actual := SolveDay2A(input)
	AssertEquals(actual, expected, t)
}

func TestSolveDay2A(t *testing.T) {
	expected := 2416
	input := ReadInputFile("day2.txt")
	actual := SolveDay2A(input)
	AssertEquals(actual, expected, t)
}

func TestSolveDay2BSample(t *testing.T) {
	expected := 2286
	input := ReadInputFile("day2_sample.txt")
	actual := SolveDay2B(input)
	AssertEquals(actual, expected, t)
}

func TestSolveDay2B(t *testing.T) {
	expected := 63307
	input := ReadInputFile("day2.txt")
	actual := SolveDay2B(input)
	AssertEquals(actual, expected, t)
}

func TestParseInput(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	}
	actual := ParseInput(input)
	AssertEquals(len(actual), 2, t)
	AssertEquals(actual[0].sets[0].blue, 3, t)
	AssertEquals(actual[1].sets[1].green, 3, t)
}
