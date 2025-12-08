package day8

import (
	"cmp"
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
	"slices"
	"strconv"
	"strings"
)

type LinkCandidate struct {
	p1   [3]int
	p2   [3]int
	dist int
}

func sqDist(a, b [3]int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	z := a[2] - b[2]
	return x*x + y*y + z*z
}

func solve(junctions [][3]int) (int, int) {
	p1stop := 10
	if len(junctions) >= 1000 {
		p1stop = 1000
	}
	cands := []LinkCandidate{}
	for i := range junctions {
		for j := i; j < len(junctions); j++ {
			if i == j {
				continue
			}
			a := junctions[i]
			b := junctions[j]
			cands = append(cands, LinkCandidate{a, b, sqDist(a, b)})
		}
	}
	dCmp := func(a, b LinkCandidate) int {
		return cmp.Compare(a.dist, b.dist)
	}
	slices.SortFunc(cands, dCmp)
	sets := map[[3]int]*utils.Set[[3]int]{}
	res1 := 0
	for i := range cands {
		if i == p1stop {
			seen := utils.NewSet[[3]int]()
			sizes := []int{}
			for junc, set := range sets {
				if seen.Contains(junc) {
					continue
				}
				sizes = append(sizes, len(*set))
				for elm := range *set {
					seen.Add(elm)
				}
			}
			slices.Sort(sizes)
			ls := len(sizes)
			res1 = sizes[ls-1] * sizes[ls-2] * sizes[ls-3]
		}
		cand := cands[i]
		found := false
		s1, ok1 := sets[cand.p1]
		s2, ok2 := sets[cand.p2]

		if !ok1 && !ok2 {
			s := utils.NewSet[[3]int]()
			s.Add(cand.p1, cand.p2)
			sets[cand.p1] = &s
			sets[cand.p2] = &s
		}

		if ok1 && !ok2 {
			s1.Add(cand.p2)
			sets[cand.p2] = s1
			if len(*s1) == len(junctions) {
				found = true
			}
		}

		if !ok1 && ok2 {
			s2.Add(cand.p1)
			sets[cand.p1] = s2
			if len(*s2) == len(junctions) {
				found = true
			}
		}

		if ok1 && ok2 {
			s := s1.Union(*s2)
			for elm := range s {
				sets[elm] = &s
			}
			if len(s) == len(junctions) {
				found = true
			}
		}

		if found {
			return res1, cand.p1[0] * cand.p2[0]
		}
	}

	return res1, 0
}

func parseLine(line string) ([3]int, error) {
	parts := strings.Split(line, ",")
	res := [3]int{0, 0, 0}
	if len(parts) != 3 {
		return res, fmt.Errorf("invalid junction: %s", line)
	}

	for i := range parts {
		nb, err := strconv.Atoi(parts[i])
		if err != nil {
			return res, err
		}
		res[i] = nb
	}

	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[[3]int]("day8/input.txt", parseLine)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	res1, res2 := solve(parsed)
	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
