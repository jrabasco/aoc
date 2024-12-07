package day6

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
)

func p1(g *grid.Grid[rune], start grid.Point, obstacles utils.Set[grid.Point]) utils.Set[grid.Point] {
	curP := start
	curDir := utils.UP
	visited := utils.NewSet[grid.Point]()
	for g.Inbound(curP) {
		visited.Add(curP)
		nextP := curP.Move(curDir, 1)
		// allow it to go out of bounds
		if !g.Inbound(nextP) || !obstacles.Contains(nextP) {
			curP = nextP
		} else {
			curDir = curDir.Turn(utils.RIGHT)
		}
	}
	return visited
}

type PosDir struct {
	p   grid.Point
	dir utils.Direction
}

func hasCycle(g *grid.Grid[rune], start grid.Point, startDir utils.Direction, obstacles utils.Set[grid.Point]) bool {
	curP := start
	curDir := startDir
	visited := utils.NewSet[PosDir]()
	for g.Inbound(curP) {
		state := PosDir{curP, curDir}
		if visited.Contains(state) {
			return true
		}
		visited.Add(state)
		nextP := curP.Move(curDir, 1)
		// allow it to go out of bounds
		if !g.Inbound(nextP) || !obstacles.Contains(nextP) {
			curP = nextP
		} else {
			curDir = curDir.Turn(utils.RIGHT)
		}
	}
	return false
}

func p2(g *grid.Grid[rune], start grid.Point, guardPath utils.Set[grid.Point], obstacles utils.Set[grid.Point]) int {
	res := utils.NewSet[grid.Point]()
	for pos := range guardPath {
		if pos == start {
			continue
		}
		nObstacles := utils.NewSet[grid.Point]()
		nObstacles.AddSet(obstacles)
		nObstacles.Add(pos)
		if hasCycle(g, start, utils.UP, nObstacles) {
			res.Add(pos)
		}
	}
	return len(res)
}
func Solution() int {
	parsed, err := parse.GetLinesAsRunes("day6/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	var start grid.Point
	obstacles := utils.NewSet[grid.Point]()
	g, _ := grid.NewGridAs[rune, rune](parsed, func(elm rune, x, y int) (rune, error) {
		if elm == '^' {
			start = grid.Point{x, y}
		}

		if elm == '#' {
			obstacles.Add(grid.Point{x, y})
		}
		return elm, nil
	})
	guardPath := p1(&g, start, obstacles)
	fmt.Printf("Part 1: %d\n", len(guardPath))
	fmt.Printf("Part 2: %d\n", p2(&g, start, guardPath, obstacles))
	return 0
}
