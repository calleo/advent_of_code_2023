package aoc2023

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

type Matrix struct {
	data [][]Point
}

type Point struct {
	x     int
	y     int
	value string
}

func (m *Matrix) valueAt(x, y int) (Point, error) {
	if y >= len(m.data) || y < 0 {
		return Point{}, errors.New(fmt.Sprintf("y %d is out of bounds", y))
	} else if x >= len(m.data[y]) || x < 0 {
		return Point{}, errors.New(fmt.Sprintf("x %d is out of bounds", x))
	}
	return m.data[y][x], nil
}

func (m *Matrix) southOf(p Point) (Point, error) {
	return m.valueAt(p.x, p.y+1)
}

func (m *Matrix) northOf(p Point) (Point, error) {
	return m.valueAt(p.x, p.y-1)
}

func (m *Matrix) westOf(p Point) (Point, error) {
	return m.valueAt(p.x-1, p.y)
}

func (m *Matrix) eastOf(p Point) (Point, error) {
	return m.valueAt(p.x+1, p.y)
}

func (m *Matrix) setAt(x, y int, value string) {
	m.data[y][x].value = value
}

func (m *Matrix) neighboursOf(p Point) []Point {
	var values []Point
	indices := [][]int{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}

	for _, index := range indices {
		val, err := m.valueAt(index[0], index[1])
		if err == nil {
			values = append(values, val)
		}
	}

	return values
}

func ReadInputFile(name string) []string {
	filename := fmt.Sprintf("input/%s", name)
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(fmt.Sprintf("Error opening file %s. Error: %s", name, err))
	}

	return strings.Split(string(content), "\n")
}

func InputAsMatrix(content []string) Matrix {
	var data [][]Point

	for y, line := range content {
		var points []Point
		splitLine := strings.Split(line, "")
		for x, value := range splitLine {
			points = append(points, Point{x: x, y: y, value: value})
		}
		data = append(data, points)
	}

	return Matrix{data: data}
}

func sumSlice(numbers []int) int {
	sum := 0
	for _, nr := range numbers {
		sum += nr
	}
	return sum
}

func AssertEquals[T comparable](actual T, expected T, t *testing.T) {
	if actual != expected {
		t.Fatalf("Actual %+v, does not equal expected %+v", actual, expected)
	}
}

func AssertNotEquals(actual int, expected int, t *testing.T) {
	if actual == expected {
		t.Fatalf("Actual %d, does equal expected %d", actual, expected)
	}
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}

	tmp := a
	a = b
	b = tmp % a

	return GCD(a, b)
}

func LCM(nums ...int) int {
	if len(nums) == 2 {
		return nums[0] * nums[1] / GCD(nums[0], nums[1])
	}
	return LCM(nums[0], LCM(nums[1:]...))
}
