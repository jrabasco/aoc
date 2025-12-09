package day9

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/grid"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
	"strconv"
	"strings"
)

func area(p1, p2 grid.Point) int {
	return (utils.AbsDiff(p1.X, p2.X) + 1) * (utils.AbsDiff(p1.Y, p2.Y) + 1)
}

func p1(points []grid.Point) int {
	maxArea := 0
	for i := range points {
		for j := i; j < len(points); j++ {
			a := area(points[i], points[j])
			if a > maxArea {
				maxArea = a
			}
		}
	}
	return maxArea
}

func doesIntersect(a1, a2, b1, b2 grid.Point) bool {
	// both are vertical
	if a1.X == a2.X && b1.X == b2.X {
		return false
	}

	// both horizontal
	if a1.Y == a2.Y && b1.Y == b2.Y {
		return false
	}

	// a is vertical, b is horizontal
	if a1.X == a2.X {
		// fully on the left
		if b1.X < a1.X && b2.X < a1.X {
			return false
		}
		// fully on the right
		if b1.X > a1.X && b2.X > a1.X {
			return false
		}

		// here they are on both sides of a
		// check if a crosses the Y coord of b
		return (a1.Y < b1.Y && a2.Y > b1.Y) || (a1.Y > b1.Y && a2.Y < b1.Y)
	}

	// a is horizontal b is vertical
	if a1.Y == a2.Y {
		// fully on top
		if b1.Y < a1.Y && b2.Y < a1.Y {
			return false
		}
		// fully on the bottom
		if b1.Y > a1.Y && b2.Y > a1.Y {
			return false
		}

		// here they are on both sides of a
		// check if a crosses the X coord of b
		return (a1.X < b1.X && a2.X > b1.X) || (a1.X > b1.X && a2.X < b1.X)
	}

	return false
}

func p2(points []grid.Point) int {
	maxArea := 0
	for i := range points {
		for j := i; j < len(points); j++ {
			c1 := points[i]
			c3 := points[j]
			c2 := grid.Point{c1.X, c3.Y}
			c4 := grid.Point{c3.X, c1.Y}
			a := area(c1, c3)
			if a <= maxArea {
				continue
			}
			conflict := false
			for k := 0; k < len(points); k++ {
				n := (k + 1) % len(points)
				if doesIntersect(c1, c2, points[k], points[n]) {
					conflict = true
					break
				}
				if doesIntersect(c2, c3, points[k], points[n]) {
					conflict = true
					break
				}
				if doesIntersect(c3, c4, points[k], points[n]) {
					conflict = true
					break
				}
				if doesIntersect(c4, c1, points[k], points[n]) {
					conflict = true
					break
				}
			}
			if conflict {
				continue
			}
			maxArea = a
		}
	}
	return maxArea
}

func parsePoint(line string) (grid.Point, error) {
	parts := strings.Split(line, ",")
	p := grid.Point{0, 0}
	if len(parts) != 2 {
		return p, fmt.Errorf("invalid point line: %s", line)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return p, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return p, err
	}
	p.X = x
	p.Y = y
	return p, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[grid.Point]("day9/input.txt", parsePoint)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
