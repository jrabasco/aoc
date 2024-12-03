package day3

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"strconv"
	"strings"
	"unicode"
)

func nextNumber(liner []rune, idx int) (int, int) {
	cur := liner[idx]
	ns := ""

	for idx < len(liner) {
		cur = liner[idx]
		if !unicode.IsDigit(cur) {
			break
		}
		ns += string(cur)
		idx++
	}

	if len(ns) == 0 {
		return 0, -1
	}

	res, _ := strconv.Atoi(ns)
	return res, idx
}

func nextMul1(line string) (int, int) {
	found := false
	res := 0
	idx := 0
	for idx < len(line) && !found {
		nidx := strings.Index(line[idx:], "mul(")
		if nidx < 0 {
			return 0, -1
		}
		idx += nidx
		if idx+4 >= len(line) {
			return 0, -1
		}
		liner := []rune(line)
		n1, n1idx := nextNumber(liner, idx+4)
		if n1idx >= 0 {
			idx = n1idx
		}

		if idx >= len(line) || idx < 0 || liner[idx] != ',' {
			continue
		}

		n2, n2idx := nextNumber(liner, idx+1)
		if n2idx >= 0 {
			idx = n2idx
		}

		if idx >= len(line) || idx < 0 || liner[idx] != ')' {
			continue
		}

		found = true
		res = n1 * n2
	}

	if found {
		return res, idx
	} else {
		return 0, -1
	}
}

func nextMul2(line string) (int, int) {
	found := false
	res := 0
	idx := 0
	for idx < len(line) && !found {
		mulidx := strings.Index(line[idx:], "mul(")
		if mulidx < 0 {
			return 0, -1
		}
		dontidx := strings.Index(line[idx:], "don't()")

		if dontidx >= 0 && dontidx < mulidx {
			idx += dontidx
			doidx := strings.Index(line[idx:], "do()")
			if doidx < 0 {
				break
			}
			idx += doidx
			continue
		}

		idx += mulidx
		if idx+4 >= len(line) {
			return 0, -1
		}
		liner := []rune(line)
		n1, n1idx := nextNumber(liner, idx+4)
		if n1idx >= 0 {
			idx = n1idx
		}

		if idx >= len(line) || idx < 0 || liner[idx] != ',' {
			continue
		}

		n2, n2idx := nextNumber(liner, idx+1)
		if n2idx >= 0 {
			idx = n2idx
		}

		if idx >= len(line) || idx < 0 || liner[idx] != ')' {
			continue
		}

		found = true
		res = n1 * n2
	}

	if found {
		return res, idx
	} else {
		return 0, -1
	}
}

func px(lines []string, nextMul func(line string) (int, int)) int {
	res := 0
	wline := strings.Join(lines, "-")

	for {
		nxt, idx := nextMul(wline)
		if idx < 0 {
			break
		}
		res += nxt

		if idx >= len(wline) {
			break
		}
		wline = wline[idx+1:]
	}
	return res
}

func p1(lines []string) int {
	return px(lines, nextMul1)
}
func p2(lines []string) int {
	return px(lines, nextMul2)
}
func Solution() int {
	parsed, err := parse.GetLines("day3/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
