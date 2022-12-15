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

func (g SparseGrid[T]) Get(x, y int) T {
	if r, ok := g.grid[Point{x, y}]; ok {
		return r
	}
	return g.empty
}

func (g *SparseGrid[T]) Add(x, y int, r T) {
	if x > g.maxX {
		g.maxX = x
	}

	if x < g.minX {
		g.minX = x
	}

	if y > g.maxY {
		g.maxY = y
	}

	if y < g.minY {
		g.minY = y
	}

	g.grid[Point{x, y}] = r
}

func (g *SparseGrid[T]) MaybeAdd(x, y int, r T) {
	if g.Get(x, y) == g.empty {
		g.Add(x, y, r)
	}
}
