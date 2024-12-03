package day17

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strconv"
	"strings"
)

func bfs(g grid.Grid[int], maxStraight, minStraight int) int {
	target := grid.Point{g.MaxX(), g.MaxY()}
	type queueEntry struct {
		pos      grid.Point
		dir      utils.Direction
		heatLoss int
	}

	q := utils.NewPriorityQueue[queueEntry](func(a, b queueEntry) int {
		return a.heatLoss - b.heatLoss
	})
	// We could use a map here but profiler pointed at a lot of time spent in
	// map allocations.
	// This array contains heat generated by going to any cell, given any
	// direction, we are not caching on length of straight line because we
	// directly enqueue the possible destinations.
	visited := [][][]int{}
	for i := 0; i < g.H(); i++ {
		visited = append(visited, [][]int{})
		for j := 0; j < g.W(); j++ {
			visited[i] = append(visited[i], []int{-1, -1, -1, -1})
		}
	}
	minStep := max(minStraight, 1)
	rpos := grid.Point{0, 1}
	rheat := 0
	// add all the heats for the steps we are skipping
	for i := 1; i < minStep; i++ {
		rheat += *g.GetAt(rpos)
		rpos = rpos.Move(utils.RIGHT, 1)
	}
	for step := minStep; step <= maxStraight; step++ {
		if !g.Inbound(rpos) {
			break
		}
		q.Enqueue(queueEntry{
			pos:      rpos,
			dir:      utils.RIGHT,
			heatLoss: rheat,
		})
		rheat += *g.GetAt(rpos)
		rpos = rpos.Move(utils.RIGHT, 1)
	}

	lpos := grid.Point{1, 0}
	lheat := 0
	// add all the heats for the steps we are skipping
	for i := 1; i < minStep; i++ {
		lheat += *g.GetAt(lpos)
		lpos = lpos.Move(utils.DOWN, 1)
	}
	for step := minStep; step <= maxStraight; step++ {
		if !g.Inbound(lpos) {
			break
		}
		q.Enqueue(queueEntry{
			pos:      lpos,
			dir:      utils.DOWN,
			heatLoss: lheat,
		})
		lheat += *g.GetAt(lpos)
		lpos = lpos.Move(utils.DOWN, 1)
	}

	for !q.Empty() {
		cur, _ := q.Dequeue()
		heat := *g.GetAt(cur.pos) + cur.heatLoss
		if cur.pos == target {
			return heat
		}

		if v := visited[cur.pos.X][cur.pos.Y][cur.dir]; v != -1 {
			// already found a better path to get there
			if v <= heat {
				continue
			}
		}
		visited[cur.pos.X][cur.pos.Y][cur.dir] = heat

		left := cur.dir.Turn(utils.LEFT)
		lpos := cur.pos.Move(left, 1)
		lheat := heat
		// add all the heats for the steps we are skipping
		for i := 1; i < minStep; i++ {
			if !g.Inbound(lpos) {
				break
			}
			lheat += *g.GetAt(lpos)
			lpos = lpos.Move(left, 1)
		}

		for step := minStep; step <= maxStraight; step++ {
			if !g.Inbound(lpos) {
				break
			}
			heat2 := *g.GetAt(lpos) + lheat
			if v := visited[lpos.X][lpos.Y][left]; v > heat2 || v == -1 {
				q.Enqueue(queueEntry{
					pos:      lpos,
					dir:      left,
					heatLoss: lheat,
				})
			}
			lheat += *g.GetAt(lpos)
			lpos = lpos.Move(left, 1)
		}
		right := cur.dir.Turn(utils.RIGHT)
		rpos := cur.pos.Move(right, 1)
		rheat := heat
		// add all the heats for the steps we are skipping
		for i := 1; i < minStep; i++ {
			if !g.Inbound(rpos) {
				break
			}
			rheat += *g.GetAt(rpos)
			rpos = rpos.Move(right, 1)
		}
		for step := minStep; step <= maxStraight; step++ {
			if !g.Inbound(rpos) {
				break
			}
			heat2 := *g.GetAt(rpos) + rheat
			if v := visited[rpos.X][rpos.Y][right]; v > heat2 || v == -1 {
				q.Enqueue(queueEntry{
					pos:      rpos,
					dir:      right,
					heatLoss: rheat,
				})
			}
			rheat += *g.GetAt(rpos)
			rpos = rpos.Move(right, 1)
		}
	}
	return 0
}

func parseGrid(lines []string) (grid.Grid[int], error) {
	parsed := [][]string{}
	for _, line := range lines {
		parsed = append(parsed, strings.Split(line, ""))
	}
	return grid.NewGridAs[int, string](parsed, func(elm string, i, j int) (int, error) {
		return strconv.Atoi(elm)
	})
}

func Solution() int {
	g, err := parse.GetLinesAsOne[grid.Grid[int]]("day17/input.txt", parseGrid)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	res := bfs(g, 3, 0)
	fmt.Printf("Part 1: %d\n", res)
	res2 := bfs(g, 10, 4)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}