package aoc2023

import (
	"strconv"
	"strings"
)

func containsSymbol(points []Point) bool {
	for _, point := range points {
		if point.value != "." && !strings.Contains("0123456789", point.value) {
			return true
		}
	}
	return false
}

func isPartNumber(number []Point, m Matrix) bool {
	if len(number) == 0 {
		return false
	}
	for _, point := range number {
		neighbours := m.neighboursOf(point)
		if containsSymbol(neighbours) {
			return true
		}
	}
	return false
}

func toInt(numberPos []Point) int {
	var number []string
	for _, point := range numberPos {
		number = append(number, point.value)
	}
	asInt, _ := strconv.Atoi(strings.Join(number, ""))
	return asInt
}

func SolveDay3A(input Matrix) int {
	var partNumbers []int
	for y, line := range input.data {
		var numberPoints []Point
		for x, _ := range line {
			point, _ := input.valueAt(x, y)
			if strings.Contains("0123456789", point.value) {
				numberPoints = append(numberPoints, point)
			} else if isPartNumber(numberPoints, input) {
				asInt := toInt(numberPoints)
				partNumbers = append(partNumbers, asInt)
				numberPoints = nil
			} else {
				numberPoints = nil
			}

			if x == len(line)-1 {
				if isPartNumber(numberPoints, input) {
					asInt := toInt(numberPoints)
					partNumbers = append(partNumbers, asInt)
				}
				numberPoints = nil
			}
		}
	}
	return sumSlice(partNumbers)
}

func SolveDay3B(input Matrix) int {
	gears := make(map[Point][]int)

	for y, line := range input.data {
		var numberPoints []Point

		for x, _ := range line {
			point, _ := input.valueAt(x, y)

			if strings.Contains("0123456789", point.value) {
				numberPoints = append(numberPoints, point)
			}

			if (!strings.Contains("0123456789", point.value) || x == len(line)-1) && len(numberPoints) > 0 {
				uniqueNeighbours := make(map[Point]int)
				for _, p := range numberPoints {
					for _, neighbour := range input.neighboursOf(p) {
						uniqueNeighbours[neighbour] = 1
					}
				}
				for _, p := range numberPoints {
					delete(uniqueNeighbours, p)
				}

				for key, _ := range uniqueNeighbours {
					if key.value == "*" {
						number := toInt(numberPoints)
						if _, ok := gears[key]; ok {
							gears[key] = append(gears[key], number)
						} else {
							gears[key] = []int{number}
						}
						break
					}
				}
				numberPoints = nil
			}

		}
	}
	sum := 0
	for _, value := range gears {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}
	return sum
}
