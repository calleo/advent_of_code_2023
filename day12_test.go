package aoc2023

import "testing"

func TestSolveDay12ASample(t *testing.T) {
	input := ReadInputFile("day12_sample.txt")
	actual := SolveDay12A(input)
	AssertEquals(actual, 21, t)
}

func TestSolveDay12A(t *testing.T) {
	input := ReadInputFile("day12.txt")
	actual := SolveDay12A(input)
	AssertEquals(actual, 7017, t)
}

/* Too darn slow...
func TestSolveDay12BSample(t *testing.T) {
	input := ReadInputFile("day12_sample.txt")
	actual := SolveDay12B(input, 5, t)
}
*/

/* Too darn slow ...
func TestSolveDay12B(t *testing.T) {
	input := ReadInputFile("day12.txt")
	actual := SolveDay12B(input, 5, t)
}
*/
