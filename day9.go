package aoc2023

import (
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func ExtrapolateReadingsSum(readings []int) int {
	var extrapolatedReadings [][]int
	extrapolatedReadings = append(extrapolatedReadings, readings)

	for {
		var extrapolated []int
		allZero := true
		for i := 0; i < len(readings)-1; i++ {
			diff := readings[i+1] - readings[i]
			extrapolated = append(extrapolated, diff)
			if diff != 0 {
				allZero = false
			}
		}
		extrapolatedReadings = append(extrapolatedReadings, extrapolated)
		readings = extrapolated
		if allZero {
			break
		}
	}

	var newReadings []int
	previous := 0
	for i := len(extrapolatedReadings) - 2; i >= 0; i-- {
		last := extrapolatedReadings[i][len(extrapolatedReadings[i])-1]
		newReadings = append(newReadings, last+previous)
		previous = last + previous
	}

	return newReadings[len(newReadings)-1]
}

func SolveDay9A(input []string, reverse bool) int {
	var sensorReadings [][]int

	for _, line := range input {
		var readings []int
		for _, reading := range strings.Split(line, " ") {
			nr, _ := strconv.Atoi(reading)
			readings = append(readings, nr)
		}
		if reverse {
			slices.Reverse(readings)
		}
		sensorReadings = append(sensorReadings, readings)
	}

	sum := 0
	for _, reading := range sensorReadings {
		sum += ExtrapolateReadingsSum(reading)
	}

	return sum
}

func SolveDay9B(input []string) int {
	return SolveDay9A(input, true)
}
