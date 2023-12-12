package aoc2023

import "strings"

func Empty(line string) bool {
	for _, ch := range strings.Split(line, "") {
		if ch != "." {
			return false
		}
	}
	return true
}

func SolveDay11A(input []string) int {
	var gravity []string
	var gravityTransposed []string

	for _, line := range input {
		gravity = append(gravity, line)
		if Empty(line) {
			gravity = append(gravity, line)
		}
	}

	for col := 0; col < len(gravity[0])-1; col++ {
		line := ""
		for row := 0; row < len(gravity)-1; row++ {
			line = line + string(gravity[row][col])
		}
		gravityTransposed = append(gravityTransposed, line)
	}

	return 1
}
