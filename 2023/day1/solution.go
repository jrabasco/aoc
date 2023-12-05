package day1

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Solution() int {
	parsed, err := parse.GetLines("day1/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	res := 0

	res, err = part1(parsed)
	if err != nil {
		fmt.Printf("Part 1 failed with %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", res)
	res, err = part2(parsed)
	if err != nil {
		fmt.Printf("Part 2 failed with %v\n", err)
		return 1
	}
	fmt.Printf("Part 2: %d\n", res)
	return 0
}

type tpl struct {
	val string
	idx int
}

func part2(lines []string) (int, error) {
	var dict map[string]string = map[string]string{
		"one":   "1",
		"1":     "1",
		"two":   "2",
		"2":     "2",
		"three": "3",
		"3":     "3",
		"four":  "4",
		"4":     "4",
		"five":  "5",
		"5":     "5",
		"six":   "6",
		"6":     "6",
		"seven": "7",
		"7":     "7",
		"eight": "8",
		"8":     "8",
		"nine":  "9",
		"9":     "9",
	}

	tot := 0
	for _, line := range lines {
		candidates := []tpl{}
		for word, value := range dict {
			fidx := strings.Index(line, word)
			if fidx != -1 {
				candidates = append(candidates, tpl{value, fidx})
			}
			lidx := strings.LastIndex(line, word)
			if lidx != -1 {
				candidates = append(candidates, tpl{value, lidx})
			}
		}
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].idx < candidates[j].idx
		})
		valStr := candidates[0].val + candidates[len(candidates)-1].val
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, err
		}
		tot += val
	}
	return tot, nil
}

func part1(lines []string) (int, error) {
	tot := 0
	for _, line := range lines {
		valStr := ""
		first := true
		last := 'x'
		for _, char := range line {
			if unicode.IsDigit(char) {
				last = char
				if first {
					valStr += string(last)
					first = false
				}
			}
		}
		valStr += string(last)
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, err
		}
		tot += val
	}
	return tot, nil
}
