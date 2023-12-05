package day2

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id       int
	cubeSets []CubeSet
}

func parseCubeSet(cubeStr string) (CubeSet, error) {
	cubeParts := strings.Split(cubeStr, ", ")
	res := CubeSet{0, 0, 0}
	for _, part := range cubeParts {
		pparts := strings.Split(part, " ")
		if len(pparts) != 2 {
			return res, fmt.Errorf("malformed part: %s", part)
		}
		nb, err := strconv.Atoi(pparts[0])
		if err != nil {
			return res, err
		}

		if pparts[1] == "red" {
			res.red = nb
		} else if pparts[1] == "green" {
			res.green = nb
		} else if pparts[1] == "blue" {
			res.blue = nb
		} else {
			return res, fmt.Errorf("malformed part: %s", part)
		}
	}
	return res, nil
}

func parseGame(gameStr string) (Game, error) {
	res := Game{0, []CubeSet{}}

	gameNSets := strings.Split(gameStr, ": ")
	if len(gameNSets) != 2 {
		return res, fmt.Errorf("malformed game: %s", gameStr)
	}

	gameNID := strings.Split(gameNSets[0], " ")
	if len(gameNID) != 2 {
		return res, fmt.Errorf("malformed game: %s", gameStr)
	}
	id, err := strconv.Atoi(gameNID[1])

	if err != nil {
		return res, err
	}

	res.id = id

	setStrs := strings.Split(gameNSets[1], "; ")
	for _, setStr := range setStrs {
		set, err := parseCubeSet(setStr)
		if err != nil {
			return res, err
		}
		res.cubeSets = append(res.cubeSets, set)
	}

	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[Game]("day2/input.txt", parseGame)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	res := 0
	for _, game := range parsed {
		possible := true
		for _, set := range game.cubeSets {
			if set.red > 12 || set.green > 13 || set.blue > 14 {
				possible = false
			}
		}
		if possible {
			res += game.id
		}
	}
	fmt.Printf("Part 1: %d\n", res)
	res = 0
	for _, game := range parsed {
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, set := range game.cubeSets {
			if set.red > minRed {
				minRed = set.red
			}

			if set.green > minGreen {
				minGreen = set.green
			}

			if set.blue > minBlue {
				minBlue = set.blue
			}
		}
		res += minRed * minGreen * minBlue
	}
	fmt.Printf("Part 2: %d\n", res)
	return 0
}
