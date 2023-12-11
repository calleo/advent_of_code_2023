package aoc2023

import (
	"strings"
)

type Path struct {
	path   []Point
	failed bool
}

func GetStart(m Matrix) Point {
	for _, line := range m.data {
		for _, p := range line {
			if p.value == "S" {
				return p
			}
		}
	}
	panic("No start found!")
}

func NotSeen(p Point, seen []Point) bool {
	for _, seenPoint := range seen {
		if p == seenPoint {
			return false
		}
	}
	return true
}

func ReachedStart(m Matrix, path Path) Point {
	var next Point
	var err error
	last := path.path[len(path.path)-1]

	if strings.Contains("-J7", last.value) {
		if next, err = m.westOf(last); err == nil && strings.Contains("S", next.value) {
			return next
		}
	}

	if strings.Contains("-LF", last.value) {
		if next, err = m.eastOf(last); err == nil && strings.Contains("S", next.value) {
			return next
		}
	}

	if strings.Contains("|LJ", last.value) {
		if next, err = m.northOf(last); err == nil && strings.Contains("S", next.value) {
			return next
		}
	}

	if strings.Contains("|7F", last.value) {
		if next, err = m.southOf(last); err == nil && strings.Contains("S", next.value) {
			return next
		}
	}

	return next
}

func GetNext(m Matrix, path Path) []Point {
	var nextAll []Point
	var next Point
	last := path.path[len(path.path)-1]
	var err error

	if last.x == 1 && last.y == 3 {
		foo := 1
		foo++
	}

	if strings.Contains("-J7S", last.value) {
		if next, err = m.westOf(last); err == nil && strings.Contains("L-F", next.value) && NotSeen(next, path.path) {
			nextAll = append(nextAll, next)
		}
	}

	if strings.Contains("-LFS", last.value) {
		if next, err = m.eastOf(last); err == nil && strings.Contains("J-7", next.value) && NotSeen(next, path.path) {
			nextAll = append(nextAll, next)
		}
	}

	if strings.Contains("|LJS", last.value) {
		if next, err = m.northOf(last); err == nil && strings.Contains("|7F", next.value) && NotSeen(next, path.path) {
			nextAll = append(nextAll, next)
		}
	}

	if strings.Contains("|7FS", last.value) {
		if next, err = m.southOf(last); err == nil && strings.Contains("|LJ", next.value) && NotSeen(next, path.path) {
			nextAll = append(nextAll, next)
		}
	}

	if nextAll == nil {
		foo := 1
		foo++
	}

	return nextAll
}

func Traverse(matrix Matrix, path Path) Path {
	nextAll := GetNext(matrix, path)
	var successfulPath Path

	if len(nextAll) == 0 {
		if start := ReachedStart(matrix, path); start.value == "" {
			path.failed = true
		} else {
			path.path = append(path.path, start)
		}
		return path
	}

	for _, next := range nextAll {
		var newPath []Point
		newPath = append(newPath, path.path...)
		if len(newPath) == 0 {
			panic("getting nowhere")
		}
		if p := Traverse(matrix, Path{path: append(newPath, next)}); !p.failed {
			successfulPath = p
		}
	}
	return successfulPath
}

func SolveDay10A(input []string) int {
	var path Path
	matrix := InputAsMatrix(input)
	path.path = append(path.path, GetStart(matrix))
	path = Traverse(matrix, path)
	return (len(path.path) - 1) / 2
}

func IsOnPath(p Point, path []Point) bool {
	for _, pathPoint := range path {
		if pathPoint == p {
			return true
		}
	}
	return false
}

func IsEnclosed(p Point, path []Point, m Matrix) bool {
	crossedLines := 0

	if IsOnPath(p, path) {
		return false
	}

	for x := p.x + 1; x < len(m.data); x++ {
		currPoint, _ := m.valueAt(x, p.y)
		// This over-simplification doesn't work since you can have a loop without pipes...
		if currPoint.value == "|" && IsOnPath(currPoint, path) {
			crossedLines++
		}
	}

	return crossedLines != 0 && (crossedLines%2 != 0)
}

func SolveDay10B(input []string) int {
	var path Path
	matrix := InputAsMatrix(input)
	path.path = append(path.path, GetStart(matrix))
	path = Traverse(matrix, path)
	enclosedPoints := 0

	for _, line := range matrix.data {
		for _, point := range line {
			if IsEnclosed(point, path.path, matrix) {
				enclosedPoints++
			}
		}
	}

	return enclosedPoints
}
