package day12

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/grid"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
)

type Square struct {
	val     rune
	isGoal  bool
	visited bool
	prev    *Square
}

func NewSquare(r rune, isGoal bool) Square {
	return Square{r, isGoal, false, nil}
}

func backTrace(square *Square) []*Square {
	cur := square
	res := []*Square{cur}
	for cur.prev != nil {
		res = append(res, cur.prev)
		cur = cur.prev
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func findPath(g *grid.Grid[Square], x int, y int, isGoal func(*Square) bool, canGo func(rune, rune) bool) []*Square {
	// queue of paths
	queue := utils.NewQueue[grid.Point]()
	queue.Enqueue(grid.Point{x, y})
	g.Get(x, y).visited = true

	for !queue.Empty() {
		// ignore error because of loop condition
		curP, _ := queue.Dequeue()
		cur := g.Get(curP.X, curP.Y)

		if isGoal(cur) {
			return backTrace(cur)
		}

		for _, neighP := range g.Neighbours(curP.X, curP.Y) {
			neigh := g.Get(neighP.X, neighP.Y)
			if !canGo(cur.val, neigh.val) {
				continue
			}

			if neigh.visited {
				continue
			}

			neigh.visited = true
			neigh.prev = cur
			queue.Enqueue(neighP)
		}
	}
	return nil
}

func solvePart(part string) int {
	parsed, err := parse.GetLinesAs[[]rune]("day12/input.txt",
		func(s string) ([]rune, error) {
			res := []rune{}
			for _, r := range s {
				res = append(res, r)
			}
			return res, nil
		})
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	startX := -1
	startY := -1

	endX := -1
	endY := -1
	g, err := grid.NewGrid[Square, rune](parsed, func(r rune, x, y int) (Square, error) {
		el := r
		if r == 'S' {
			el = 'a'
			startX = x
			startY = y
		}

		isGoal := false
		if r == 'E' {
			el = 'z'
			isGoal = true
			endX = x
			endY = y
		}
		return NewSquare(el, isGoal), nil
	})
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	path := []*Square{}

	if part == "1" {
		path = findPath(&g, startX, startY, func(s *Square) bool { return s.isGoal }, func(v1, v2 rune) bool { return v2-v1 <= 1 })
	} else {
		path = findPath(&g, endX, endY, func(s *Square) bool { return s.val == 'a' }, func(v1, v2 rune) bool { return v1-v2 <= 1 })
	}

	// path contains start so need -1
	fmt.Printf("Part %s: %v\n", part, len(path)-1)
	return 0
}

func Solution(part string) int {
	if part != "1" && part != "2" {
		p1 := solvePart("1")
		if p1 != 0 {
			return p1
		}
		return solvePart("2")
	} else {
		return solvePart(part)
	}
}
