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
	steps := 0
	for _, m := range p.moves {
		diff := seekEmpty(&p.g, cur, m)
		if diff == -1 {
			//fmt.Println("Do nothing")
			continue
		}

		empty := cur.Move(m, diff)
		back := m.Reverse()
		//fmt.Println("Moving things", m)
		for i := diff; i > 0; i-- {
			n := empty.Move(back, 1)
			//fmt.Println("Move things from:", n, "to:", empty)
			toMove := *p.g.GetAt(n)
			*p.g.GetAt(empty) = toMove
			empty = n
		}
		*p.g.GetAt(empty) = "."
		steps++
		cur = cur.Move(m, 1)
		//if steps == 2 {
		//	break
		//}
		//fmt.Println(p.g)
	}
	res := 0
	for i, r := range p.g.Rows() {
		for j, elm := range r {
			if *elm == "O" {
				res += 100*i + j
			}
		}
	}
	return res
}

func p2(p Problem) int {
	return 0
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
