package day5

import (
	"cmp"
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
	"slices"
	"strconv"
	"strings"
)

type Database struct {
	fresh       []utils.Range
	ingredients []int
}

func p1(db Database) int {
	res := 0
	for _, ing := range db.ingredients {
		for _, r := range db.fresh {
			if r.Contains(ing) {
				res++
				break
			}
		}
	}
	return res
}

func cmpFreshRange(a, b utils.Range) int {
	return cmp.Compare(a.Start(), b.Start())
}

func p2(db Database) int {
	res := 0
	for _, f := range db.fresh {
		res += f.Len()
	}
	return res
}

func parseRange(line string) (utils.Range, error) {
	res := utils.EmptyRange
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return res, fmt.Errorf("invalid range spec: %s", line)
	}

	st, err := strconv.Atoi(parts[0])
	if err != nil {
		return res, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return res, err
	}

	return utils.NewRange(st, end), nil
}

func parseDatabase(lines []string) (Database, error) {
	res := Database{}
	li := 0
	for ; li < len(lines); li++ {
		line := lines[li]
		if line == "" {
			break
		}
		rng, err := parseRange(line)
		if err != nil {
			return res, err
		}
		res.fresh = append(res.fresh, rng)
	}
	li++

	for ; li < len(lines); li++ {
		ing, err := strconv.Atoi(lines[li])
		if err != nil {
			return res, err
		}
		res.ingredients = append(res.ingredients, ing)
	}

	slices.SortFunc(res.fresh, cmpFreshRange)
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(res.fresh)-1; i++ {
			if res.fresh[i].Overlaps(res.fresh[i+1]) {
				merged := res.fresh[i].Merge(res.fresh[i+1])
				// remove element i
				res.fresh = append(res.fresh[:i], res.fresh[i+1:]...)
				// now override i+1 with new one
				res.fresh[i] = merged
				changed = true
			}
		}
	}
	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[Database]("day5/input.txt", parseDatabase)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
