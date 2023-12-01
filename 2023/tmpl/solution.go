package dayX

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
)

func solvePart(part string) int {
	parsed, err := parse.GetLines("dayX/input_test.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	fmt.Printf("Part %s: %v\n", part, parsed)
	return 0
}

func Solution(part string) int {
	if part != "1" && part != "2" {
		p1 := solvePart("1")
		if p1 != 0 {
			return p1
		}
		return solvePart("2")
	} else {
		return solvePart(part)
	}
}
