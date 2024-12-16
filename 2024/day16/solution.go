package day16

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"strings"
)

type Problem struct {
	g     grid.Grid[string]
	start grid.Point
	end   grid.Point
}

var UNDEFINED = grid.Point{-1, -1}

type Cost struct {
	cost int
	dir  utils.Direction
}

func findSmallest(q *[]grid.Point, dist *map[grid.Point]Cost) int {
	idx := -1
	minCost := -1

	for i, elm := range *q {
		ci := (*dist)[elm]
		if minCost == -1 {
			idx = i
			minCost = ci.cost
			continue
		}
		if ci.cost == -1 {
			continue
		}
		if ci.cost < minCost {
			idx = i
			minCost = ci.cost
		}
	}
	return idx
}

func dijkstra(g *grid.Grid[string], start grid.Point) {
	dist := map[grid.Point]Cost{}
	prev := map[grid.Point]grid.Point{}
	Q := []grid.Point{}
	considering := utils.NewSet[grid.Point]()
	for i, row := range g.Rows() {
		for j, elm := range row {
			if *elm == "#" {
				continue
			}
			cur := grid.Point{i, j}
			dist[cur] = Cost{-1, utils.NODIR}
			prev[cur] = UNDEFINED
			Q = append(Q, cur)
			considering.Add(cur)
		}
	}
	dist[start] = Cost{0, utils.RIGHT}

	for len(Q) > 0 {
		ui := findSmallest(&Q, &dist)
		u := Q[ui]
		// remove from Q
		Q[ui] = Q[len(Q)-1]
		Q = Q[:len(Q)-1]
		considering.Remove(u)

		for _, v := range g.Neighbours(u.X, u.Y) {
			elm := g.GetAt(v)
			if *elm == "#" || !considering.Contains(v) {
				continue
			}
			cu := dist[u]
			costUV := 0
			dir := cu.dir
			if u.Move(cu.dir, 1) == v {
				costUV = 1
			} else {
				costUV = 1001
				if u.X < v.X {
					dir = utils.DOWN
				} else if u.X > v.X {
					dir = utils.UP
				} else if u.Y < v.Y {
					dir = utils.RIGHT
				} else if u.Y > v.Y {
					dir = utils.LEFT
				}
			}

			cv := dist[v]

			if cv.cost == -1 || cv.cost > cu.cost+costUV {
				dist[v] = Cost{cu.cost + costUV, dir}
				prev[v] = u
			}
		}
	}
	cur := grid.Point{1, 13}
	for cur != start {
		fmt.Println(cur, dist[cur])
		cur = prev[cur]
	}
}

func p1(p Problem) int {
	g := grid.Copy[string](p.g)
	fmt.Println(g)
	fmt.Println(p.start)
	dijkstra(&g, p.start)
	//dir := utils.RIGHT
	//fmt.Println(dir)
	//q := utils.NewQueue[grid.Point]()
	//q.Enqueue(start)

	//for !q.Empty() {
	//	cur,_ := q.Dequeue()
	//	for _, n := range g.Neighbours(cur.X, cur.Y) {
	//		if !visited.Contains(n) {
	//			q.Enqueue(n)
	//		}
	//	}
	//	visited.Add(cur)
	//}

	return 0
}

func p2(p Problem) int {
	return 0
}

func parseProblem(lines []string) (Problem, error) {
	elms := [][]string{}
	for _, line := range lines {
		elms = append(elms, strings.Split(line, ""))
	}

	var start, end grid.Point

	g, _ := grid.NewGridAs[string, string](elms, func(c string, i, j int) (string, error) {
		if c == "S" {
			start = grid.Point{i, j}
		}
		if c == "E" {
			end = grid.Point{i, j}
		}
		return c, nil
	})
	return Problem{g, start, end}, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[Problem]("day16/input.txt", parseProblem)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
