package day1

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"sort"
	"strconv"
)

func p1(intPairs [][]int) int {
	l1 := []int{}
	l2 := []int{}
	for _, pair := range intPairs {
		l1 = append(l1, pair[0])
		l2 = append(l2, pair[1])
	}

	sort.Slice(l1, func(i, j int) bool {
		return l1[i] < l1[j]
	})
	sort.Slice(l2, func(i, j int) bool {
		return l2[i] < l2[j]
	})

	res := 0
	for i := 0; i < len(l1); i++ {
		res += utils.AbsDiff(l1[i], l2[i])
	}
	return res
}

func p2(intPairs [][]int) int {
	l1 := []int{}
	l2 := map[int]int{}
	for _, pair := range intPairs {
		l1 = append(l1, pair[0])
		l2[pair[1]] += 1
	}

	res := 0
	for _, i := range l1 {
		cnt, found := l2[i]
		if found {
			res += cnt * i
		}
	}
	return res
}

func Solution() int {
	parsedInts, err := parse.GetLinesAsFields[int]("day1/input.txt", strconv.Atoi)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	fmt.Printf("Part 1: %d\n", p1(parsedInts))
	fmt.Printf("Part 2: %d\n", p2(parsedInts))
	return 0
}
