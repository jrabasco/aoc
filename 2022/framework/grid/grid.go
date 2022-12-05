package grid

import (
    "fmt"
    "strings"
)

type Point struct {
    x int
    y int
}

type Grid[T any] struct{
    grid [][]T
    h int
    w int
    maxX int
    maxY int
}


func NewGrid[T any, S any](lines [][]S, conv func(S)T) Grid[T] {
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
            grid[x] = append(grid[x], conv(item))
        }
    }

    return Grid[T]{
        grid: grid,
        h: h,
        w: w,
        maxX: maxX,
        maxY: maxY,
    }
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

func (g *Grid[T]) Neighbours(x int, y int) []Point {
    res := []Point{}
    if g.h == 0 || g.w == 0 {
        return res
    }

    if x < g.maxX {
        res = append(res, Point{x+1, y})
    }

    if x > 0 {
        res = append(res, Point{x-1, y})
    }

    if y < g.maxY {
        res = append(res, Point{x, y+1})
    }

    if y > 0 {
        res = append(res, Point{x, y-1})
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

    return Point{x+1, y}
}

func (g *Grid[T]) Up(x, y int, wrap bool) Point {
    if x == 0 {
        if wrap {
            return Point{g.maxX, y}
        }
        panic("Cannot go up")
    }

    return Point{x-1, y}
}

func (g *Grid[T]) Right(x, y int, wrap bool) Point {
    if y >= g.maxY {
        if wrap {
            return Point{x, 0}
        }
        panic("Cannot go right")
    }

    return Point{x, y+1}
}

func (g *Grid[T]) Left(x, y int, wrap bool) Point {
    if y == 0 {
        if wrap {
            return Point{x, g.maxY}
        }
        panic("Cannot go left")
    }

    return Point{x, y-1}
}

func (g Grid[T]) Rows() [][]T {
    return g.grid
}

func (g Grid[T]) Row(x int) []T {
    return g.grid[x]
}

func (g Grid[T]) Column(y int) []T {
    res := []T{}
    for x:=0; x <= g.maxX; x++ {
        res = append(res, g.grid[x][y])
    }
    return res
}

func (g Grid[T]) Columns() [][]T {
    res := [][]T{}
    for y := 0; y <= g.maxY; y++ {
        res = append(res, g.Column(y))
    }
    return res
}

func RowAs[T any, S any](g Grid[T], x int, conv func(T)S) []S {
    res := []S{}
    for _, elm := range g.Row(x) {
        res = append(res, conv(elm))
    }
    return res
}

func RowsAs[T any, S any](g Grid[T], conv func(T)S) [][]S {
    res := [][]S{}
    for x := 0; x <= g.maxX; x++ {
        res = append(res, RowAs[T,S](g, x, conv))
    }
    return res
}

func ColumnAs[T any, S any](g Grid[T], y int, conv func(T)S) []S {
    res := []S{}
    for x:=0; x <= g.maxX; x++ {
        res = append(res, conv(g.grid[x][y]))
    }
    return res
}

func ColumnsAs[T any, S any](g Grid[T], conv func(T)S) [][]S {
    res := [][]S{}
    for y := 0; y <= g.maxY; y++ {
        res = append(res, ColumnAs[T,S](g, y, conv))
    }
    return res
}

func (g Grid[T]) String() string{
    var lines []string
    rowsStr := RowsAs[T, string](g, func(elm T) string{return fmt.Sprintf("%v", elm)})
    for _, line := range rowsStr {
        lines = append(lines, strings.Join(line, ""))
    }
    return strings.Join(lines, "\n")
}
