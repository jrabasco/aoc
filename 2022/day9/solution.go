package day9

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type Pos struct {
	x int
	y int
}

func NewPos() Pos {
	return Pos{0, 0}
}

func (p Pos) Touches(o Pos) bool {
	return abs(p.x-o.x) <= 1 && abs(p.y-o.y) <= 1
}

func (p *Pos) MoveR() {
	p.x++
}

func (p *Pos) MoveL() {
	p.x--
}

func (p *Pos) MoveU() {
	p.y++
}

func (p *Pos) MoveD() {
	p.y--
}

func (p *Pos) CatchUp(o Pos) {
	// if it's only 1 away in both direction, no move
	if p.Touches(o) {
		return
	}
	// on the same row
	if o.y == p.y {
		if o.x > p.x+1 {
			p.MoveR()
		}

		if o.x < p.x-1 {
			p.MoveL()
		}
	}

	// on the same column
	if p.x == o.x {
		if o.y > p.y+1 {
			p.MoveU()
		}

		if o.y < p.y-1 {
			p.MoveD()
		}
	}

	// diagonal
	if p.x != o.x && p.y != o.y {
		if o.x > p.x {
			p.MoveR()
		}

		if o.x < p.x {
			p.MoveL()
		}

		if o.y > p.y {
			p.MoveU()
		}

		if o.y < p.y {
			p.MoveD()
		}
	}
}

func parseMove(line string) (func(p *Pos), int, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return func(p *Pos) {}, 0, fmt.Errorf("invalid move line: '%s'", line)
	}

	amount, err := strconv.Atoi(parts[1])

	if err != nil {
		return func(p *Pos) {}, 0, fmt.Errorf("error while parsing '%s': %v", parts[1], err)
	}

	switch parts[0] {
	case "R":
		return (*Pos).MoveR, amount, nil
	case "L":
		return (*Pos).MoveL, amount, nil
	case "U":
		return (*Pos).MoveU, amount, nil
	case "D":
		return (*Pos).MoveD, amount, nil
	default:
		return func(p *Pos) {}, 0, fmt.Errorf("invalid direction: %s", parts[0])
	}
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day9/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	H := NewPos()
	tail := []Pos{}

	tailLen := 1
	if part == "2" {
		tailLen = 9
	}

	for i := 0; i < tailLen; i++ {
		tail = append(tail, NewPos())
	}

	T := &tail[tailLen-1]
	visited := utils.Set[Pos]{}
	visited.Add(*T)

	for _, line := range parsed {
		m, a, err := parseMove(line)
		if err != nil {
			fmt.Printf("Could not parse move: %v\n", err)
			return 1
		}
		for i := 0; i < a; i++ {
			m(&H)
			tail[0].CatchUp(H)
			for j := 1; j < tailLen; j++ {
				tail[j].CatchUp(tail[j-1])
			}
			visited.Add(*T)
		}
	}

	fmt.Printf("Part %s: %v\n", part, len(visited))
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
