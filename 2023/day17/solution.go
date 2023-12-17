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
		straight int
		heatLoss int
	}
	//type visitedEntry struct {
	//	pos      grid.Point
	//	dir      utils.Direction
	//	straight int
	//}

	q := utils.NewPriorityQueue[queueEntry](func(a, b queueEntry) int {
		return a.heatLoss - b.heatLoss
	})
	// try without a map since profiler says map is main cost
	visited := [][][][]int{}
	for i := 0; i < g.H(); i++ {
		visited = append(visited, [][][]int{})
		for j := 0; j < g.W(); j++ {
			visited[i] = append(visited[i], [][]int{})
			for d := 0; d < 4; d++ {
				visited[i][j] = append(visited[i][j], make([]int, maxStraight))
			}
		}
	}
	q.Enqueue(queueEntry{
		pos:      grid.Point{0, 1},
		straight: 1,
		dir:      utils.RIGHT,
	})
	q.Enqueue(queueEntry{
		pos:      grid.Point{1, 0},
		straight: 1,
		dir:      utils.DOWN,
	})

	count := 0
	for !q.Empty() {
		count += 1
		cur, _ := q.Dequeue()

		if !g.Inbound(cur.pos) {
			continue
		}

		heat := *g.GetAt(cur.pos) + cur.heatLoss
		if cur.pos == target && cur.straight >= minStraight {
			return heat
		}

		if v := visited[cur.pos.X][cur.pos.Y][cur.dir][cur.straight-1]; v != 0 {
			// already found a better path to get there
			if v <= heat {
				continue
			}
		}
		visited[cur.pos.X][cur.pos.Y][cur.dir][cur.straight-1] = heat

		if cur.straight >= minStraight {
			left := cur.dir.Turn(utils.LEFT)
			q.Enqueue(queueEntry{
				pos:      cur.pos.Move(left, 1),
				dir:      left,
				heatLoss: heat,
				straight: 1,
			})

			right := cur.dir.Turn(utils.RIGHT)
			q.Enqueue(queueEntry{
				pos:      cur.pos.Move(right, 1),
				dir:      right,
				heatLoss: heat,
				straight: 1,
			})
		}

		if cur.straight < maxStraight {
			q.Enqueue(queueEntry{
				pos:      cur.pos.Move(cur.dir, 1),
				dir:      cur.dir,
				heatLoss: heat,
				straight: cur.straight + 1,
			})
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
