package day7

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
)

func Solution() int {
	parsed, err := parse.GetLines("day7/input_test.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %v\n", parsed)
	return 0
}
