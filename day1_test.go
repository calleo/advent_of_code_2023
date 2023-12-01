package aoc2023

import (
	"os"
	"strings"
	"testing"
)

func TestSolveDay1(t *testing.T) {
	expected := 142
	document := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	actual := SolveDay1A(document)
	if actual != expected {
		t.Fatalf("Actual %d, does not equal expected %d", actual, expected)
	}
}

func TestSolveDay1AFull(t *testing.T) {
	expected := 54940
	content, _ := os.ReadFile("input/day1.txt")
	lines := strings.Split(string(content), "\n")
	actual := SolveDay1A(lines)
	if actual != expected {
		t.Fatalf("Actual %d, does not equal expected %d", actual, expected)
	}
}

func TestSolveDay1B(t *testing.T) {
	expected := 281
	document := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	actual := SolveDay1B(document)
	if actual != expected {
		t.Fatalf("Actual %d, does not equal expected %d", actual, expected)
	}
}

func TestSolveDay1BFull(t *testing.T) {
	expected := 54194
	content, _ := os.ReadFile("input/day1.txt")
	lines := strings.Split(string(content), "\n")
	actual := SolveDay1B(lines)
	if actual == expected {
		t.Fatalf("Actual %d, does not equal expected %d", actual, expected)
	}
}

func TestReplaceNumericDigits(t *testing.T) {
	expected := "8wo3"
	actual := ReplaceNumericDigits("eightwothree")
	if actual != expected {
		t.Fatalf("Actual %s, does not equal expected %s", actual, expected)
	}

	expected = "219"
	actual = ReplaceNumericDigits("two1nine")
	if actual != expected {
		t.Fatalf("Actual %s, does not equal expected %s", actual, expected)
	}
}
