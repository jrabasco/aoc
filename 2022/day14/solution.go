package day14

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/grid"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
	"strings"
)

type Cave interface {
	Add(x, y int, r rune)
	String() string
	MaxX() int
	MaxY() int
	MinX() int
	MinY() int
	Get(x, y int) rune
}

type BaseCave struct {
	grid.SparseGrid[rune]
}

func (c *BaseCave) String() string {
	lines := []string{}
	for y := c.MinY(); y <= c.MaxY(); y++ {
		lines = append(lines, "")
		for x := c.MinX(); x <= c.MaxX(); x++ {
			lines[y-c.MinY()] += string(c.Get(x, y))
		}
	}
	return strings.Join(lines, "\n")
}

type CaveWithFloor struct {
	bc    *BaseCave
	floor int
}

func NewCaveWithFloor(bc *BaseCave) CaveWithFloor {
	return CaveWithFloor{bc, bc.MaxY() + 2}
}

func (c CaveWithFloor) Get(x, y int) rune {
	if y == c.floor {
		return '#'
	}

	return c.bc.Get(x, y)
}

func (c CaveWithFloor) MaxX() int {
	return c.bc.MaxX()
}

func (c CaveWithFloor) MaxY() int {
	return c.floor
}

func (c CaveWithFloor) MinX() int {
	return c.bc.MinX()
}

func (c CaveWithFloor) MinY() int {
	return c.bc.MinY()
}

func (c *CaveWithFloor) Add(x, y int, r rune) {
	c.bc.Add(x, y, r)
}

func (c CaveWithFloor) String() string {
	res := c.bc.String()
	res += "\n"
	for x := c.bc.MinX(); x <= c.bc.MaxX(); x++ {
		res += "#"
	}
	return res
}

func parseCoord(coordS string) (int, int, error) {
	coordPS := strings.Split(coordS, ",")
	if len(coordPS) == 0 {
		return 0, 0, fmt.Errorf("invalid coordinate: %s", coordS)
	}

	x, err := strconv.Atoi(coordPS[0])
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse x: %v", err)
	}

	y, err := strconv.Atoi(coordPS[1])
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse y: %v", err)
	}

	return x, y, nil
}

func populateGrid(g *BaseCave, lines []string) error {
	for _, line := range lines {
		pts := strings.Split(line, " -> ")
		if len(pts) == 0 {
			continue
		}

		startS := pts[0]
		rest := pts[1:]

		x, y, err := parseCoord(startS)
		if err != nil {
			return err
		}

		g.Add(x, y, '#')

		for len(rest) > 0 {
			next := rest[0]
			rest = rest[1:]

			nextX, nextY, err := parseCoord(next)
			if err != nil {
				return err
			}

			for x != nextX || y != nextY {
				if x > nextX {
					x--
				}
				if x < nextX {
					x++
				}

				if y > nextY {
					y--
				}

				if y < nextY {
					y++
				}

				g.Add(x, y, '#')
			}

		}

	}
	return nil
}

func dropSand(c Cave) int {
	res := 0
	cont := true
	for cont {
		cont = false
		sandX := 500
		sandY := 0

		falling := true
		for falling {
			if sandY > c.MaxY() {
				break
			}
			falling = false
			if c.Get(sandX, sandY+1) == '.' {
				sandY += 1
				falling = true
				continue
			}

			if c.Get(sandX-1, sandY+1) == '.' {
				sandX -= 1
				sandY += 1
				falling = true
				continue
			}

			if c.Get(sandX+1, sandY+1) == '.' {
				sandX += 1
				sandY += 1
				falling = true
				continue
			}
		}

		if !falling {
			c.Add(sandX, sandY, 'o')
			res += 1
		}

		// check if it didn't settle on the starting point
		// and that it's not falling below the floor
		if (sandX != 500 || sandY != 0) && !falling {
			cont = true
		}
	}
	return res
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day14/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	g := BaseCave{grid.NewSparseGrid[rune]('.')}
	populateGrid(&g, parsed)

	numSand := 0
	if part == "1" {
		numSand = dropSand(&g)
	} else {
		c := NewCaveWithFloor(&g)
		numSand = dropSand(&c)
	}
	fmt.Printf("Part %s: %v\n", part, numSand)
	return 0
}

func Solution(part string) int {
	if part != "1" && part != "2" {
		p1 := solvePart("1")
		if p1 != 0 {
			return p1
		}
		return solvePart("2")
	} else {
		return solvePart(part)
	}
}
