package day1

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
)

func findPos(top []int64, val int64) int {
	for i, t := range top {
		if val > t {
			return i
		}
	}
	return len(top)
}

func insertInTop(top *[]int64, val int64) {
	max := len(*top)
	pos := findPos(*top, val)
	if pos == max {
		return
	}

	for i := pos + 1; i < max; i++ {
		(*top)[i] = (*top)[i-1]
	}

	(*top)[pos] = val
}

func sum(top []int64) int64 {
	var res int64
	for _, t := range top {
		res += t
	}
	return res
}

func Solution(part string) int {
	lines, err := parse.GetLines("day1/input.txt")

	if err != nil {
		fmt.Printf("Could not read lines: %s\n", err)
		return 1
	}

	var curTot int64
	var topThree = []int64{0, 0, 0}

	for _, line := range lines {
		if len(line) == 0 {
			insertInTop(&topThree, curTot)
			curTot = 0
			continue
		}

		cal, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}

		curTot += cal
	}

	insertInTop(&topThree, curTot)

	if part != "2" {
		fmt.Printf("Part 1: %d\n", topThree[0])
	}

	if part != "1" {
		fmt.Printf("Part 2: %d\n", sum(topThree))
	}

	return 0
}
