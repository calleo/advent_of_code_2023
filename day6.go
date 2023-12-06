package aoc2023

import (
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func ParseRaces(input []string) []Race {
	var races []Race

	times := strings.Split(input[0], ":")[1]
	times = strings.Trim(times, " ")

	for _, time := range strings.Split(times, " ") {
		timeInt, err := strconv.Atoi(time)
		if err == nil {
			races = append(races, Race{time: timeInt})
		}
	}

	distances := strings.Split(input[1], ":")[1]
	distances = strings.Trim(distances, " ")
	index := 0

	for _, distance := range strings.Split(distances, " ") {
		distanceInt, err := strconv.Atoi(distance)
		if err == nil {
			races[index].distance = distanceInt
			index++
		}
	}

	return races
}

func ParseRaceB(input []string) Race {
	var race = Race{}

	time := strings.Split(input[0], ":")[1]
	time = strings.Replace(time, " ", "", -1)
	race.time, _ = strconv.Atoi(time)

	distance := strings.Split(input[1], ":")[1]
	distance = strings.Replace(distance, " ", "", -1)
	race.distance, _ = strconv.Atoi(distance)

	return race
}

func GetWinningStrategies(race Race) []int {
	var winningStrategies []int

	for holdTime := 1; holdTime < race.time; holdTime++ {
		travelTime := race.time - holdTime
		if travelTime*holdTime > race.distance {
			winningStrategies = append(winningStrategies, holdTime)
		}
	}

	return winningStrategies
}

func SolveDay6A(input []string) int {
	races := ParseRaces(input)
	product := 1

	for _, race := range races {
		s := GetWinningStrategies(race)
		product = product * len(s)
	}

	return product
}

func SolveDay6B(input []string) int {
	race := ParseRaceB(input)
	return len(GetWinningStrategies(race))
}
