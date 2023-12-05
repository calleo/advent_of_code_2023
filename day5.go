package aoc2023

import (
	"math"
	"strconv"
	"strings"
)

type Almanac struct {
	seeds []int
	//mappers map[string][][]int
	mappers [][][]int
}

func ParseAlmanac(input []string) Almanac {
	almanac := Almanac{}
	//almanac.mappers = make(map[string][][]int)
	seeds := strings.Split(strings.Split(input[0], ":")[1], " ")
	for _, seed := range seeds {
		if seedAsInt, err := strconv.Atoi(seed); err == nil {
			almanac.seeds = append(almanac.seeds, seedAsInt)
		}
	}

	//var mapperName = ""
	var mappings [][]int
	for index, line := range input[2:] {
		if strings.Contains(line, ":") || len(input[2:])-1 == index {
			if len(mappings) > 0 {
				almanac.mappers = append(almanac.mappers, mappings)
				//mapperName = ""
				mappings = nil
			}
			//mapperName = strings.Split(line, " ")[0]
		} else {
			var mappingRow []int
			numbers := strings.Split(line, " ")
			for _, nr := range numbers {
				if nrAsInt, err := strconv.Atoi(nr); err == nil {
					mappingRow = append(mappingRow, nrAsInt)
				}
			}
			if len(mappingRow) > 0 {
				mappings = append(mappings, mappingRow)
			}
		}
	}

	return almanac
}

func GetMapping(seed int, mappings [][]int) int {
	for _, mapping := range mappings {
		input := mapping[1]
		output := mapping[0]
		length := mapping[2]
		if seed >= input && seed < input+length {
			diff := seed - input
			return output + diff
		}
	}
	return seed
}

func SolveDay5A(input []string) int {
	almanac := ParseAlmanac(input)
	nextSeed := 0
	min := math.MaxInt

	for _, seed := range almanac.seeds {
		nextSeed = seed
		for _, mapper := range almanac.mappers {
			nextSeed = GetMapping(nextSeed, mapper)
		}
		min = int(math.Min(float64(min), float64(nextSeed)))
	}

	return min
}

func SolveDay5B(input []string) int {
	almanac := ParseAlmanac(input)
	min := math.MaxInt

	var seedRanges [][]int

	for i := 0; i < len(almanac.seeds); i += 2 {
		seedRanges = append(seedRanges, []int{almanac.seeds[i], almanac.seeds[i+1]})
	}

	for _, seedRange := range seedRanges {
		for seed := seedRange[0]; seed < (seedRange[0] + seedRange[1]); seed++ {
			nextSeed := seed
			for _, mapper := range almanac.mappers {
				nextSeed = GetMapping(nextSeed, mapper)
			}
			min = int(math.Min(float64(min), float64(nextSeed)))
		}
	}

	return min
}
