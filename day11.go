package aoc2023

import (
	"fmt"
	"github.com/mowshon/iterium"
	"math"
	"strings"
)

func Empty(line string) bool {
	for _, ch := range strings.Split(line, "") {
		if ch != "." {
			return false
		}
	}
	return true
}

func EmptyRowsAndCols(m Matrix) ([]int, []int) {
	var emptyRows []int
	var emptyCols []int

	// Count empty rows
	for y, line := range m.data {
		fullLine := ""
		for _, p := range line {
			fullLine += p.value
		}
		if Empty(fullLine) {
			emptyRows = append(emptyRows, y)
		}
	}

	// Count empty columns
	for x := 0; x < len(m.data[0]); x++ {
		fullLine := ""
		for _, line := range m.data {
			fullLine += line[x].value
		}
		if Empty(fullLine) {
			emptyCols = append(emptyCols, x)
		}
	}

	return emptyRows, emptyCols
}

func SolveDay11AB(input []string, expansion int) int {
	matrix := InputAsMatrix(input)
	count := 1
	var galaxyIDs []int
	var galaxies = make(map[int]Point)
	emptyRows, emptyCols := EmptyRowsAndCols(matrix)
	expansion = expansion - 1 // Each empty row is counted once normally, so subtract that one

	for _, line := range matrix.data {
		for _, point := range line {
			if point.value == "#" {
				matrix.setAt(point.x, point.y, fmt.Sprintf("%d", count))
				galaxyIDs = append(galaxyIDs, count)
				galaxies[count], _ = matrix.valueAt(point.x, point.y)
				count++
			}
		}
	}

	CountEmpty := func(from Point, to Point) int {
		minY := int(math.Min(float64(from.y), float64(to.y)))
		maxY := int(math.Max(float64(from.y), float64(to.y)))
		minX := int(math.Min(float64(from.x), float64(to.x)))
		maxX := int(math.Max(float64(from.x), float64(to.x)))
		empty := 0

		for _, row := range emptyRows {
			if row > minY && row < maxY {
				empty++
			}
		}

		for _, col := range emptyCols {
			if col > minX && col < maxX {
				empty++
			}
		}

		return empty
	}

	combinations := iterium.Combinations(galaxyIDs, 2)
	toSlice, _ := combinations.Slice()
	totalDistance := 0

	for _, combination := range toSlice {
		from := galaxies[combination[0]]
		to := galaxies[combination[1]]
		empty := CountEmpty(from, to)
		totalDistance += expansion * empty
		totalDistance += int(math.Abs(float64(from.x - to.x)))
		totalDistance += int(math.Abs(float64(from.y - to.y)))
	}

	return totalDistance
}
