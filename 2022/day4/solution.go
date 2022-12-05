package day4

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"strconv"
	"strings"
)

type ElfPair struct {
	e1 utils.Range
	e2 utils.Range
}

var ErrorPair = ElfPair{utils.EmptyRange, utils.EmptyRange}

func lineAsElfPair(line string) (ElfPair, error) {
	elvesStr := strings.Split(line, ",")
	if len(elvesStr) != 2 {
		return ErrorPair, fmt.Errorf("invalid input line %s", line)
	}

	r1, err := parseRange(elvesStr[0])

	if err != nil {
		return ErrorPair, err
	}

	r2, err := parseRange(elvesStr[1])

	if err != nil {
		return ErrorPair, err
	}

	return ElfPair{r1, r2}, nil
}

func parseRange(s string) (utils.Range, error) {
	rStr := strings.Split(s, "-")
	if len(rStr) != 2 {
		return utils.EmptyRange, fmt.Errorf("invalid input %s", s)
	}
	start, err := strconv.ParseInt(rStr[0], 10, 64)
	if err != nil {
		return utils.EmptyRange, err
	}

	end, err := strconv.ParseInt(rStr[1], 10, 64)
	if err != nil {
		return utils.EmptyRange, err
	}

	return utils.NewRange(int(start), int(end)), nil
}

func Solution(part string) int {
	pairs, err := parse.GetLinesAs[ElfPair]("day4/input.txt", lineAsElfPair)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	part1 := 0
	part2 := 0
	for _, pair := range pairs {
		if pair.e1.FullyContains(pair.e2) || pair.e2.FullyContains(pair.e1) {
			part1 += 1
		}

		if pair.e1.Overlaps(pair.e2) {
			part2 += 1
		}
	}

	if part != "2" {
		fmt.Printf("Part 1: %d\n", part1)
	}

	if part != "1" {
		fmt.Printf("Part 2: %d\n", part2)
	}

	return 0
}
