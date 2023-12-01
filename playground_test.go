package aoc2023

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	expected := "hello world"
	actual := Solve()

	if actual != expected {
		t.Fatalf("Actual %s, does not equal expected %s", expected, actual)
	}
}

func TestSliceLength(t *testing.T) {
	expected := 7
	array := []int{1, 1, 1, 1, 1, 2, 2}
	actual := SliceLength(array)

	if actual != expected {
		t.Fatalf("Actual %d, does not equal expected %d", expected, actual)
	}
}

func TestSliceAppend(t *testing.T) {
	expected := []int{1, 1, 1, 2, 3}
	array := []int{1, 1}
	actual := SliceAppend(array)

	if !reflect.DeepEqual(actual, expected) {
		_actual, _ := json.Marshal(actual)
		_expected, _ := json.Marshal(expected)
		t.Fatalf("Actual %s, does not equal expected %s", _actual, _expected)
	}
}
