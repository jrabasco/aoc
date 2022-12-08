package day8

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/grid"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
)

type Tree struct {
	height  int
	visible bool
	scores  [4]int //left, right, down, up
}

func NewTree(height int) Tree {
	return Tree{height, false, [4]int{0, 0, 0, 0}}
}

func (t Tree) Score() int {
	return t.scores[0] * t.scores[1] * t.scores[2] * t.scores[3]
}

type Forest = grid.Grid[Tree]

// dir will be 0 for left right, 2 for down up
func computeScores(r []*Tree, dir int) {
	minH := -1
	prevH := -1
	for i := 0; i < len(r); i++ {
		t := r[i]
		if t.height > minH {
			t.visible = true
			minH = t.height
		}
		if t.height <= prevH {
			t.scores[dir] = 1
		} else {
			// walk back to find how far we can see
			if i == 0 {
				t.scores[dir] = 0
				continue
			}

			b := i - 1
			for b >= 0 {
				t.scores[dir] += 1
				bt := r[b]
				if bt.height >= t.height {
					break
				}
				b--
			}
		}
		prevH = t.height
	}

	minH = -1
	prevH = -1
	for i := len(r) - 1; i >= 0; i-- {
		t := r[i]
		if t.height > minH {
			t.visible = true
			minH = t.height
		}
		if t.height <= prevH {
			t.scores[dir+1] = 1
		} else {
			// walk back to find how far we can see
			if i == len(r)-1 {
				t.scores[dir+1] = 0
				continue
			}

			b := i + 1
			for b < len(r) {
				t.scores[dir+1] += 1
				bt := r[b]
				if bt.height >= t.height {
					break
				}
				b++
			}
		}
		prevH = t.height
	}
}

func countVisibles(f Forest) int {
	res := 0
	for _, r := range f.Rows() {
		for _, t := range r {
			if t.visible {
				res++
			}
		}
	}
	return res
}

func findLargestScore(f Forest) int {
	max := 0

	for _, r := range f.Rows() {
		for _, t := range r {
			cScore := t.Score()
			if cScore > max {
				max = cScore
			}
		}
	}
	return max
}

func solvePart(part string) int {
	parsed, err := parse.GetLinesAs[[]string]("day8/input.txt",
		func(line string) ([]string, error) {
			res := []string{}
			for _, c := range line {
				res = append(res, string(c))
			}
			return res, nil
		})
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	f, err := grid.NewGrid[Tree, string](parsed, func(cell string) (Tree, error) {
		h, err := strconv.Atoi(cell)
		return NewTree(h), err
	})

	if err != nil {
		fmt.Printf("Failed to make a grid: %v\n", err)
		return 1
	}

	for _, r := range f.Rows() {
		computeScores(r, 0)
	}

	for _, c := range f.Columns() {
		computeScores(c, 2)
	}

	if part == "1" {
		visible := countVisibles(f)

		fmt.Printf("Part %s: %d\n", part, visible)
	} else {
		maxScore := findLargestScore(f)

		fmt.Printf("Part %s: %d\n", part, maxScore)
	}
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
