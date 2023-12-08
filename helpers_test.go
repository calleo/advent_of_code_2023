package aoc2023

import "testing"

func TestGCD(t *testing.T) {
	actual := GCD(15, 20)
	AssertEquals(actual, 5, t)
}

func TestLCM(t *testing.T) {
	actual := LCM(2, 7, 3, 9, 4)
	AssertEquals(actual, 252, t)
}
