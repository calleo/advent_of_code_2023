package aoc2023

import (
	"math"
	"strings"
)

func SolveDay14A(input []string) int {
	matrix := InputAsMatrix(input)
	totalNrOfRows := len(matrix.data)
	totalLoad := 0

	for _, line := range matrix.data {
		for _, cell := range line {
			if cell.value == "O" {
				rocks := 0
				row := cell.y - 1
				for row >= 0 {
					p, _ := matrix.valueAt(cell.x, row)
					if p.value == "O" {
						rocks++
					} else if p.value == "#" {
						row++
						break
					}
					row--
				}
				rockLoad := totalNrOfRows - int(math.Max(float64(row), float64(0))) - rocks
				totalLoad += rockLoad
			}
		}
	}

	return totalLoad
}

func TiltNorth(m *Matrix) {
	for _, line := range m.data {
		for _, cell := range line {
			if cell.value == "O" {
				newY := cell.y
				for y := cell.y - 1; y >= 0; y-- {
					value, _ := m.valueAt(cell.x, y)
					if value.value == "." {
						newY = y
					} else if strings.Contains("#O", value.value) {
						break
					}
				}
				if newY != cell.y {
					// swap
					m.setAt(cell.x, newY, "O")
					m.setAt(cell.x, cell.y, ".")
				}
			}
		}
	}
}

func TiltSouth(m *Matrix) {
	for y := len(m.data) - 1; y >= 0; y-- {
		line := m.data[y]
		for _, cell := range line {
			if cell.value == "O" {
				newY := cell.y
				for cellY := cell.y + 1; cellY < len(m.data); cellY++ {
					value, _ := m.valueAt(cell.x, cellY)
					if value.value == "." {
						newY = cellY
					} else if strings.Contains("#O", value.value) {
						break
					}
				}
				if newY != cell.y {
					// swap
					m.setAt(cell.x, newY, "O")
					m.setAt(cell.x, cell.y, ".")
				}
			}
		}
	}
}

func TiltWest(m *Matrix) {
	for x := 0; x < len(m.data[0]); x++ {
		for y, _ := range m.data {
			cell, _ := m.valueAt(x, y)
			if cell.value == "O" {
				newX := cell.x
				for cellX := cell.x - 1; cellX >= 0; cellX-- {
					value, _ := m.valueAt(cellX, cell.y)
					if value.value == "." {
						newX = value.x
					} else if strings.Contains("#O", value.value) {
						break
					}
				}
				if newX != cell.x {
					// swap
					m.setAt(newX, cell.y, "O")
					m.setAt(cell.x, cell.y, ".")
				}
			}
		}
	}
}

func TiltEast(m *Matrix) {
	for x := len(m.data[0]) - 1; x >= 0; x-- {
		for y, _ := range m.data {
			cell, _ := m.valueAt(x, y)
			if cell.value == "O" {
				newX := cell.x
				for cellX := cell.x + 1; cellX < len(m.data[0]); cellX++ {
					value, _ := m.valueAt(cellX, cell.y)
					if value.value == "." {
						newX = value.x
					} else if strings.Contains("#O", value.value) {
						break
					}
				}
				if newX != cell.x {
					// swap
					m.setAt(newX, cell.y, "O")
					m.setAt(cell.x, cell.y, ".")
				}
			}
		}
	}
}

func (m *Matrix) Load() int {
	load := 0
	totalNrOfRows := len(m.data)
	for y, line := range m.data {
		for _, cell := range line {
			if cell.value == "O" {
				load += totalNrOfRows - y
			}
		}
	}
	return load
}

func SolveDay14B(input []string) int {
	sequences := 1_000_000_000
	matrix := InputAsMatrix(input)
	fromTo := make(map[string]string)
	matrixLoad := make(map[string]int)
	nextMatrix := ""

	for i := 0; i < sequences; i++ {
		fromChecksum := matrix.Checksum()

		next, exists := fromTo[fromChecksum]
		if exists {
			// If we've seen this start position, we know the next
			// position and there is no need to do the actual tilting
			nextMatrix = next
			i++
			for i < sequences {
				nextMatrix, exists = fromTo[nextMatrix]
				i++
			}
		} else {
			TiltNorth(&matrix)
			TiltWest(&matrix)
			TiltSouth(&matrix)
			TiltEast(&matrix)
			toChecksum := matrix.Checksum()
			fromTo[fromChecksum] = toChecksum
			matrixLoad[toChecksum] = matrix.Load()
			nextMatrix = toChecksum
		}
	}

	return matrixLoad[nextMatrix]
}
