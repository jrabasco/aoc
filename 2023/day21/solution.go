package day21

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
)

func possibleArrivals(g *grid.Grid[string], start grid.Point, steps int) utils.Set[grid.Point] {
	res := utils.NewSet[grid.Point]()
	if steps == 0 {
		res.Add(start)
		return res
	}
	prevs := possibleArrivals(g, start, steps-1)
	for prev := range prevs {
		neighs := []grid.Point{
			grid.Point{prev.X, prev.Y + 1},
			grid.Point{prev.X, prev.Y - 1},
			grid.Point{prev.X + 1, prev.Y},
			grid.Point{prev.X - 1, prev.Y},
		}
		for _, p := range neighs {
			x := p.X % g.H()
			if x < 0 {
				x += g.H()
			}
			y := p.Y % g.W()
			if y < 0 {
				y += g.W()
			}
			if *g.Get(x, y) != "#" {
				res.Add(p)
			}
		}
	}
	return res
}

// Lagrange interpolation of degree 2
type Lagrange2 struct {
	xf [3]float64
	yf [3]float64
}

func FromInts(xi [3]int, yi [3]int) Lagrange2 {
	xf := [3]float64{0, 0, 0}
	yf := [3]float64{0, 0, 0}
	for i := 0; i < 3; i++ {
		xf[i] = float64(xi[i])
		yf[i] = float64(yi[i])
	}
	return Lagrange2{xf, yf}
}

func (l Lagrange2) P(x float64) float64 {
	var res float64 = 0
	for j := 0; j < 3; j++ {
		res += l.Pj(j, x)
	}
	return res
}

func (l Lagrange2) Pj(j int, x float64) float64 {
	var res float64 = 1
	for k := 0; k < 3; k++ {
		if k == j {
			continue
		}
		res *= (x - l.xf[k]) / (l.xf[j] - l.xf[k])
	}
	return l.yf[j] * res
}

func Solution() int {
	parsed, err := parse.GetLinesAsRunes("day21/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	spos := grid.Point{}
	g, _ := grid.NewGridAs[string, rune](parsed, func(elm rune, i, j int) (string, error) {
		if elm == 'S' {
			spos = grid.Point{i, j}
		}
		return string(elm), nil
	})
	arr := possibleArrivals(&g, spos, 64)
	fmt.Printf("Part 1: %d\n", len(arr))
	// Apparently because of the fact that the edges and the straight lines
	// from S are . it's going to be a quadratic every g.H() steps.
	x := 26501365
	x0 := x % g.H()
	x1 := x0 + g.H()
	x2 := x1 + g.H()
	y0 := 3725  // len(possibleArrivals(&g, spos, x0))
	y1 := 32896 // len(possibleArrivals(&g, spos, x1))
	y2 := 91055 // len(possibleArrivals(&g, spos, x2))
	l := FromInts([3]int{x0, x1, x2}, [3]int{y0, y1, y2})
	fmt.Printf("Part 2: %.0f\n", l.P(float64(x)))
	return 0
}
