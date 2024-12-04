package day4

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
)

func findXMas(g *grid.Grid[rune], p grid.Point, dir utils.Direction) bool {
	xmas := []rune("XMAS")
	i := 0
	for g.Inbound(p) && i < len(xmas) {
		if *g.GetAt(p) != xmas[i] {
			break
		}
		if i == len(xmas)-1 {
			return true
		}
		p = p.Move(dir, 1)
		i++
	}
	return false
}

func p1(g *grid.Grid[rune]) int {
	res := 0
	for x := 0; x <= g.MaxX(); x++ {
		for y := 0; y <= g.MaxY(); y++ {
			c := *g.Get(x, y)
			if c != 'X' {
				continue
			}
			p := grid.Point{x, y}
			for _, dir := range utils.ALL_DIRS {
				if findXMas(g, p, dir) {
					res += 1
				}
			}
		}
	}
	return res
}

func p2(g *grid.Grid[rune]) int {
	res := 0
	for x := 1; x < g.MaxX(); x++ {
		for y := 1; y < g.MaxY(); y++ {
			c := *g.Get(x, y)
			if c != 'A' {
				continue
			}
			// upleft -> downright diagonal
			upleft := *g.GetAt(grid.Point{x - 1, y - 1})
			downright := *g.GetAt(grid.Point{x + 1, y + 1})
			if upleft != 'M' && upleft != 'S' {
				continue
			}
			if upleft == 'M' && downright != 'S' {
				continue
			}
			if upleft == 'S' && downright != 'M' {
				continue
			}

			// downleft -> upright diagonal
			downleft := *g.GetAt(grid.Point{x + 1, y - 1})
			upright := *g.GetAt(grid.Point{x - 1, y + 1})
			if downleft != 'M' && downleft != 'S' {
				continue
			}
			if downleft == 'M' && upright != 'S' {
				continue
			}
			if downleft == 'S' && upright != 'M' {
				continue
			}
			res++
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
