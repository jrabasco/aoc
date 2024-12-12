package day12

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
)

type Region struct {
	area      int
	sides     int
	perimeter int
}

type Tile struct {
	tpe         rune
	x           int
	y           int
	region      *Region
	fenceLeft   bool
	fenceRight  bool
	fenceTop    bool
	fenceBottom bool
}

type Map struct {
	g         grid.Grid[Tile]
	regions   []*Region
	processed bool
}

func fillRegion(g *grid.Grid[Tile], x, y int) *Region {
	t0 := g.Get(x, y)
	r := Region{1, 0, 0}
	t0.region = &r

	q := utils.NewQueue[*Tile]()
	q.Enqueue(t0)

	for !q.Empty() {
		t, _ := q.Dequeue()

		for _, p := range g.Neighbours(t.x, t.y) {
			tn := g.GetAt(p)
			if tn.region != nil || tn.tpe != t0.tpe {
				continue
			}
			r.area += 1
			tn.region = &r
			q.Enqueue(tn)
		}
	}
	return &r
}

func process(m *Map) {
	for x := 0; x <= m.g.MaxX(); x++ {
		for y := 0; y <= m.g.MaxY(); y++ {
			if m.g.Get(x, y).region != nil {
				continue
			}
			m.regions = append(m.regions, fillRegion(&m.g, x, y))
		}
	}

	// each row, cast a ray and update perimeters
	for x := 0; x <= m.g.MaxX(); x++ {
		cur := m.g.Get(x, 0)
		// this is for the fence on the left
		cur.region.perimeter++
		cur.fenceLeft = true

		// is this a new side? answer this by checking the previous
		// row and see if it was in the same region.
		if x == 0 || m.g.Get(x-1, 0).tpe != cur.tpe {
			cur.region.sides++
		}

		for y := 0; y <= m.g.MaxY(); y++ {
			n := m.g.Get(x, y)
			if cur.tpe == n.tpe {
				cur = n
				continue
			}
			cur.region.perimeter++
			cur.fenceRight = true
			n.region.perimeter++
			n.fenceLeft = true
			// is this a new side? answer this by checking the previous
			// row and see if it was in the same region and had the same fence
			if x == 0 {
				cur.region.sides++
				n.region.sides++
			} else {
				prevC := m.g.Get(x-1, cur.y)
				prevN := m.g.Get(x-1, n.y)

				if cur.tpe != prevC.tpe || !prevC.fenceRight {
					cur.region.sides++
				}
				if n.tpe != prevN.tpe || !prevN.fenceLeft {
					n.region.sides++
				}
			}
			cur = n
		}

		// this is for the fence on the right
		cur.region.perimeter++
		cur.fenceRight = true
		// is this a new side? answer this by checking the previous
		// row and see if it was in the same region.
		if x == 0 || m.g.Get(x-1, m.g.MaxY()).tpe != cur.tpe {
			cur.region.sides++
		}
	}

	// each column, cast a ray and update perimeters
	for y := 0; y <= m.g.MaxY(); y++ {
		cur := m.g.Get(0, y)
		// this is for the fence on the top
		cur.region.perimeter++
		cur.fenceTop = true
		// is this a new side? answer this by checking the previous
		// column and see if it was in the same region.
		if y == 0 || m.g.Get(0, y-1).tpe != cur.tpe {
			cur.region.sides++
		}

		for x := 0; x <= m.g.MaxX(); x++ {
			n := m.g.Get(x, y)
			if cur.tpe == n.tpe {
				cur = n
				continue
			}
			cur.region.perimeter++
			cur.fenceBottom = true
			n.region.perimeter++
			n.fenceTop = true
			// is this a new side? answer this by checking the previous
			// row and see if it was in the same region and had the same fence
			if y == 0 {
				cur.region.sides++
				n.region.sides++
			} else {
				prevC := m.g.Get(cur.x, y-1)
				prevN := m.g.Get(n.x, y-1)

				if cur.tpe != prevC.tpe || !prevC.fenceBottom {
					cur.region.sides++
				}
				if n.tpe != prevN.tpe || !prevN.fenceTop {
					n.region.sides++
				}
			}
			cur = n
		}

		// this is for the fence on the bottom
		cur.region.perimeter++
		cur.fenceBottom = true
		// is this a new side? answer this by checking the previous
		// column and see if it was in the same region.
		if y == 0 || m.g.Get(m.g.MaxX(), y-1).tpe != cur.tpe {
			cur.region.sides++
		}
	}
	m.processed = true
}

func p1(m *Map) int {
	if !m.processed {
		process(m)
	}
	cost := 0
	for _, r := range m.regions {
		cost += r.area * r.perimeter
	}
	return cost
}

func p2(m *Map) int {
	if !m.processed {
		process(m)
	}
	cost := 0
	for _, r := range m.regions {
		cost += r.area * r.sides
	}
	return cost
}

func parseMap(lines [][]rune) Map {
	g, _ := grid.NewGridAs[Tile, rune](lines, func(r rune, i, j int) (Tile, error) {
		return Tile{r, i, j, nil, false, false, false, false}, nil
	})
	return Map{g, []*Region{}, false}
}

func Solution() int {
	parsed, err := parse.GetLinesAsRunes("day12/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	g := parseMap(parsed)
	fmt.Printf("Part 1: %d\n", p1(&g))
	fmt.Printf("Part 2: %d\n", p2(&g))
	return 0
}
