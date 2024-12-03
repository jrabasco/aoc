package grid

type SparseGrid[T comparable] struct {
	grid  map[Point]T
	empty T
	maxX  int
	maxY  int
	minX  int
	minY  int
}

func NewSparseGrid[T comparable](empty T) SparseGrid[T] {
	return SparseGrid[T]{map[Point]T{},
		empty,
		-9223372036854775808,
		-9223372036854775808,
		9223372036854775807,
		9223372036854775807,
	}
}

func (g SparseGrid[T]) MaxX() int {
	return g.maxX
}

func (g SparseGrid[T]) MaxY() int {
	return g.maxY
}

func (g SparseGrid[T]) MinX() int {
	return g.minX
}

func (g SparseGrid[T]) MinY() int {
	return g.minY
}

func (g SparseGrid[T]) GetP(p Point) T {
	if r, ok := g.grid[p]; ok {
		return r
	}
	return g.empty
}

func (g SparseGrid[T]) Get(x, y int) T {
	return g.GetP(Point{x, y})
}

func (g *SparseGrid[T]) AddP(p Point, r T) {
	if p.X > g.maxX {
		g.maxX = p.X
	}

	if p.X < g.minX {
		g.minX = p.X
	}

	if p.Y > g.maxY {
		g.maxY = p.Y
	}

	if p.Y < g.minY {
		g.minY = p.Y
	}

	g.grid[p] = r
}

func (g *SparseGrid[T]) Add(x, y int, r T) {
	g.AddP(Point{x, y}, r)
}

func (g *SparseGrid[T]) MaybeAddP(p Point, r T) {
	if g.GetP(p) == g.empty {
		g.AddP(p, r)
	}
}

func (g *SparseGrid[T]) MaybeAdd(x, y int, r T) {
	g.MaybeAddP(Point{x, y}, r)
}

func (g SparseGrid[T]) EmptyVal() T {
	return g.empty
}
