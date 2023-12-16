package aoc2023

import (
	"strings"
)

type Pattern struct {
	rows    []string
	columns []string
}

func FillColumns(p *Pattern) {
	cols := make([]string, len(p.rows[0]))

	for _, line := range p.rows {
		for pos, char := range strings.Split(line, "") {
			cols[pos] = cols[pos] + char
		}
	}

	p.columns = cols
}

func ParsePatterns(input []string) []*Pattern {
	var patterns []*Pattern
	pattern := &Pattern{}

	for _, line := range input {
		if line == "" {
			FillColumns(pattern)
			patterns = append(patterns, pattern)
			pattern = &Pattern{}
			continue
		}
		pattern.rows = append(pattern.rows, line)
	}

	FillColumns(pattern)
	patterns = append(patterns, pattern)

	return patterns
}

func FindReflection(pattern []string, skip int) int {
	for index, line := range pattern {
		if index == 0 || index == skip {
			continue
		}
		if line == pattern[index-1] {
			// we found a potential match, now check all the others
			found := true
			for bw := index - 2; bw >= 0; bw-- {
				fw := index + index - bw - 1
				if fw >= len(pattern) {
					break
				}
				if pattern[bw] != pattern[fw] {
					found = false
					break
				}
			}
			if found {
				return index
			}
		}
	}
	return 0
}

func FindSmudgedReflection(lines []string) int {
	reflectionValue := 0
	existingRowReflection := FindReflection(lines, -1)
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			old := lines[row]
			lines[row] = ReplaceAt(lines[row], col)
			reflectionValue = FindReflection(lines, existingRowReflection)
			lines[row] = old
			if reflectionValue != 0 {
				return reflectionValue
			}
		}
	}
	return 0
}

func ReplaceAt(s string, at int) string {
	chars := strings.Split(s, "")
	if chars[at] == "." {
		chars[at] = "#"
	} else {
		chars[at] = "."
	}
	return strings.Join(chars, "")
}

func SolveDay13A(input []string) int {
	patterns := ParsePatterns(input)
	sum := 0

	for _, pattern := range patterns {
		reflectionValue := FindReflection(pattern.rows, -1) * 100
		if reflectionValue == 0 {
			reflectionValue = FindReflection(pattern.columns, -1)
		}
		sum += reflectionValue
	}

	return sum
}

func SolveDay13B(input []string) int {
	patterns := ParsePatterns(input)
	sum := 0

	for _, pattern := range patterns {
		reflectionValue := FindSmudgedReflection(pattern.rows) * 100
		if reflectionValue == 0 {
			reflectionValue = FindSmudgedReflection(pattern.columns)
		}
		sum += reflectionValue
	}

	return sum
}
