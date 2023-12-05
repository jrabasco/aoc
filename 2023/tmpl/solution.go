package dayX

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
)

func Solution() int {
	parsed, err := parse.GetLines("dayX/input_test.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	fmt.Printf("Part %s: %v\n", part, parsed)
	return 0
}
