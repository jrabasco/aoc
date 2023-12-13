package day13

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
)

type Grid struct {
	byRows []string
	byCols []string
}

// assumes rectangles
func parseGrids(lines []string) []Grid {
	grids := []Grid{}
	g := Grid{[]string{}, []string{}}
	for i, line := range lines {
		if line != "" {
			g.byRows = append(g.byRows, line)
		}
		// assume found at least one non-empty line
		if line == "" || i == len(lines)-1 {
			rowLen := len(g.byRows[0])
			for j := 0; j < rowLen; j++ {
				col := ""
				for i := 0; i < len(g.byRows); i++ {
					col += string(g.byRows[i][j])
				}
				g.byCols = append(g.byCols, col)
			}
			grids = append(grids, g)
			g = Grid{[]string{}, []string{}}
		}
	}
	return grids
}

// return which row and which col is a symmetry line
// one should always be -1
// also returns a second pair for the smudged equivalent
func (g Grid) findSymmetry() (int, int, int, int) {
	rSym, rSymS := findSymmetry(g.byRows)
	if rSym != -1 && rSymS != -1 {
		return rSym, -1, rSymS, -1
	}

	cSym, cSymS := findSymmetry(g.byCols)
	return rSym, cSym, rSymS, cSymS
}

type Poss struct {
	noSmudge  bool
	smudge    bool // if we find diffs on two different places, this turns to false
	smudgePos int  // when first encountering a diff, where is the smudge
}

// returns the number of differences and the position of the first difference
func strDiff(a, b string) (int, int) {
	// this should not happen in our case but we'll return a number of diffs
	// larger than 1 to avoid considering this a smudge
	if len(a) != len(b) {
		return 2, -1
	}
	diffs := 0
	pos := -1
	for i := range a {
		if a[i] != b[i] {
			if pos == -1 {
				pos = i
			}

			diffs += 1

			// for our problem, more than one diff doesn't make a difference
			// anymore
			if diffs > 1 {
				return diffs, pos
			}
		}
	}
	return diffs, pos
}

// Find symmetry line and potentially one with a smudge
func findSymmetry(lines []string) (int, int) {
	possLines := []Poss{}
	nLines := len(lines)
	// after which line in the input can a symmetry line exist
	for i := 0; i < nLines-1; i++ {
		possLines = append(possLines, Poss{true, true, -1})
	}

	// For each line, compare it to all possible reflections
	// - if it's not equal then mark that reflection as impossible
	// - if the first reflection is still possible (no previous line marked it
	//   false) and the current line matches its reflection through that first
	//   one, then we found the symmetry
	// For the smudge version, when we find a difference we record the
	// potential location of the smudge and use it in upcoming comparisons
	noSmudgePos := -1
	smudgePos := -1
	for i, line := range lines {
		// we reached the end
		if i >= len(possLines) {
			break
		}

		// check the next line
		diffsNL, posNL := strDiff(line, lines[i+1])
		// if it's still possible without a smudge, check the diff to see if
		// this last line matches
		// if it's possible with one smudge, check if the position of the
		// difference matches
		if possLines[i].noSmudge && noSmudgePos == -1 {
			if diffsNL == 0 {
				noSmudgePos = i
			}
		}

		// - we know of a possible smudge (smudgePos != -1), there is a
		//   diff of 1 and the current diff matched the smudge -> i is a
		//   symmetry point
		// - we know of a possible smudge (smudgePos != -1) and there is
		//   no current diff -> i is a symmetry point
		// - we don't know of any smudge, but we don't know for sure there
		//   aren't any (smudge was not set to false) and there is only one
		//   diff -> i is a symmetry point with the  smudge at this diff
		// CAREFUL: if we don't know of any smudge and there aren't any
		// diffs then this would fall under the previous condition
		if possLines[i].smudge && smudgePos == -1 {

			if possLines[i].smudgePos != -1 {
				if diffsNL == 0 {
					smudgePos = i
				} else if diffsNL == 1 && possLines[i].smudgePos == posNL {
					smudgePos = i
				}
			} else if diffsNL == 1 {
				smudgePos = i
			}
		}

		if noSmudgePos != -1 && smudgePos != -1 {
			return noSmudgePos, smudgePos
		}

		// check each reflection beyond that first one
		for j := i + 1; j < len(possLines); j++ {
			// we already know this one is not possible
			if !possLines[j].noSmudge && possLines[j].smudgePos == -1 {
				continue
			}

			// which line would be the reflection of line i if we put a mirror
			// after line j
			reflectedIdx := i + 2*(j-i) + 1

			// beyond the lines we have so we can skip
			if reflectedIdx >= nLines {
				break
			}

			diffs, pos := strDiff(line, lines[reflectedIdx])
			if diffs > 0 {
				possLines[j].noSmudge = false

				// we determined that this line needed at least two smudges
				// already
				if !possLines[j].smudge {
					continue
				}

				// only one diff
				if diffs == 1 {
					if possLines[j].smudgePos == -1 {
						// first time we see a possibility with a smudge for that
						// symmetry line (we already checked the smudge boolean
						// that would tell us it's definitely not possible)
						possLines[j].smudgePos = pos
					} else if possLines[j].smudgePos != pos {
						// we had already found a potential smudge but now
						// this one would be at a different place
						possLines[j].smudge = false
						possLines[j].smudgePos = -1
					}
				} else {
					// two or more diffs is definitely impossible
					possLines[j].smudge = false
					possLines[j].smudgePos = -1
				}
			}
		}
	}

	return noSmudgePos, smudgePos
}

func Solution() int {
	parsed, err := parse.GetLines("day13/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	grids := parseGrids(parsed)

	res := 0
	res2 := 0
	for _, grid := range grids {
		rSym, cSym, rSymS, cSymS := grid.findSymmetry()
		// those work because findSymmetry returns -1 when it doesn't find
		// anything
		res += cSym + 1
		res += (rSym + 1) * 100
		res2 += cSymS + 1
		res2 += (rSymS + 1) * 100
	}

	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
