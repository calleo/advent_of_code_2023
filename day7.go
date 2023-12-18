package aoc2023

import (
	"fmt"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards         []int
	bid           int
	rank          int
	cardsOriginal string
}

func (h *Hand) GetCardOccurrences() map[int]int {
	occurrences := make(map[int]int)
	for _, card := range h.cards {
		occurrences[card]++
	}
	return occurrences
}

func CountOccurrencesOf(slice []int, value int) int {
	sum := 0
	for _, val := range slice {
		if val == value {
			sum++
		}
	}
	return sum
}

func (h *Hand) GetValueWithJokers(jokers bool) string {
	var occurrences = h.GetCardOccurrences()
	nrOfJokers := 0
	if jokers {
		nrOfJokers = occurrences[11]
		delete(occurrences, 11)
	}
	values := maps.Values(occurrences)
	nrOfPairs := CountOccurrencesOf(values, 2)
	value := strings.Split("ZZZZZZZZZZZ", "")

	if slices.Contains(values, 5) || (slices.Contains(values, 4) && nrOfJokers == 1) || (slices.Contains(values, 3) && nrOfJokers == 2) || nrOfJokers == 5 || nrOfJokers == 4 || (nrOfJokers == 3 && nrOfPairs == 1) {
		value[0] = "A"
	} else if slices.Contains(values, 4) || (slices.Contains(values, 3) && nrOfJokers == 1) || (slices.Contains(values, 2) && nrOfJokers == 2) || (slices.Contains(values, 1) && nrOfJokers == 3) || nrOfJokers == 4 {
		value[1] = "A"
	} else if (slices.Contains(values, 3) && slices.Contains(values, 2)) || (nrOfPairs == 2 && nrOfJokers == 1) {
		value[2] = "A" // Full house
	} else if slices.Contains(values, 3) || (slices.Contains(values, 2) && nrOfJokers == 1) || (slices.Contains(values, 1) && nrOfJokers == 2) || nrOfJokers == 3 {
		value[3] = "A" // Three of a kind
	} else if nrOfPairs == 2 || (nrOfPairs == 1 && nrOfJokers == 1) || (nrOfJokers == 2) {
		value[4] = "A"
	} else if nrOfPairs == 1 || nrOfJokers == 1 {
		value[5] = "A"
	}

	for pos, card := range h.cards {
		if card == 11 && jokers {
			card = 1
		}
		value[5+1+pos] = string(rune(90 - card))
	}

	return strings.Join(value, "")
}

func ParseHands(input []string) []Hand {
	var hands []Hand

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		var cards []int
		for _, card := range splitLine[0] {
			cardAsInt := 0
			switch string(card) {
			case "A":
				cardAsInt = 14
			case "K":
				cardAsInt = 13
			case "Q":
				cardAsInt = 12
			case "J":
				cardAsInt = 11
			case "T":
				cardAsInt = 10
			default:
				cardAsInt, _ = strconv.Atoi(string(card))
			}
			cards = append(cards, cardAsInt)
		}
		bid, _ := strconv.Atoi(splitLine[1])
		hands = append(hands, Hand{cards: cards, bid: bid, cardsOriginal: splitLine[0]})
	}

	return hands
}

func SortHands(hands []Hand, jokers bool) []Hand {
	less := func(a, b int) bool {
		aValue := hands[a].GetValueWithJokers(jokers)
		bValue := hands[b].GetValueWithJokers(jokers)

		if aValue == bValue {
			panic(fmt.Sprintf("Found equal hands: %s", aValue))
		}

		return aValue > bValue
	}
	sort.Slice(hands, less)
	return hands
}

func SolveDay7A(input []string) int {
	winnings := 0
	hands := SortHands(ParseHands(input), false)
	for rank, hand := range hands {
		winnings += hand.bid * (rank + 1)
	}
	return winnings
}

func SolveDay7B(input []string) int {
	winnings := 0
	hands := SortHands(ParseHands(input), true)
	for rank, hand := range hands {
		winnings += hand.bid * (rank + 1)
	}
	return winnings
}
