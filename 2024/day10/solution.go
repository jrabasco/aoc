package day10

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"strconv"
	"strings"
)

type Topology struct {
	Map   grid.Grid[int]
	Heads []grid.Point
}

func score(g *grid.Grid[int], s grid.Point) int {
	res := 0
	visited := utils.NewSet[grid.Point]()
	toVisit := utils.NewQueue[grid.Point]()
	toVisit.Enqueue(s)

	for !toVisit.Empty() {
		// guarantee to work by construction, no need to check errors
		cur, _ := toVisit.Dequeue()
		if visited.Contains(cur) {
			continue
		}
		visited.Add(cur)
		curVal := *g.GetAt(cur)
		if curVal == 9 {
			res += 1
			continue
		}

		for _, p := range g.Neighbours(cur.X, cur.Y) {
			val := *g.GetAt(p)
			if val == curVal+1 && !visited.Contains(p) {
				toVisit.Enqueue(p)
			}
		}
	}
	return res
}

func rating(g *grid.Grid[int], s grid.Point) int {
	res := 0
	visited := utils.NewSet[grid.Point]()
	toVisit := utils.NewQueue[grid.Point]()
	toVisit.Enqueue(s)

	for !toVisit.Empty() {
		// guarantee to work by construction, no need to check errors
		cur, _ := toVisit.Dequeue()
		curVal := *g.GetAt(cur)
		if curVal == 9 {
			res += 1
			continue
		}

		for _, p := range g.Neighbours(cur.X, cur.Y) {
			val := *g.GetAt(p)
			if val == curVal+1 && !visited.Contains(p) {
				toVisit.Enqueue(p)
			}
		}
		visited.Add(cur)
	}
	return res
}

func p1(t *Topology) int {
	res := 0
	for _, s := range t.Heads {
		res += score(&t.Map, s)
	}
	return res
}

func p2(t *Topology) int {
	res := 0
	for _, s := range t.Heads {
		res += rating(&t.Map, s)
	}
	return res
}

func parseGrid(lines []string) (Topology, error) {
	parsed := [][]string{}
	for _, line := range lines {
		parsed = append(parsed, strings.Split(line, ""))
	}
	heads := []grid.Point{}
	g, err := grid.NewGridAs[int, string](parsed, func(elm string, i, j int) (int, error) {
		if elm == "0" {
			heads = append(heads, grid.Point{i, j})
		}
		if elm == "." {
			return -1, nil
		}
		return strconv.Atoi(elm)
	})
	return Topology{g, heads}, err
}

func Solution() int {
	t, err := parse.GetLinesAsOne[Topology]("day10/input.txt", parseGrid)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(&t))
	fmt.Printf("Part 2: %d\n", p2(&t))
	return 0
}
