package day4

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/grid"
	"github.com/jrabasco/aoc/2025/framework/parse"
)

func canBeRemoved(g *grid.Grid[rune], x, y int) bool {
	if *g.Get(x, y) != '@' {
		return false
	}
	busy := 0
	for _, p := range g.DNeighbours(x, y) {
		if *g.GetAt(p) == '@' {
			busy++
		}
	}
	return busy < 4
}

func p1(g *grid.Grid[rune]) int {
	res := 0
	for x := 0; x <= g.MaxX(); x++ {
		for y := 0; y <= g.MaxY(); y++ {
			if canBeRemoved(g, x, y) {
				res++
			}
		}
	}
	return res
}

func p2(g *grid.Grid[rune]) int {
	res := 0
	toRemove := []grid.Point{}
	first := true
	for first || len(toRemove) > 0 {
		first = false
		for _, p := range toRemove {
			*g.GetAt(p) = '.'
			res++
		}
		toRemove = []grid.Point{}
		for x := 0; x <= g.MaxX(); x++ {
			for y := 0; y <= g.MaxY(); y++ {
				if canBeRemoved(g, x, y) {
					toRemove = append(toRemove, grid.Point{x, y})
				}
			}
		}
	}
	return res
}

func Solution() int {
	parsed, err := parse.GetLinesAsRunes("day4/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	g := grid.NewGrid[rune](parsed)
	fmt.Printf("Part 1: %d\n", p1(&g))
	fmt.Printf("Part 2: %d\n", p2(&g))
	return 0
}
