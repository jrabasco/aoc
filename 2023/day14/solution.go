package day14

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
)

type Dir struct {
	vertical bool
	dir      int
}

var NORTH = Dir{true, 1}
var WEST = Dir{false, 1}
var SOUTH = Dir{true, -1}
var EAST = Dir{false, -1}

func shift(g *grid.Grid[rune], dir Dir) {
	var slices [][]*rune
	if dir.vertical {
		slices = g.Columns()
	} else {
		slices = g.Rows()
	}

	for _, sl := range slices {
		shiftSlice(sl, dir.dir)
	}
}

func shiftSlice(sl []*rune, dir int) {
	// landing position, i.e. where the next O would fall
	// if it's equal to -1 it means the O doesn't move
	landingPos := -1

	lsl := len(sl)
	start := 0
	end := lsl
	if dir < 0 {
		start = lsl - 1
		end = -1
	}
	for x := start; x != end; x += dir {
		elm := sl[x]
		// found an empty space to land on
		if *elm == '.' && landingPos == -1 {
			landingPos = x
		}

		// now there is a fixed rock between us and wherever the emtpy
		// space was
		if *elm == '#' {
			landingPos = -1
		}

		// found a rounded rock
		if *elm == 'O' {
			// means there are no free space to fall into
			if landingPos == -1 {
				continue
			}

			// fall into the new space and adjust landingPos
			*sl[landingPos] = 'O'
			*sl[x] = '.'
			landingPos = landingPos + dir
		}
	}
}

func score(g *grid.Grid[rune]) int {
	res := 0
	for x, row := range g.Rows() {
		score := g.H() - x
		for _, r := range row {
			if *r == 'O' {
				res += score
			}
		}
	}
	return res
}

var DIRS = []Dir{NORTH, WEST, SOUTH, EAST}

type Map struct {
	g        grid.Grid[rune]
	curShift int
}

func (m *Map) nextShift() {
	g := &(m.g)
	dir := DIRS[m.curShift]
	shift(g, dir)
	m.curShift = (m.curShift + 1) % 4
}

func (m Map) score() int {
	return score(&(m.g))
}

type MapState = utils.Set[grid.Point]

func (m Map) state() MapState {
	res := MapState{}
	for x, row := range m.g.Rows() {
		for y := range row {
			if *(m.g.Get(x, y)) == 'O' {
				res.Add(grid.Point{x, y})
			}
		}
	}
	return res
}

func (m *Map) cycle() int {
	// start a new cycle
	if m.curShift == 0 {
		m.nextShift()
	}

	// finish the cycle
	for m.curShift != 0 {
		m.nextShift()
	}
	return m.score()
}

func Solution() int {
	parsed, err := parse.GetLinesAs[[]rune]("day14/input.txt", func(line string) ([]rune, error) {
		return []rune(line), nil
	})
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	grid := grid.NewGrid[rune](parsed)
	mp := Map{grid, 0}
	mp.nextShift()
	fmt.Printf("Part 1: %d\n", mp.score())

	scores := []int{mp.cycle()}
	states := []MapState{mp.state()}
	found := false
	cycleStart := 0
	cycleEnd := 0
	for !found {
		score := mp.cycle()
		state := mp.state()
		for i, st := range states {
			if state.Equal(st) {
				found = true
				cycleStart = i
				break
			}
		}
		cycleEnd++
		if !found {
			states = append(states, state)
			scores = append(scores, score)
		}
	}
	cycleLength := cycleEnd - cycleStart

	target := 1000000000
	target -= cycleStart + 1 // 0-index start

	endState := cycleStart + (target % cycleLength)
	fmt.Printf("Part 2: %d\n", scores[endState])
	return 0
}
