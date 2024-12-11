package day11

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"strconv"
	"strings"
)

type Stones map[int]int

func split(i int, size int) (int, int) {
	a := i
	exp := 1
	for i := 0; i < size/2; i++ {
		a /= 10
		exp *= 10
	}
	b := i - a*exp
	return a, b
}

func size(i int) int {
	res := 0
	for i > 0 {
		i /= 10
		res += 1
	}
	return res
}

func simulate(stones Stones) Stones {
	res := Stones{}
	for stone, count := range stones {
		if stone == 0 {
			res[1] += count
			continue
		}
		sz := size(stone)
		if sz%2 == 0 {
			a, b := split(stone, sz)
			res[a] += count
			res[b] += count
			continue
		}

		res[stone*2024] += count
	}
	return res
}

func px(stones Stones, iter int) int {
	for i := 0; i < iter; i++ {
		stones = simulate(stones)
	}
	res := 0
	for _, count := range stones {
		res += count
	}
	return res
}

func p1(stones Stones) int {
	return px(stones, 25)
}

func p2(stones Stones) int {
	return px(stones, 75)
}

func parseStones(lines []string) (Stones, error) {
	res := Stones{}
	if len(lines) != 1 {
		return res, fmt.Errorf("wrong number of lines")
	}

	stonesStr := strings.Fields(lines[0])
	for _, stoneStr := range stonesStr {
		stone, err := strconv.Atoi(stoneStr)
		if err != nil {
			return res, err
		}
		res[stone] += 1
	}
	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[Stones]("day11/input.txt", parseStones)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
