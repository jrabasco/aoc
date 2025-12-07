package day7

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
)

func isTachyon(r rune) bool {
	return r == 'S' || r == '|'
}

func isSplitter(r rune) bool {
	return r == '^'
}

/*
func p1(lines [][]rune) int {
	res := 0
	for i := range lines {
		if i == len(lines)-1 {
			break
		}
		for j := range lines[i] {
			if !isTachyon(lines[i][j]) {
				continue
			}

			if !isSplitter(lines[i+1][j]) {
				lines[i+1][j] = '|'
				continue
			}

			res++

			if j > 0 && !isSplitter(lines[i+1][j-1]) {
				lines[i+1][j-1] = '|'
			}

			if j < len(lines[i+1])-1 && !isSplitter(lines[i+1][j+1]) {
				lines[i+1][j+1] = '|'
			}
		}
	}
	return res
}
*/

func solve(lines [][]rune) (int, int) {
	history := [][]int{}
	for i := range lines {
		history = append(history, []int{})
		for j := range lines[i] {
			if isTachyon(lines[i][j]) {
				history[i] = append(history[i], 1)
			} else {
				history[i] = append(history[i], 0)
			}
		}
	}
	res1 := 0
	for i := range lines {
		if i == len(lines)-1 {
			break
		}
		for j := range lines[i] {
			if !isTachyon(lines[i][j]) {
				continue
			}
			count := history[i][j]

			if !isSplitter(lines[i+1][j]) {
				lines[i+1][j] = '|'
				history[i+1][j] += count
				continue
			}
			res1++

			if j > 0 && !isSplitter(lines[i+1][j-1]) {
				lines[i+1][j-1] = '|'
				history[i+1][j-1] += count
			}

			if j < len(lines[i+1])-1 && !isSplitter(lines[i+1][j+1]) {
				lines[i+1][j+1] = '|'
				history[i+1][j+1] += count
			}
		}
	}
	res2 := 0
	for _, nb := range history[len(history)-1] {
		res2 += nb
	}
	return res1, res2
}

func Solution() int {
	parsed, err := parse.GetLinesAsRunes("day7/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	res1, res2 := solve(parsed)
	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
