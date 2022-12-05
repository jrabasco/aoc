package day3

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
)

type Compartment = utils.Set[rune]

type Rucksack struct {
	c1  Compartment
	c2  Compartment
	tot Compartment
}

func (r Rucksack) String() string {
	return fmt.Sprintf("%s|%s", r.c1, r.c2)
}

func makeRucksack(s string) (Rucksack, error) {
	cutoff := len(s) / 2
	c1 := make(Compartment)
	c2 := make(Compartment)
	tot := make(Compartment)

	for i, v := range s {
		if i < cutoff {
			c1.Add(v)
		} else {
			c2.Add(v)
		}
		tot.Add(v)
	}
	return Rucksack{c1, c2, tot}, nil
}

func getPrio(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}
}

func Solution(part string) int {
	ruckSacks, err := parse.GetLinesAs[Rucksack]("day3/input.txt", makeRucksack)
	if err != nil {
		fmt.Printf("Cannot parse input: %v\n", err)
		return 1
	}

	if part != "2" {
		// part1
		part1 := 0
		for _, r := range ruckSacks {
			common := r.c1.Intersect(r.c2).Peek()
			part1 += getPrio(common)
		}
		fmt.Printf("Part 1: %d\n", part1)
	}

	if part != "1" {
		// part2
		if len(ruckSacks)%3 != 0 {
			fmt.Printf("Invalid number of sacks: %d\n", len(ruckSacks))
			return 1
		}
		part2 := 0
		for i := 0; i < len(ruckSacks)/3; i++ {
			r1 := ruckSacks[3*i]
			r2 := ruckSacks[3*i+1]
			r3 := ruckSacks[3*i+2]
			common := r1.tot.Intersect(r2.tot).Intersect(r3.tot).Peek()
			part2 += getPrio(common)
		}
		fmt.Printf("Part 2: %d\n", part2)
	}
	return 0
}
