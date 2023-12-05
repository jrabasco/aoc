package day5

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strconv"
	"strings"
)

type Map map[utils.Range]int // map a range to the delta

func (m Map) findDelta(elm int) int {
	for rng, delta := range m {
		if rng.Contains(elm) {
			return delta
		}
	}
	return 0
}

func (m Map) transform(r utils.Range) []utils.Range {
	res := []utils.Range{}
	q := utils.NewQueue[utils.Range]()
	q.Enqueue(r)

	for !q.Empty() {
		// ignore err since for condition prevents empty err
		elm, _ := q.Dequeue()
		foundMatch := false
		for rng, delta := range m {
			match, rest := rng.Split(elm)
			if !match.Empty() {
				foundMatch = true
				nElm := utils.NewRange(match.Start()+delta, match.End()+delta)
				res = append(res, nElm)

				for _, todo := range rest {
					q.Enqueue(todo)
				}
				break
			}
		}

		if !foundMatch {
			res = append(res, elm)
		}
	}
	return res
}

type Almanac struct {
	seeds      []int
	seedRanges []utils.Range
	phases     []Map
}

func parseAlmanac(lines []string) (Almanac, error) {
	res := Almanac{[]int{}, []utils.Range{}, []Map{}}

	seedsStr := strings.TrimPrefix(lines[0], "seeds: ")
	seedsParts := strings.Split(seedsStr, " ")
	for _, seedStr := range seedsParts {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			return res, err
		}
		res.seeds = append(res.seeds, seed)
	}

	for i := 0; i < len(res.seeds)/2; i++ {
		rng := utils.NewRange(res.seeds[2*i], res.seeds[2*i]+res.seeds[2*i+1]-1)
		res.seedRanges = append(res.seedRanges, rng)
	}

	i := 2
	for i < len(lines) {
		mp, ni, err := parseMap(i+1, lines)
		if err != nil {
			return res, err
		}
		res.phases = append(res.phases, mp)
		i = ni
		i++
	}
	return res, nil
}

func parseMap(start int, lines []string) (Map, int, error) {
	i := start
	res := Map{}
	for i < len(lines) && lines[i] != "" {
		parts := strings.Split(lines[i], " ")
		if len(parts) != 3 {
			return res, i, fmt.Errorf("malformed line: %s", lines[i])
		}

		dst, err := strconv.Atoi(parts[0])
		if err != nil {
			return res, i, err
		}
		src, err := strconv.Atoi(parts[1])
		if err != nil {
			return res, i, err
		}
		ln, err := strconv.Atoi(parts[2])
		if err != nil {
			return res, i, err
		}

		rng := utils.NewRange(src, src+ln-1) // inclusive boundaries
		res[rng] = dst - src                 // delta
		i++
	}
	return res, i, nil // i is now at the next empty line
}

func Solution() int {
	parsed, err := parse.GetLines("day5/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	almanac, err := parseAlmanac(parsed)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	res := 9223372036854775807

	for _, seed := range almanac.seeds {
		cur := seed
		for _, phase := range almanac.phases {
			cur = cur + phase.findDelta(cur)
		}

		if cur < res {
			res = cur
		}
	}
	fmt.Printf("Part 1: %d\n", res)

	res = 9223372036854775807
	for _, seedR := range almanac.seedRanges {
		curs := []utils.Range{seedR}
		for _, phase := range almanac.phases {
			ncurs := []utils.Range{}
			for _, cur := range curs {
				nparts := phase.transform(cur)
				ncurs = append(ncurs, nparts...)
			}
			curs = ncurs
		}

		for _, loc := range curs {
			if loc.Start() < res {
				res = loc.Start()
			}
		}
	}
	fmt.Printf("Part 2: %d\n", res)
	return 0
}
