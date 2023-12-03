package day3

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strconv"
	"unicode"
)

type PosPart struct {
	x      int
	y      int
	symbol rune
}

type Parts map[PosPart][]int

func isSymbol(c rune) bool {
	return c != '.' && !unicode.IsDigit(c)
}

func neighSymbols(g *grid.Grid[rune], i, j int) []PosPart {
	res := []PosPart{}
	for _, p := range g.DNeighbours(i, j) {
		c := g.Get(p.X, p.Y)
		if isSymbol(*c) {
			res = append(res, PosPart{p.X, p.Y, *c})
		}
	}
	return res
}

func getPart(g *grid.Grid[rune], i, j int, parts *Parts) (int, int, bool) {
	c := g.Get(i, j)
	if !unicode.IsDigit(*c) {
		return 0, 0, false
	}
	delta := -1
	isPart := false
	nbStr := ""
	allNeighs := utils.Set[PosPart]{}
	for j <= g.MaxY() {
		c = g.Get(i, j)
		if !unicode.IsDigit(*c) {
			break
		}
		nbStr += string(*c)
		neighs := neighSymbols(g, i, j)
		if len(neighs) != 0 {
			isPart = true
		}
		for _, neigh := range neighs {
			allNeighs.Add(neigh)
		}
		j++
		delta++
	}
	nb, _ := strconv.Atoi(nbStr)
	for neigh, _ := range allNeighs {
		p, ok := (*parts)[neigh]
		if !ok {
			(*parts)[neigh] = []int{nb}
		} else {
			(*parts)[neigh] = append(p, nb)
		}
	}
	return nb, delta, isPart
}

func Solution(string) int {
	parsed, err := parse.GetLinesAs[[]rune]("day3/input.txt",
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
	g, _ := grid.NewGrid[rune, rune](parsed, func(c rune, x, y int) (rune, error) { return c, nil })
	res1 := 0
	parts := Parts{}
	for i := 0; i <= g.MaxX(); i++ {
		for j := 0; j <= g.MaxY(); j++ {
			nb, delta, isPart := getPart(&g, i, j, &parts)
			if isPart {
				res1 += nb
			}
			j += delta
		}
	}
	res2 := 0
	for pospart, nbs := range parts {
		if pospart.symbol == '*' && len(nbs) == 2 {
			res2 += nbs[0] * nbs[1]
		}
	}
	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
