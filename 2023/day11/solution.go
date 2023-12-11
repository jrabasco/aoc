package day11

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
)

func findEmptyRows(lines []string) []int {
	res := []int{}
	for i := range lines {
		allDots := true
		for j := range lines[i] {
			if lines[i][j] != '.' {
				allDots = false
				break
			}
		}
		if allDots {
			res = append(res, i)
		}
	}

	return res
}

func findEmptyCols(lines []string) []int {
	res := []int{}
	if len(lines) == 0 {
		return res
	}

	// assumes a rectangle (i.e. all lines have the same length)
	for j := range lines[0] {
		allDots := true
		for i := range lines {
			if lines[i][j] != '.' {
				allDots = false
				break
			}
		}
		if allDots {
			res = append(res, j)
		}
	}
	return res
}

// count how many indices in idxs are smaller than i
// assumes idxs is sorted
func countSmaller(i int, idxs []int) int {
	count := 0
	for _, idx := range idxs {
		if i <= idx {
			break
		}
		count++
	}
	return count
}

type Galaxy struct {
	x int
	y int
}

// expansion is how many rows/cols have to be replaced into the spot
func sumDistances(lines []string, expansion int, emptyRows []int, emptyCols []int) int {
	galaxies := []Galaxy{}
	for i, row := range lines {
		for j, c := range row {
			if c == '#' {
				actualI := i + (expansion-1)*countSmaller(i, emptyRows)
				actualJ := j + (expansion-1)*countSmaller(j, emptyCols)
				galaxies = append(galaxies, Galaxy{actualI, actualJ})
			}
		}
	}

	res := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i >= j {
				continue
			}
			res += utils.AbsDiff(g1.x, g2.x)
			res += utils.AbsDiff(g1.y, g2.y)
		}
	}

	return res
}

func Solution() int {
	lines, err := parse.GetLines("day11/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	emptyRows := findEmptyRows(lines)
	emptyCols := findEmptyCols(lines)
	res1 := sumDistances(lines, 2, emptyRows, emptyCols)
	res2 := sumDistances(lines, 1000000, emptyRows, emptyCols)

	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
