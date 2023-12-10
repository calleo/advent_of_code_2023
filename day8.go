package aoc2023

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type Position struct {
	name  string
	left  *Position
	right *Position
}

func ParseMap(input []string, partB bool) []*Position {
	var starts []*Position
	re := regexp.MustCompile("^.*A$")
	positions := make(map[string]*Position)

	for _, line := range input[2:] {
		position := &Position{}
		nameAndDirections := strings.Split(line, " = ")
		name := nameAndDirections[0]

		if val, ok := positions[name]; ok {
			position = val
		} else {
			position = &Position{name: name}
			positions[name] = position
		}

		leftAndRight := strings.Split(nameAndDirections[1], ", ")

		left := strings.Replace(leftAndRight[0], "(", "", -1)
		right := strings.Replace(leftAndRight[1], ")", "", -1)

		if val, ok := positions[left]; ok {
			position.left = val
		} else {
			position.left = &Position{name: left}
			positions[left] = position.left
		}

		if val, ok := positions[right]; ok {
			position.right = val
		} else {
			position.right = &Position{name: right}
			positions[right] = position.right
		}

		if name == "AAA" && !partB {
			starts = append(starts, position)
		} else if partB {
			if re.MatchString(name) {
				starts = append(starts, position)
			}
		}
	}
	return starts
}

func SolveDay8A(input []string) int {
	directions := strings.Split(input[0], "")
	currentPosition := ParseMap(input, false)[0]
	steps := 0

	for {
		for _, direction := range directions {
			steps += 1
			switch direction {
			case "L":
				currentPosition = currentPosition.left
			case "R":
				currentPosition = currentPosition.right
			}
			if currentPosition.name == "ZZZ" {
				break
			}
		}
		if currentPosition.name == "ZZZ" {
			break
		}
	}

	return steps
}

func Done(position *Position) bool {
	re := regexp.MustCompile("^.*Z$")
	if !re.MatchString(position.name) {
		return false
	}
	return true
}

func SolveDay8B(input []string, t *testing.T) map[string]int {
	directions := strings.Split(input[0], "")
	positions := ParseMap(input, true)
	var startPositions []string
	positionSteps := make(map[string]int)
	for _, pos := range positions {
		startPositions = append(startPositions, pos.name)
	}

	steps := 0

	for {
		for _, direction := range directions {
			for index := 0; index < len(positions); index++ {
				position := positions[index]

				if position.name == position.left.name && position.name == position.right.name {
					panic(fmt.Sprintf("Endless loop: %s", position.name))
				}

				switch direction {
				case "L":
					positions[index] = position.left
				case "R":
					positions[index] = position.right
				}
				if Done(position) {
					if _, ok := positionSteps[startPositions[index]]; !ok {
						positionSteps[startPositions[index]] = steps
					} else {
						val, _ := positionSteps[startPositions[index]]
						if val != steps {
							t.Log(fmt.Sprintf("Got a new step value for %d, %d", index, steps))
						}
					}
				}
			}
			steps += 1
			if len(positionSteps) == len(positions) {
				break
			}
		}
		if len(positionSteps) == len(positions) {
			break
		}
	}

	return positionSteps
}
