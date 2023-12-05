package grid

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Grid[T any] struct {
	grid [][]T
	h    int
	w    int
	maxX int
	maxY int
}

func NewGrid[T any, S any](lines [][]S, conv func(S, int, int) (T, error)) (Grid[T], error) {
	var empty Grid[T]
	var grid [][]T
	h := 0
	w := 0
	maxX := 0
	maxY := 0
	for x, line := range lines {
		if x > maxX {
			maxX = x
			h = x + 1
		}
		grid = append(grid, []T{})
		for y, item := range line {
			if y > maxY {
				maxY = y
				w = y + 1
			}
			nitem, err := conv(item, x, y)
			if err != nil {
				return empty, err
			}
			grid[x] = append(grid[x], nitem)
		}
	}

	return Grid[T]{
		grid: grid,
		h:    h,
		w:    w,
		maxX: maxX,
		maxY: maxY,
	}, nil
}

func (g *Grid[T]) H() int {
	return g.h
}

func (g *Grid[T]) W() int {
	return g.w
}

func (g *Grid[T]) MaxX() int {
	return g.maxX
}

func (g *Grid[T]) MaxY() int {
	return g.maxY
}

func (g *Grid[T]) Get(x, y int) *T {
	return &g.grid[x][y]
}

func (g *Grid[T]) Neighbours(x int, y int) []Point {
	res := []Point{}
	if g.h == 0 || g.w == 0 {
		return res
	}

	if x < g.maxX {
		res = append(res, Point{x + 1, y})
	}

	if x > 0 {
		res = append(res, Point{x - 1, y})
	}

	if y < g.maxY {
		res = append(res, Point{x, y + 1})
	}

	if y > 0 {
		res = append(res, Point{x, y - 1})
	}
	return res
}

func (g *Grid[T]) DNeighbours(x int, y int) []Point {
	res := []Point{}
	if g.h == 0 || g.w == 0 {
		return res
	}

	if x < g.maxX {
		res = append(res, Point{x + 1, y})
		if y < g.maxY {
			res = append(res, Point{x + 1, y + 1})
		}

		if y > 0 {
			res = append(res, Point{x + 1, y - 1})
		}
	}

	if x > 0 {
		res = append(res, Point{x - 1, y})
		if y < g.maxY {
			res = append(res, Point{x - 1, y + 1})
		}

		if y > 0 {
			res = append(res, Point{x - 1, y - 1})
		}
	}

	if y < g.maxY {
		res = append(res, Point{x, y + 1})
	}

	if y > 0 {
		res = append(res, Point{x, y - 1})
	}
	return res
}

func (g *Grid[T]) Down(x, y int, wrap bool) Point {
	if x >= g.maxX {
		if wrap {
			return Point{0, y}
		}
		panic("Cannot go down")
	}

	return Point{x + 1, y}
}

func (g *Grid[T]) Up(x, y int, wrap bool) Point {
	if x == 0 {
		if wrap {
			return Point{g.maxX, y}
		}
		panic("Cannot go up")
	}

	return Point{x - 1, y}
}

func (g *Grid[T]) Right(x, y int, wrap bool) Point {
	if y >= g.maxY {
		if wrap {
			return Point{x, 0}
		}
		panic("Cannot go right")
	}

	return Point{x, y + 1}
}

func (g *Grid[T]) Left(x, y int, wrap bool) Point {
	if y == 0 {
		if wrap {
			return Point{x, g.maxY}
		}
		panic("Cannot go left")
	}

	return Point{x, y - 1}
}

func (g Grid[T]) Rows() [][]*T {
	res := [][]*T{}
	for x := 0; x <= g.maxX; x++ {
		res = append(res, g.Row(x))
	}
	return res
}

func (g Grid[T]) RowsCopy() [][]T {
	return g.grid
}

func (g Grid[T]) Row(x int) []*T {
	res := []*T{}
	for y := 0; y <= g.maxY; y++ {
		res = append(res, &g.grid[x][y])
	}
	return res
}

func (g Grid[T]) Column(y int) []*T {
	res := []*T{}
	for x := 0; x <= g.maxX; x++ {
		res = append(res, &g.grid[x][y])
	}
	return res
}

func (g Grid[T]) Columns() [][]*T {
	res := [][]*T{}
	for y := 0; y <= g.maxY; y++ {
		res = append(res, g.Column(y))
	}
	return res
}

func Convert[T any, S any](g Grid[T], conv func(T, int, int) (S, error)) (Grid[S], error) {
	return NewGrid[S, T](g.RowsCopy(), conv)
}

func (g Grid[T]) String() string {
	var lines []string
	// no error can be returned since the conv function doesn't error
	gridStr, _ := Convert[T, string](g, func(elm T, x, y int) (string, error) { return fmt.Sprintf("%v", elm), nil })
	for _, row := range gridStr.RowsCopy() {
		lines = append(lines, strings.Join(row, ""))
	}
	return strings.Join(lines, "\n")
}

func BasicTest() int {
	lines := [][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	}
	g, _ := NewGrid[int, string](lines, func(s string, x, y int) (int, error) {
		e, err := strconv.Atoi(s)
		return e, err
	})
	// no error can be returned since the conv function doesn't error
	t, _ := NewGrid[int, *int](g.Columns(), func(e *int, x, y int) (int, error) { return *e, nil })
	fmt.Println(g)
	fmt.Println("-------")
	fmt.Println(t)
	return 0
}
