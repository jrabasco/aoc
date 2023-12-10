package day10

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
)

type Dir int

const (
	NORTH Dir = iota
	SOUTH
	WEST
	EAST
)

// used to swap from/to, if you're going TO your NORTH neighbour
// then you came FROM the SOUTH
func reflect(d Dir) Dir {
	if d == NORTH {
		return SOUTH
	} else if d == SOUTH {
		return NORTH
	} else if d == WEST {
		return EAST
	} else {
		return WEST
	}
}

func closeLoop(n1 Dir, n2 Dir) rune {
	if (n1 == NORTH && n2 == SOUTH) || (n2 == NORTH && n1 == SOUTH) {
		return '|'
	} else if (n1 == NORTH && n2 == EAST) || (n2 == NORTH && n1 == EAST) {
		return 'L'
	} else if (n1 == EAST && n2 == WEST) || (n2 == EAST && n1 == WEST) {
		return '-'
	} else if (n1 == NORTH && n2 == WEST) || (n2 == NORTH && n1 == WEST) {
		return 'J'
	} else if (n1 == SOUTH && n2 == WEST) || (n2 == SOUTH && n1 == WEST) {
		return '7'
	} else {
		return 'F'
	}
}

// move to the next position that's not the previous one
func next(g *grid.Grid[rune], from grid.Point, prev Dir) (grid.Point, Dir, error) {
	val := *g.Get(from.X, from.Y)
	if prev == NORTH {
		if val == '|' {
			return g.Down(from.X, from.Y, false), SOUTH, nil
		}
		if val == 'L' {
			return g.Right(from.X, from.Y, false), EAST, nil
		}
		if val == 'J' {
			return g.Left(from.X, from.Y, false), WEST, nil
		}
		return grid.Point{0, 0}, NORTH, fmt.Errorf("coming from north, found: %s", string(val))
	} else if prev == SOUTH {
		if val == '|' {
			return g.Up(from.X, from.Y, false), NORTH, nil
		}
		if val == '7' {
			return g.Left(from.X, from.Y, false), WEST, nil
		}
		if val == 'F' {
			return g.Right(from.X, from.Y, false), EAST, nil
		}
		return grid.Point{0, 0}, NORTH, fmt.Errorf("coming from south, found: %s", string(val))
	} else if prev == WEST {
		if val == '-' {
			return g.Right(from.X, from.Y, false), EAST, nil
		}
		if val == '7' {
			return g.Down(from.X, from.Y, false), SOUTH, nil
		}
		if val == 'J' {
			return g.Up(from.X, from.Y, false), NORTH, nil
		}
		return grid.Point{0, 0}, NORTH, fmt.Errorf("coming from west, found: %s", string(val))
	} else { // prev == EAST
		if val == '-' {
			return g.Left(from.X, from.Y, false), WEST, nil
		}
		if val == 'L' {
			return g.Up(from.X, from.Y, false), NORTH, nil
		}
		if val == 'F' {
			return g.Down(from.X, from.Y, false), SOUTH, nil
		}
		return grid.Point{0, 0}, NORTH, fmt.Errorf("coming from east, found: %s", string(val))
	}
}

type Neighbour struct {
	pos grid.Point
	dir Dir
}

func connectingNeighbours(g *grid.Grid[rune], pos grid.Point) []Neighbour {
	res := []Neighbour{}
	if pos.X < g.MaxX() {
		val := *g.Get(pos.X+1, pos.Y)
		if val == '|' || val == 'L' || val == 'J' {
			res = append(res, Neighbour{grid.Point{pos.X + 1, pos.Y}, SOUTH})
		}
	}
	if pos.X > 0 {
		val := *g.Get(pos.X-1, pos.Y)
		if val == '|' || val == '7' || val == 'F' {
			res = append(res, Neighbour{grid.Point{pos.X - 1, pos.Y}, NORTH})
		}
	}

	if pos.Y < g.MaxY() {
		val := *g.Get(pos.X, pos.Y+1)
		if val == '-' || val == '7' || val == 'J' {
			res = append(res, Neighbour{grid.Point{pos.X, pos.Y + 1}, EAST})
		}
	}

	if pos.Y > 0 {
		val := *g.Get(pos.X, pos.Y-1)
		if val == '-' || val == 'L' || val == 'F' {
			res = append(res, Neighbour{grid.Point{pos.X, pos.Y - 1}, WEST})
		}
	}

	return res
}

func Solution() int {
	parsed, err := parse.GetLinesAs[[]rune]("day10/input.txt",
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

	var sPos grid.Point
	g, _ := grid.NewGrid[rune, rune](parsed, func(c rune, x, y int) (rune, error) {
		if c == 'S' {
			sPos = grid.Point{x, y}
		}
		return c, nil
	})

	neighs := connectingNeighbours(&g, sPos)
	if len(neighs) != 2 {
		fmt.Printf("Didn't find two neighbours, found %d\n", len(neighs))
		return 1
	}

	loop := utils.Set[grid.Point]{}
	loop.Add(sPos)
	p1 := neighs[0].pos
	d1 := neighs[0].dir
	p2 := neighs[1].pos
	d2 := neighs[1].dir
	g.Set(sPos.X, sPos.Y, closeLoop(d1, d2))
	loop.Add(p1)
	loop.Add(p2)
	res := 1
	for p1 != p2 {
		p1, d1, err = next(&g, p1, reflect(d1))
		if err != nil {
			fmt.Printf("Error walking the loop: %v\n", err)
		}
		p2, d2, err = next(&g, p2, reflect(d2))
		if err != nil {
			fmt.Printf("Error walking the loop: %v\n", err)
		}
		loop.Add(p1)
		loop.Add(p2)
		res += 1
	}

	fmt.Printf("Part 1: %d\n", res)

	// remove the unnecessary pipes
	g2 := grid.Copy[rune](g)
	for x := 0; x <= g.MaxX(); x++ {
		for y := 0; y <= g.MaxY(); y++ {
			if loop.Contains(grid.Point{x, y}) {
				continue
			}
			g2.Set(x, y, '.')
		}
	}

	count := 0
	for x := 0; x <= g2.MaxX(); x++ {
		for y := 0; y <= g2.MaxY(); y++ {
			if loop.Contains(grid.Point{x, y}) {
				continue
			}
			// cast ray to the right if we find an odd number of pipes then we
			// are inside the loop
			inside := false
			for yt := y + 1; yt <= g2.MaxY(); yt++ {
				val := *g2.Get(x, yt)
				if val == '|' || val == 'L' || val == 'J' {
					inside = !inside
				}
			}
			if inside {
				count += 1
			}
		}
	}
	fmt.Printf("Part 2: %d\n", count)

	return 0
}
