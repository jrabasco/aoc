package day1

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
	"strconv"
)

func p1(nbs []int) int {
	cur := 50
	zeros := 0
	for _, nb := range nbs {
		cur = (cur + nb) % 100
		if cur < 0 {
			cur += 100
		}

		if cur == 0 {
			zeros++
		}
	}
	return zeros
}

func p2(nbs []int) int {
	cur := 50
	zeros := 0
	for _, nb := range nbs {
		if utils.Abs(nb) >= 100 {
			zeros += utils.Abs(nb / 100)
			nb -= (nb / 100) * 100
		}

		if (cur != 0 && cur+nb < 0) || cur+nb > 100 {
			zeros++
		}

		cur = (cur + nb) % 100
		if cur < 0 {
			cur += 100
		}

		if cur == 0 {
			zeros++
		}
	}
	return zeros
}

func parseLine(line string) (int, error) {
	if len(line) < 2 {
		return 0, fmt.Errorf("invalid entry: %s", line)
	}

	sgn := 1
	if line[0] == 'L' {
		sgn = -1
	}

	nb, err := strconv.Atoi(line[1:])

	if err != nil {
		return 0, err
	}

	return sgn * nb, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[int]("day1/input.txt", parseLine)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
