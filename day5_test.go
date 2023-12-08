package aoc2023

import (
	"math"
	"testing"
)

func TestSolveDay5ASample(t *testing.T) {
	input := ReadInputFile("day5_sample.txt")
	actual := SolveDay5A(input)
	AssertEquals(actual, 35, t)
}

func TestSolveDay5A(t *testing.T) {
	input := ReadInputFile("day5.txt")
	actual := SolveDay5A(input)
	AssertEquals(actual, 318728750, t)
}

func TestSolveDay5BSample(t *testing.T) {
	input := ReadInputFile("day5_sample.txt")
	actual := SolveDay5B(input)
	AssertEquals(actual, 46, t)
}

/* This takes a long time, ~3 minutes */
/*func TestSolveDay5B(t *testing.T) {
	input := ReadInputFile("day5.txt")
	actual := SolveDay5B(input)
	AssertEquals(actual, 37384986, t)
}*/

func TestParseAlmanac(t *testing.T) {
	input := ReadInputFile("day5_sample.txt")
	almanac := ParseAlmanac(input)
	AssertEquals(len(almanac.seeds), 4, t)
	AssertEquals(len(almanac.mappers), 7, t)
	AssertEquals(len(almanac.mappers[2]), 4, t)
	AssertEquals(len(almanac.mappers[2][3]), 3, t)
}

func TestGetMappingSample(t *testing.T) {
	input := ReadInputFile("day5_sample.txt")
	almanac := ParseAlmanac(input)
	min := math.MaxInt

	for _, mapping := range almanac.mappers {
		actual := GetMapping(57, mapping)
		min = int(math.Min(float64(min), float64(actual)))
	}

	AssertEquals(min, 50, t)
}

func TestGetMapping(t *testing.T) {
	var mappings1 [][]int
	actual1 := GetMapping(10, mappings1)
	AssertEquals(actual1, 10, t)

	var mappings2 = [][]int{{50, 98, 2}, {52, 50, 48}}
	actual2 := GetMapping(98, mappings2)
	AssertEquals(actual2, 50, t)

	var mappings3 = [][]int{{50, 98, 2}, {52, 50, 48}}
	actual3 := GetMapping(99, mappings3)
	AssertEquals(actual3, 51, t)

	actual6 := GetMapping(75, mappings3)
	AssertEquals(actual6, 77, t)

	var mappings4 = [][]int{{90, 50, 10}}
	actual4 := GetMapping(59, mappings4)
	AssertEquals(actual4, 99, t)

	actual5 := GetMapping(60, mappings4)
	AssertEquals(actual5, 60, t)
}
