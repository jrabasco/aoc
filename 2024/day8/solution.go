package day8

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
)

func p1(g *grid.Grid[rune], antennas map[rune][]grid.Point) int {
	antinodes := utils.NewSet[grid.Point]()
	for _, list := range antennas {
		for i, a := range list {
			for j, b := range list {
				if i == j {
					continue
				}
				xdiff := b.X - a.X
				ydiff := b.Y - a.Y
				antinode1 := grid.Point{a.X - xdiff, a.Y - ydiff}
				antinode2 := grid.Point{b.X + xdiff, b.Y + ydiff}
				if g.Inbound(antinode1) {
					antinodes.Add(antinode1)
				}
				if g.Inbound(antinode2) {
					antinodes.Add(antinode2)
				}
			}
		}
	}
	return len(antinodes)
}

func p2(g *grid.Grid[rune], antennas map[rune][]grid.Point) int {
	antinodes := utils.NewSet[grid.Point]()
	for _, list := range antennas {
		for i, a := range list {
			for j, b := range list {
				if i == j {
					continue
				}
				diff := grid.Vector{b.X - a.X, b.Y - a.Y}
				curA := a
				for g.Inbound(curA) {
					antinodes.Add(curA)
					curA.MoveV(diff)
				}

				diff.Reverse()
				curB := b
				for g.Inbound(curB) {
					antinodes.Add(curB)
					curB.MoveV(diff)
				}
			}
		}
	}
	return len(antinodes)
}

func Solution() int {
	parsed, err := parse.GetLinesAsRunes("day8/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	antennas := map[rune][]grid.Point{}
	g, _ := grid.NewGridAs[rune, rune](parsed, func(elm rune, x, y int) (rune, error) {
		if elm != '.' {
			pos := grid.Point{x, y}
			antennas[elm] = append(antennas[elm], pos)
		}
		return elm, nil
	})
	fmt.Printf("Part 1: %d\n", p1(&g, antennas))
	fmt.Printf("Part 1: %d\n", p2(&g, antennas))
	return 0
}
