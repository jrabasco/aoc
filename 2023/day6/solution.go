package day6

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"math"
	"strconv"
	"strings"
)

type Race struct {
	time     float64
	distance float64
}

func NewRace(time int, distance int) Race {
	return Race{float64(time), float64(distance)}
}

// If we try to plot the distance a boat will reach as a function of the amount of
// time the button was pressed, we get the following equation:
//
// distance = x*(time - x) = x*time - x^2
//
// where x is the amount of time you pressed on the button
//
// This function is a parabola that has a maximum on x=time/2. We are looking to
// find how many integer values give a result higher than the input distance.
// For that we need to first solve for equality:
//
// distance = x*time - x^2
//
// which we can re-arrange as
//
// x^2 - time * x + distance = 0
//
// This is a quadratic equation which we can then solve:
// delta = time^2 - 4 * distance
//
// if delta > 0 then we know our solutions are all the integers between
//
// (time - sqrt(delta))/2 and (time + sqrt(delta))/2
//
// if delta == 0 then there is no solution as it means there's only one point
// where it's equal so there is no higher point
//
// if delta < 0 something weird is happening to these boats.
func (r Race) betterRacesCount() int {
	delta := r.time*r.time - 4*r.distance

	if delta == 0 {
		return 0
	}

	if delta < 0 {
		panic("wtf?")
	}

	under := (r.time - math.Sqrt(delta)) / 2
	over := (r.time + math.Sqrt(delta)) / 2

	// Find the next integer above the "under" limit as that is the lowest
	// amount of time we can press to beat it.
	// examples:
	// 10 -> floor(11) -> 11
	// 5.5 -> floor(6.5) -> 6
	nextUnderInt := math.Floor(under + 1)

	// Find the previous integer under the "over" limit as that is the highest
	// amount of time we can press to beat it.
	// examples:
	// 20 -> ceil(19) -> 19
	// 7.5 -> ceil(6.5) -> 7
	prevOverInt := math.Ceil(over - 1)

	return int(prevOverInt-nextUnderInt) + 1
}

func parseRaces(lines []string) ([]Race, error) {
	res := []Race{}
	if len(lines) != 2 {
		return res, fmt.Errorf("malformed input, wrong number of lines")
	}

	timeParts := strings.Fields(lines[0])
	distanceParts := strings.Fields(lines[1])

	if len(timeParts) != len(distanceParts) {
		return res, fmt.Errorf("malformed input, time and distance lengths differ")
	}

	if len(timeParts) < 2 {
		return res, fmt.Errorf("malformed input, not enough elements")
	}

	timeParts = timeParts[1:]
	distanceParts = distanceParts[1:]

	for i, timeStr := range timeParts {
		distanceStr := distanceParts[i]

		time, err := strconv.Atoi(timeStr)
		if err != nil {
			return res, err
		}

		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			return res, err
		}

		race := NewRace(time, distance)
		res = append(res, race)
	}

	return res, nil
}

func parseRace(lines []string) (Race, error) {
	if len(lines) != 2 {
		return Race{}, fmt.Errorf("malformed input, wrong number of lines")
	}

	timeParts := strings.Fields(lines[0])
	distanceParts := strings.Fields(lines[1])

	if len(timeParts) < 2 || len(distanceParts) < 2 {
		return Race{}, fmt.Errorf("malformed input, not enough elements")
	}

	timeStr := strings.Join(timeParts[1:], "")
	distanceStr := strings.Join(distanceParts[1:], "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return Race{}, err
	}

	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return Race{}, err
	}
	return NewRace(time, distance), nil
}

func Solution() int {
	parsed, err := parse.GetLines("day6/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	// Part 1
	races, err := parseRaces(parsed)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	res := 1
	for _, race := range races {
		res *= race.betterRacesCount()
	}

	fmt.Printf("Part 1: %d\n", res)

	// Part 2
	race, err := parseRace(parsed)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	fmt.Printf("Part 2: %d\n", race.betterRacesCount())
	return 0
}
