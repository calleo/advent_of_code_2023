package aoc2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type game struct {
	id   int
	sets []set
}

type set struct {
	blue  int
	red   int
	green int
}

func ParseInput(input []string) []game {
	games := []game{}
	for _, line := range input {
		gamePart := strings.Split(line, ":")[0]
		setsPart := strings.Split(line, ":")[1]
		id, _ := strconv.Atoi(strings.Split(gamePart, " ")[1])
		newGame := game{id: id}

		setsParts := strings.Split(setsPart, ";")

		for _, setParts := range setsParts {
			newSet := set{}
			for _, setPart := range strings.Split(setParts, ",") {
				setPart = strings.Trim(setPart, " ")
				countStr := strings.Split(setPart, " ")[0]
				count, _ := strconv.Atoi(strings.Trim(countStr, " "))
				color := strings.Trim(strings.Split(setPart, " ")[1], " ")
				switch color {
				case "green":
					newSet.green = count
				case "red":
					newSet.red = count
				case "blue":
					newSet.blue = count
				default:
					panic(fmt.Sprintf("Color: %s unknown", color))
				}
			}
			newGame.sets = append(newGame.sets, newSet)
		}
		games = append(games, newGame)
	}
	return games
}

func SolveDay2A(input []string) int {
	allowedRed := 12
	allowedGreen := 13
	allowedBlue := 14
	sum := 0

	games := ParseInput(input)

	for _, game := range games {
		possible := true

		for _, set := range game.sets {
			if set.red > allowedRed || set.blue > allowedBlue || set.green > allowedGreen {
				possible = false
			}
		}

		if possible {
			sum += game.id
		}
	}

	return sum
}

func SolveDay2B(input []string) int {
	games := ParseInput(input)
	sum := 0
	for _, currGame := range games {
		minRed, minGreen, minBlue := 0, 0, 0

		for _, set := range currGame.sets {
			minRed = int(math.Max(float64(minRed), float64(set.red)))
			minBlue = int(math.Max(float64(minBlue), float64(set.blue)))
			minGreen = int(math.Max(float64(minGreen), float64(set.green)))
		}

		sum += minRed * minBlue * minGreen
	}

	return sum
}
