package aoc2023

import "strings"

func Hash(s string) int {
	hash := 0

	for _, char := range strings.Split(s, "") {
		hash += int(rune(char[0]))
		hash *= 17
		hash %= 256
	}
	return hash
}

func SolveDay15(input []string) int {
	sum := 0
	for _, instr := range strings.Split(input[0], ",") {
		sum += Hash(instr)
	}
	return sum
}
