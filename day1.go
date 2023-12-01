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
		old := strings.Clone(line)
		document[index] = ReplaceNumericDigits(line)
		fmt.Printf("%s => %s\n", old, document[index])
	}
	return SolveDay1A(document)
}

func ReplaceNumericDigits(line string) string {
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

	to := 1
	for to <= len(line) {
		for from := 0; from <= to; from++ {
			nr := line[from:to]
			translation, ok := lookup[nr]
			if ok {
				line = strings.Replace(line, nr, translation, 1)
				to = 0 // start over, if something has been replaced since string length changes
				break
			}
		}
		to = to + 1
	}

	return line
}
