package aoc2023

import (
	"testing"
)

/*
func TestSolveDay10Utils(t *testing.T) {
	var seen []Point
	input := ReadInputFile("day10_sample.txt")
	matrix := InputAsMatrix(input)

	actual := GetStart(matrix)
	AssertEquals(actual.value, "S", t)

	east, _ := matrix.eastOf(actual)
	AssertEquals(east.value, "-", t)

	west, _ := matrix.westOf(actual)
	AssertEquals(west.value, ".", t)

	north, _ := matrix.northOf(actual)
	AssertEquals(north.value, ".", t)

	south, _ := matrix.southOf(actual)
	AssertEquals(south.value, "|", t)

	// Test traversal
	var next []Point
	next = append(next, actual)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "-", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "7", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "|", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "J", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "-", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "L", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(next[0].value, "|", t)
	seen = append(seen, next[0])

	next = GetNext(matrix, next[0], seen)
	AssertEquals(len(next), 0, t)
}
*/

func TestSolveDay10Sample(t *testing.T) {
	input := ReadInputFile("day10_sample.txt")
	actual := SolveDay10A(input)
	AssertEquals(actual, 4, t)
}

func TestSolveDay10Sample2(t *testing.T) {
	input := ReadInputFile("day10_sample2.txt")
	actual := SolveDay10A(input)
	AssertEquals(actual, 8, t)
}

func TestSolveDay10A(t *testing.T) {
	input := ReadInputFile("day10.txt")
	actual := SolveDay10A(input)
	AssertEquals(actual, 6897, t)
}

/*
these examples doesn't work due to the way how

	'S' is being handled and counted
*/
func TestSolveDay10BSample(t *testing.T) {
	input := ReadInputFile("day10_sample.txt")
	actual := SolveDay10B(input)
	AssertEquals(actual, 1, t)
}

func TestSolveDay10BSampleB(t *testing.T) {
	input := ReadInputFile("day10_sample3.txt")
	actual := SolveDay10B(input)
	AssertEquals(actual, 4, t)
}

func TestSolveDay10BSampleC(t *testing.T) {
	input := ReadInputFile("day10_sample4.txt")
	actual := SolveDay10B(input)
	AssertEquals(actual, 8, t)
}

func TestSolveDay10BSampleD(t *testing.T) {
	input := ReadInputFile("day10_sample5.txt")
	actual := SolveDay10B(input)
	AssertEquals(actual, 10, t)
}

func TestSolveDay10BSampleE(t *testing.T) {
	input := ReadInputFile("day10_sample6.txt")
	actual := SolveDay10B(input)
	AssertEquals(actual, 4, t)
}

func TestSolveDay10B(t *testing.T) {
	input := ReadInputFile("day10.txt")
	actual := SolveDay10B(input)
	AssertEquals(actual, 367, t)
}
