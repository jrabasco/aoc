package day9

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"strconv"
)

func findPredictions(history []int) (int, int) {
	lh := len(history)
	if lh == 0 {
		return 0, 0
	}

	if lh == 1 {
		return history[0], history[lh-1]
	}

	nh := []int{}
	allzero := true
	for i := 0; i < lh-1; i++ {
		diff := history[i+1] - history[i]
		if diff != 0 {
			allzero = false
		}
		nh = append(nh, diff)
	}
	if allzero {
		return history[0], history[lh-1]
	}

	beg, end := findPredictions(nh)
	return history[0] - beg, history[lh-1] + end
}

func Solution() int {
	histories, err := parse.GetLinesAsFields[int]("day9/input.txt", strconv.Atoi)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	res1 := 0
	res2 := 0
	for _, history := range histories {
		beg, end := findPredictions(history)
		res1 += end
		res2 += beg
	}
	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
