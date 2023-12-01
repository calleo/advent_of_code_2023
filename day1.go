package aoc2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolveDay1A(document []string) int {
	re := regexp.MustCompile("[0-9]")

	sum := 0
	for _, line := range document {
		matches := re.FindAllString(line, -1)
		if len(matches) > 0 {
			nr, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
			sum = sum + nr
		} else {
			panic(fmt.Sprintf("Line '%s' does not contain any numbers!", line))
		}
	}

	return sum
}

func SolveDay1B(document []string) int {
	for index, line := range document {
		document[index] = GetAllNumerical(line)
	}
	return SolveDay1A(document)
}

func GetAllNumerical(line string) string {
	lookup := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}
	var numerical []string

	for to := 1; to <= len(line); to++ {
		for from := 0; from <= to; from++ {
			nr := line[from:to]
			if len(nr) == 1 && strings.Contains("1234567890", nr) {
				numerical = append(numerical, nr)
			} else {
				translation, ok := lookup[nr]
				if ok {
					numerical = append(numerical, translation)
				}
			}
		}
	}

	return strings.Join(numerical, "")
}
