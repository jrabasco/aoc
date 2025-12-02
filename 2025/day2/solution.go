package day2

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
	"strconv"
	"strings"
)

func p1(ranges []utils.Range) int {
	res := 0
	for _, r := range ranges {
		start := r.Start()
		end := r.End()
		ln := utils.IntLen(start)
		if ln%2 != 0 {
			// next number that has an even number of digits
			nxtPow10 := utils.IntPow(10, ln)
			if nxtPow10 > end {
				continue
			}

			ln = ln + 1
			start = nxtPow10 + utils.IntPow(10, ln/2-1)
			if start > end {
				continue
			}
		}

		lnEnd := utils.IntLen(end)
		if lnEnd%2 != 0 {
			// prev number that has an even number of digits
			prevPow10 := utils.IntPow(10, lnEnd-1) - 1
			if prevPow10 < start {
				continue
			}

			lnEnd = lnEnd - 1
			end = prevPow10
		}
		if ln != lnEnd {
			panic(end)
		}

		mask := utils.IntPow(10, ln/2)
		cur := start
		for cur <= end {
			upper := cur / mask
			lower := cur - upper*mask
			if upper == lower {
				res += cur
			}
			cur++
		}
	}
	return res
}

func p2(ranges []utils.Range) int {
	res := 0
	for _, r := range ranges {
		cur := r.Start()
		for cur <= r.End() {
			curStr := strconv.Itoa(cur)
			ln := len(curStr)

			for i := 1; i <= ln/2; i++ {
				if ln%i != 0 {
					continue
				}
				pattern := curStr[:i]
				tst := strings.Repeat(pattern, ln/i)
				if tst == curStr {
					res += cur
					break
				}
			}
			cur++
		}
	}
	return res
}

func parseRanges(lines []string) ([]utils.Range, error) {
	res := []utils.Range{}
	if len(lines) != 1 {
		return res, fmt.Errorf("wrong size of input: %d", len(lines))
	}

	rangeStrs := strings.Split(lines[0], ",")

	for _, rangeStr := range rangeStrs {
		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			return res, fmt.Errorf("invalid range spec: %s", rangeStr)
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return res, err
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return res, err
		}

		res = append(res, utils.NewRange(start, end))
	}
	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[[]utils.Range]("day2/input.txt", parseRanges)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
