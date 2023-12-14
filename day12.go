package aoc2023

import (
	"github.com/schwarmco/go-cartesian-product"
	"golang.org/x/exp/slices"
	"math"
	"strconv"
	"strings"
)

type SpringSequence struct {
	length            int
	position          int
	possiblePositions []interface{}
}

type ConditionRecord struct {
	springs         []string
	possibleSprings []string
	springSequences []*SpringSequence
}

func ParseConditionRecords(input []string) []*ConditionRecord {
	var readings []*ConditionRecord
	for _, line := range input {
		r := &ConditionRecord{}
		r.springs = strings.Split(strings.Split(line, " ")[0], "")
		sequences := strings.Split(strings.Split(line, " ")[1], ",")

		for _, reading := range r.springs {
			if reading != "." {
				reading = "#"
			}
			r.possibleSprings = append(r.possibleSprings, reading)
		}

		for pos, seq := range sequences {
			seqInt, _ := strconv.Atoi(seq)
			r.springSequences = append(r.springSequences, &SpringSequence{length: seqInt, position: pos})
		}
		readings = append(readings, r)
	}
	return readings
}

func CountPossibleArrangements(record *ConditionRecord) int {
	var allPossibilities [][]interface{}
	arrangements := 0

	for _, seq := range record.springSequences {
		allPossibilities = append(allPossibilities, seq.possiblePositions)
	}

	possibleArrangements := cartesian.Iter(allPossibilities...)

	for arr := range possibleArrangements {

		positionValid := true
		for index, position := range arr {
			if index == 0 {
				// this one is always ok
				continue
			}

			// get the earliest starting position by looking at the spring to the left
			lengthLeft := arr[index-1].(int) + 1 + record.springSequences[index-1].length

			if position.(int) < lengthLeft {
				positionValid = false
				break
			}
		}

		if positionValid {
			springs := make([]string, len(record.springs))
			copy(springs, record.springs)
			for index, springSeq := range arr {
				length := record.springSequences[index].length
				for i := springSeq.(int); i < springSeq.(int)+length; i++ {
					springs[i] = ""
				}
			}
			if slices.Contains(springs, "#") {
				positionValid = false
			}
		}

		if positionValid {
			arrangements++
		}
	}

	return arrangements
}

func IsValidPosition(length int, start int, end int, springs []string) bool {
	area := strings.Join(springs[start:end], "")
	if !strings.Contains(area, ".") && len(area) == length {
		if start > 0 {
			if springs[start-1] == "#" {
				return false
			}
		}
		if end < len(springs) {
			if springs[end] == "#" {
				return false
			}
		}
		return true
	}
	return false
}

func FindTotalArrangements(records []*ConditionRecord) int {
	for _, f := range records {
		accumulatedLength := 0
		totalLength := 0
		for _, seq := range f.springSequences {
			totalLength += seq.length
		}
		for _, seq := range f.springSequences {
			for pos, _ := range f.springs {
				end := int(math.Min(float64(pos+seq.length), float64(len(f.springs))))

				if accumulatedLength > 0 && pos <= accumulatedLength {
					// if knowing that we have at least N springs to the left, we can disregard those positions
					continue
				}

				spaceLeft := totalLength - accumulatedLength - seq.length
				if end-1 > len(f.springs)-spaceLeft {
					// if knowing that we have at least N springs to the right, we can disregard those positions
					continue
				}

				if IsValidPosition(seq.length, pos, end, f.springs) {
					seq.possiblePositions = append(seq.possiblePositions, pos)
				}

			}
			accumulatedLength += seq.length
		}
	}

	totalArrangements := 0
	for _, f1 := range records {
		this := CountPossibleArrangements(f1)
		totalArrangements += this
	}

	return totalArrangements
}

func SolveDay12A(input []string) int {
	records := ParseConditionRecords(input)
	return FindTotalArrangements(records)
}

func SolveDay12B(input []string, expansions int) int {
	var newInput []string

	for _, line := range input {
		springs := strings.Split(line, " ")[0]
		newSprings := springs

		if expansions < 2 {
			newSprings += "?"
		}

		seq := strings.Split(line, " ")[1]
		newSeq := seq
		for i := 0; i < expansions-1; i++ {
			newSprings = newSprings + "?" + springs
			newSeq = newSeq + "," + seq
		}
		newInput = append(newInput, newSprings+" "+newSeq)
	}

	records := ParseConditionRecords(newInput)
	return FindTotalArrangements(records)
}
