package day12

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/grid"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Square struct {
	val     rune
	x       int
	y       int
	isGoal  bool
	visited bool
	prev    *Square
}

func NewSquare(r rune, x, y int, isGoal bool) Square {
	return Square{r, x, y, isGoal, false, nil}
}

func backTrace(square *Square) []*Square {
	cur := square
	res := []*Square{cur}
	for cur.prev != nil {
		res = append(res, cur.prev)
		cur = cur.prev
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func findPath(g *grid.Grid[Square], x int, y int, isGoal func(*Square) bool, canGo func(rune, rune) bool) []*Square {
	// queue of paths
	queue := utils.NewQueue[grid.Point]()
	queue.Enqueue(grid.Point{x, y})
	g.Get(x, y).visited = true

	for !queue.Empty() {
		// ignore error because of loop condition
		curP, _ := queue.Dequeue()
		cur := g.Get(curP.X, curP.Y)

		if isGoal(cur) {
			return backTrace(cur)
		}

		for _, neighP := range g.Neighbours(curP.X, curP.Y) {
			neigh := g.Get(neighP.X, neighP.Y)
			if !canGo(cur.val, neigh.val) {
				continue
			}

			if neigh.visited {
				continue
			}

			neigh.visited = true
			neigh.prev = cur
			queue.Enqueue(neighP)
		}
	}
	return nil
}

func solvePart(part string) int {
	parsed, err := parse.GetLinesAs[[]rune]("day12/input.txt",
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

	startX := -1
	startY := -1

	endX := -1
	endY := -1
	g, err := grid.NewGrid[Square, rune](parsed, func(r rune, x, y int) (Square, error) {
		el := r
		if r == 'S' {
			el = 'a'
			startX = x
			startY = y
		}

		isGoal := false
		if r == 'E' {
			el = 'z'
			isGoal = true
			endX = x
			endY = y
		}
		return NewSquare(el, x, y, isGoal), nil
	})
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	path := []*Square{}

	if part == "1" {
		path = findPath(&g, startX, startY, func(s *Square) bool { return s.isGoal }, func(v1, v2 rune) bool { return v2-v1 <= 1 })
	} else {
		path = findPath(&g, endX, endY, func(s *Square) bool { return s.val == 'a' }, func(v1, v2 rune) bool { return v1-v2 <= 1 })
	}

	// path contains start so need -1
	fmt.Printf("Part %s: %v\n", part, len(path)-1)
	return 0
}

func getColor(r rune, maxCol color.RGBA) color.RGBA {
	max := 'z' - 'a'
	v := r - 'a'
	ratio := float64(v) / float64(max)
	R := uint8(float64(maxCol.R) * ratio)
	G := uint8(float64(maxCol.G) * ratio)
	B := uint8(float64(maxCol.B) * ratio)
	return color.RGBA{R, G, B, 0xff}
}

var inflate = 32

func putSquare(x, y int, img *image.RGBA, col color.RGBA) {
	for i := inflate * x; i < inflate*(x+1); i++ {
		for j := inflate * y; j < inflate*(y+1); j++ {
			img.Set(j, i, col)
		}
	}
}

func getCol(x, y int, img *image.RGBA) color.RGBA {
	return img.RGBAAt(inflate*y, inflate*x)
}

func createHeatMap(parsed [][]rune, startCol, endCol, maxCol color.RGBA) (*image.RGBA, error) {
	startX := -1
	startY := -1
	g, err := grid.NewGrid[Square, rune](parsed, func(r rune, x, y int) (Square, error) {
		el := r
		if r == 'S' {
			el = 'a'
			startX = x
			startY = y
		}

		isGoal := false
		if r == 'E' {
			el = 'z'
			isGoal = true
		}
		return NewSquare(el, x, y, isGoal), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse input : %v", err)
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{inflate * g.W(), inflate * g.H()}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x, r := range g.Rows() {
		for y, e := range r {
			if e.isGoal {
				putSquare(x, y, img, endCol)
				continue
			}

			if x == startX && y == startY {
				putSquare(x, y, img, startCol)
				continue
			}
			col := getColor(e.val, maxCol)
			putSquare(x, y, img, col)
		}
	}

	return img, nil
}

func saveHeatMap() int {
	parsed, err := parse.GetLinesAs[[]rune]("day12/input.txt",
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

	img, err := createHeatMap(parsed, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0xff, 0, 0x44, 0xff})
	if err != nil {
		fmt.Printf("Failed heatmap generation: %v\n", err)
		return 1
	}

	f, err := os.Create("output/day12/heatmap.png")
	if err != nil {
		fmt.Printf("Could not create image file: %v\n", err)
		return 1
	}
	png.Encode(f, img)
	return 0
}

func tracePathIMG(g *grid.Grid[Square], x int, y int, img *image.RGBA, name string, isGoal func(*Square) bool, canGo func(rune, rune) bool) error {
	// queue of paths
	queue := utils.NewQueue[grid.Point]()
	queue.Enqueue(grid.Point{x, y})
	g.Get(x, y).visited = true

	i := 0

	for !queue.Empty() {
		// ignore error because of loop condition
		curP, _ := queue.Dequeue()
		cur := g.Get(curP.X, curP.Y)
		prevCol := getCol(cur.x, cur.y, img)
		putSquare(cur.x, cur.y, img, color.RGBA{0, 0xff, 0, 0xff})
		f, err := os.Create(fmt.Sprintf("%s%d.png", name, i))
		if err != nil {
			return fmt.Errorf("could not create image file: %v", err)
		}
		png.Encode(f, img)
		putSquare(curP.X, curP.Y, img, prevCol)

		if isGoal(cur) {
			putSquare(cur.x, cur.y, img, color.RGBA{0, 0xff, 0, 0xff})
			for cur.prev != nil {
				cur = cur.prev
				putSquare(cur.x, cur.y, img, color.RGBA{0, 0xff, 0, 0xff})
			}
			f, err := os.Create(fmt.Sprintf("%s%d.png", name, i))
			if err != nil {
				return fmt.Errorf("could not create image file: %v", err)
			}
			png.Encode(f, img)
			return nil
		}

		for _, neighP := range g.Neighbours(curP.X, curP.Y) {
			neigh := g.Get(neighP.X, neighP.Y)
			if !canGo(cur.val, neigh.val) {
				continue
			}

			if neigh.visited {
				continue
			}

			neigh.visited = true
			neigh.prev = cur
			queue.Enqueue(neighP)
		}
		i++
	}
	return nil
}

func saveTraces() int {
	parsed, err := parse.GetLinesAs[[]rune]("day12/input.txt",
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
	img, err := createHeatMap(parsed, color.RGBA{0, 0, 0, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff})
	if err != nil {
		fmt.Printf("Failed heatmap generation: %v\n", err)
		return 1
	}

	startX := -1
	startY := -1

	endX := -1
	endY := -1
	g, err := grid.NewGrid[Square, rune](parsed, func(r rune, x, y int) (Square, error) {
		el := r
		if r == 'S' {
			el = 'a'
			startX = x
			startY = y
		}

		isGoal := false
		if r == 'E' {
			el = 'z'
			isGoal = true
			endX = x
			endY = y
		}
		return NewSquare(el, x, y, isGoal), nil
	})
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	err = tracePathIMG(&g, startX, startY, img, "output/day12/part1/step", func(s *Square) bool { return s.isGoal }, func(v1, v2 rune) bool { return v2-v1 <= 1 })

	if err != nil {
		fmt.Printf("Failed to print steps: %v\n", err)
		return 1
	}

	// reset heatmap and grid
	img, err = createHeatMap(parsed, color.RGBA{0, 0, 0, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff})
	if err != nil {
		fmt.Printf("Failed heatmap generation: %v\n", err)
		return 1
	}

	for _, r := range g.Rows() {
		for _, e := range r {
			e.visited = false
			e.prev = nil
		}
	}

	err = tracePathIMG(&g, endX, endY, img, "output/day12/part2/step", func(s *Square) bool { return s.val == 'a' }, func(v1, v2 rune) bool { return v1-v2 <= 1 })
	fmt.Println(endX, endY)
	fmt.Println("Done")
	return 0
}

func Solution(part string) int {
	if part != "1" && part != "2" && part != "heat" && part != "trace" {
		p1 := solvePart("1")
		if p1 != 0 {
			return p1
		}
		return solvePart("2")
	} else if part != "heat" && part != "trace" {
		return solvePart(part)
	} else if part == "heat" {
		return saveHeatMap()
	} else if part == "trace" {
		return saveTraces()
	}
	return 7
}
