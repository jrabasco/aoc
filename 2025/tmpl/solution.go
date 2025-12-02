package dayX

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
)

func p1(lines []string) int {
	return 0
}

func p2(lines []string) int {
	return 0
}

func Solution() int {
	parsed, err := parse.GetLines("dayX/input_test.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
