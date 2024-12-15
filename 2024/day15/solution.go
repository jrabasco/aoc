package day15

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"strings"
)

type Problem struct {
	g     grid.Grid[string]
	moves []utils.Direction
	start grid.Point
}

func seekEmpty(g *grid.Grid[string], start grid.Point, dir utils.Direction) int {
	count := 1
	cur := start.Move(dir, 1)
	for g.Inbound(cur) {
		if *g.GetAt(cur) == "#" {
			return -1
		}
		if *g.GetAt(cur) == "." {
			return count
		}
		count++
		cur = cur.Move(dir, 1)
	}
	return -1
}

func p1(p Problem) int {
	cur := p.start
	g := grid.Copy[string](p.g)
	steps := 0
	for _, m := range p.moves {
		diff := seekEmpty(&g, cur, m)
		if diff == -1 {
			continue
		}

		empty := cur.Move(m, diff)
		back := m.Reverse()
		for i := diff; i > 0; i-- {
			n := empty.Move(back, 1)
			toMove := *g.GetAt(n)
			*g.GetAt(empty) = toMove
			empty = n
		}
		*g.GetAt(empty) = "."
		steps++
		cur = cur.Move(m, 1)
	}
	res := 0
	for i, r := range g.Rows() {
		for j, elm := range r {
			if *elm == "O" {
				res += 100*i + j
			}
		}
	}
	return res
}

func canMove(g *grid.SparseGrid[string], pos grid.Point, dir utils.Direction) bool {
	npos := pos.Move(dir, 1)
	elm := g.GetP(npos)
	if elm == "#" {
		return false
	}

	if elm == "[" {
		if dir == utils.RIGHT || dir == utils.LEFT {
			return canMove(g, npos, dir)
		}

		nposRight := grid.Point{npos.X, npos.Y + 1}
		return canMove(g, npos, dir) && canMove(g, nposRight, dir)
	}

	if elm == "]" {
		if dir == utils.RIGHT || dir == utils.LEFT {
			return canMove(g, npos, dir)
		}

		nposLeft := grid.Point{npos.X, npos.Y - 1}
		return canMove(g, npos, dir) && canMove(g, nposLeft, dir)
	}

	// then it's empty
	return true
}

func move(g *grid.SparseGrid[string], pos grid.Point, dir utils.Direction) {
	npos := pos.Move(dir, 1)
	elm := g.GetP(npos)
	if elm == "#" {
		return
	}

	if elm == "[" {
		if dir == utils.RIGHT || dir == utils.LEFT {
			move(g, npos, dir)
		} else {
			nposRight := grid.Point{npos.X, npos.Y + 1}
			move(g, npos, dir)
			move(g, nposRight, dir)
		}
	}

	if elm == "]" {
		if dir == utils.RIGHT || dir == utils.LEFT {
			move(g, npos, dir)
		} else {
			nposLeft := grid.Point{npos.X, npos.Y - 1}
			move(g, npos, dir)
			move(g, nposLeft, dir)
		}
	}

	// spot has been freed
	toMove := g.GetP(pos)
	g.RemoveP(pos)
	g.AddP(npos, toMove)
}

func p2(p Problem) int {
	g2 := grid.NewSparseGrid[string](".")
	var start2 grid.Point
	for i, row := range p.g.Rows() {
		for j, elm := range row {
			if *elm == "#" {
				g2.Add(i, 2*j, "#")
				g2.Add(i, 2*j+1, "#")
			}

			if *elm == "O" {
				g2.Add(i, 2*j, "[")
				g2.Add(i, 2*j+1, "]")
			}

			if *elm == "@" {
				g2.Add(i, 2*j, "@")
				start2 = grid.Point{i, 2 * j}
			}
		}
	}

	for _, m := range p.moves {
		if !canMove(&g2, start2, m) {
			continue
		}
		move(&g2, start2, m)
		start2 = start2.Move(m, 1)
		//if i > 0 {
		//	break
		//}
	}

	res := 0
	for i := 0; i <= g2.MaxX(); i++ {
		for j := 0; j <= g2.MaxY(); j++ {
			elm := g2.Get(i, j)
			if elm != "[" {
				continue
			}
			res += 100*i + j
		}
	}
	return res
}

func parseGrid(lines []string) (Problem, error) {
	warehouse := [][]string{}
	i := 0
	for ; i < len(lines) && len(lines[i]) > 0; i++ {
		warehouse = append(warehouse, strings.Split(lines[i], ""))
	}

	var start grid.Point
	g, _ := grid.NewGridAs[string, string](warehouse, func(c string, i, j int) (string, error) {
		if c == "@" {
			start = grid.Point{i, j}
		}
		return c, nil
	})

	moves := []utils.Direction{}
	for ; i < len(lines); i++ {
		for _, m := range strings.Split(lines[i], "") {
			if m == ">" {
				moves = append(moves, utils.RIGHT)
			}

			if m == "v" {
				moves = append(moves, utils.DOWN)
			}

			if m == "<" {
				moves = append(moves, utils.LEFT)
			}

			if m == "^" {
				moves = append(moves, utils.UP)
			}
		}
	}
	return Problem{g, moves, start}, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[Problem]("day15/input.txt", parseGrid)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
