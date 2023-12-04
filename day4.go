package aoc2023

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math"
	"strconv"
	"strings"
)

type Scratchcard struct {
	winners []int
	numbers []int
	copies  int
}

func parseScratchcards(line string) Scratchcard {
	parts := strings.Split(line, ":")
	sets := strings.Split(parts[1], "|")

	parse := func(s string) []int {
		var numbers []int
		for _, nr := range strings.Split(s, " ") {
			nrClean := strings.Trim(nr, " ")
			nrInt, ok := strconv.Atoi(nrClean)
			if ok == nil {
				numbers = append(numbers, nrInt)
			}
		}
		return numbers
	}

	return Scratchcard{winners: parse(sets[0]), numbers: parse(sets[1]), copies: 1}
}

func SolveDay4A(input []string) int {
	score := 0
	for _, cardGame := range input {
		card := parseScratchcards(cardGame)
		result := mapset.NewSet[int](card.winners...).Intersect(mapset.NewSet[int](card.numbers...))
		score += int(math.Pow(2, float64(len(result.ToSlice())-1)))
	}
	return score
}

func SolveDay4B(input []string) int {
	currentCard := 0
	var cards []Scratchcard

	for _, cardGame := range input {
		cards = append(cards, parseScratchcards(cardGame))
	}

	for {
		if currentCard >= len(cards) {
			break
		}

		card := cards[currentCard]
		result := mapset.NewSet[int](card.winners...).Intersect(mapset.NewSet[int](card.numbers...))
		matches := len(result.ToSlice())

		for i := currentCard + 1; i <= currentCard+matches; i++ {
			cards[i].copies += 1 * card.copies
		}
		currentCard += 1
	}

	sum := 0
	for _, card := range cards {
		sum += card.copies
	}

	return sum
}
