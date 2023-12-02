package aoc2023

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func ReadInputFile(name string) []string {
	filename := fmt.Sprintf("input/%s", name)
	content, _ := os.ReadFile(filename)
	return strings.Split(string(content), "\n")
}

func AssertEquals(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Fatalf("Actual %d, does not equal expected %d", actual, expected)
	}
}

func AssertNotEquals(actual int, expected int, t *testing.T) {
	if actual == expected {
		t.Fatalf("Actual %d, does equal expected %d", actual, expected)
	}
}
