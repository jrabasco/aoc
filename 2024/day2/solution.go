package day2

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"strconv"
)

func p1(reports [][]int) int {
	safeNb := 0
	for _, report := range reports {
		safe := isSafe(report)
		if safe {
			safeNb++
		}
	}
	return safeNb
}

func isSafe(report []int) bool {
	first := true
	curDir := 0
	last := report[0]
	for i := 1; i < len(report); i++ {
		cur := report[i]
		step := cur - last
		absStep := utils.Abs(step)
		nDir := 0
		if absStep != 0 {
			nDir = step / absStep
		}
		if first {
			first = false
			curDir = nDir
		}
		if absStep == 0 || absStep > 3 || curDir != nDir {
			return false
		}
		last = cur
		curDir = nDir
	}
	return true
}
func p2(reports [][]int) int {
	safeNb := 0
	for _, report := range reports {
		safe := isSafe(report)
		if safe {
			safeNb++
			continue
		}

		for i := 0; i < len(report); i++ {
			skip := utils.RemoveIndex[int](report, i)
			safe := isSafe(skip)
			if safe {
				safeNb++
				break
			}
		}
	}
	return safeNb
}

func Solution() int {
	reports, err := parse.GetLinesAsFields[int]("day2/input.txt", strconv.Atoi)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(reports))
	fmt.Printf("Part 2: %d\n", p2(reports))
	return 0
}
