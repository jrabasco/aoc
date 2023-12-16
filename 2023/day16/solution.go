package day16

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
)

type Direction int

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
)

type Tile struct {
	visited []bool
	val     rune
}

func NewTile(val rune) Tile {
	return Tile{[]bool{false, false, false, false}, val}
}

func parseGrid(lines [][]rune) grid.Grid[Tile] {
	g, _ := grid.NewGridAs[Tile, rune](lines, func(elm rune, i, j int) (Tile, error) {
		return NewTile(elm), nil
	})
	return g
}

func printG(g grid.Grid[Tile]) {
	for _, row := range g.Rows() {
		rowStr := ""
		for _, tile := range row {
			visited := false
			for _, v := range tile.visited {
				if v {
					visited = true
					break
				}
			}
			if visited {
				rowStr += "#"
			} else {
				rowStr += string(tile.val)
			}
		}
		fmt.Println(rowStr)
	}
}

func energy(g grid.Grid[Tile]) int {
	res := 0
	for _, row := range g.Rows() {
		for _, tile := range row {
			for _, v := range tile.visited {
				if v {
					res += 1
					break
				}
			}
		}
	}
	return res
}

func reset(g *grid.Grid[Tile]) {
	for _, row := range g.Rows() {
		for _, tile := range row {
			for i := range tile.visited {
				tile.visited[i] = false
			}
		}
	}
}

type Beam struct {
	i   int
	j   int
	dir Direction
}

func followBeams(g *grid.Grid[Tile], direction Direction, i, j int) {
	beams := utils.NewQueue[Beam]()
	beams.Enqueue(Beam{i, j, direction})
	for !beams.Empty() {
		beam, _ := beams.Dequeue()
		if beam.i < g.H() && beam.j < g.W() && beam.i >= 0 && beam.j >= 0 {
			bi := beam.i
			bj := beam.j
			tile := g.Get(beam.i, beam.j)
			if tile.visited[beam.dir] {
				continue
			}
			tile.visited[beam.dir] = true
			if beam.dir == RIGHT {
				if tile.val == '/' {
					beams.Enqueue(Beam{bi - 1, bj, UP})
					continue
				}

				if tile.val == '\\' {
					beams.Enqueue(Beam{bi + 1, bj, DOWN})
					continue
				}

				if tile.val == '|' {
					beams.Enqueue(Beam{bi - 1, bj, UP})
					beams.Enqueue(Beam{bi + 1, bj, DOWN})
					continue
				}
				beams.Enqueue(Beam{bi, bj + 1, RIGHT})
			} else if beam.dir == LEFT {
				if tile.val == '/' {
					beams.Enqueue(Beam{bi + 1, bj, DOWN})
					continue
				}

				if tile.val == '\\' {
					beams.Enqueue(Beam{bi - 1, bj, UP})
					continue
				}

				if tile.val == '|' {
					beams.Enqueue(Beam{bi - 1, bj, UP})
					beams.Enqueue(Beam{bi + 1, bj, DOWN})
					continue
				}
				beams.Enqueue(Beam{bi, bj - 1, LEFT})
			} else if beam.dir == DOWN {
				if tile.val == '/' {
					beams.Enqueue(Beam{bi, bj - 1, LEFT})
					continue
				}

				if tile.val == '\\' {
					beams.Enqueue(Beam{bi, bj + 1, RIGHT})
					continue
				}

				if tile.val == '-' {
					beams.Enqueue(Beam{bi, bj - 1, LEFT})
					beams.Enqueue(Beam{bi, bj + 1, RIGHT})
					continue
				}
				beams.Enqueue(Beam{bi + 1, bj, DOWN})
			} else if beam.dir == UP {
				if tile.val == '/' {
					beams.Enqueue(Beam{bi, bj + 1, RIGHT})
					continue
				}

				if tile.val == '\\' {
					beams.Enqueue(Beam{bi, bj - 1, LEFT})
					continue
				}

				if tile.val == '-' {
					beams.Enqueue(Beam{bi, bj - 1, LEFT})
					beams.Enqueue(Beam{bi, bj + 1, RIGHT})
					continue
				}
				beams.Enqueue(Beam{bi - 1, bj, UP})
			}
		}
	}
}

func Solution() int {
	parsed, err := parse.GetLinesAs[[]rune]("day16/input.txt", func(line string) ([]rune, error) {
		return []rune(line), nil
	})
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	g := parseGrid(parsed)
	followBeams(&g, RIGHT, 0, 0)
	fmt.Printf("Part 1: %d\n", energy(g))

	res2 := 0
	for i := 0; i <= g.MaxX(); i++ {
		for j := 0; j <= g.MaxY(); j++ {
			if i != 0 && j != 0 && i != g.MaxX() && j != g.MaxY() {
				// not on the border
				continue
			}

			if i == 0 {
				reset(&g)
				followBeams(&g, DOWN, i, j)
				res2 = max(res2, energy(g))
			}
			if i == g.MaxX() {
				reset(&g)
				followBeams(&g, UP, i, j)
				res2 = max(res2, energy(g))
			}
			if j == 0 {
				reset(&g)
				followBeams(&g, RIGHT, i, j)
				res2 = max(res2, energy(g))
			}
			if j == g.MaxY() {
				reset(&g)
				followBeams(&g, LEFT, i, j)
				res2 = max(res2, energy(g))
			}
		}
	}
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
