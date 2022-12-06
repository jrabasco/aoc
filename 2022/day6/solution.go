package day6

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
)

type Window struct {
	data   []rune
	mSize  int
	unique map[rune]int
}

func (w *Window) Add(r rune) {
	if len(w.data) == w.mSize {
		rm := w.data[0]
		w.data = w.data[1:]
		if w.unique[rm] > 1 {
			w.unique[rm] -= 1
		} else {
			delete(w.unique, rm)
		}
	}
	w.data = append(w.data, r)
	if _, exists := w.unique[r]; !exists {
		w.unique[r] = 0
	}
	w.unique[r] += 1
}

func (w Window) IsMarker() bool {
	return len(w.data) == w.mSize && len(w.unique) == w.mSize
}

func getMarkerPos(message string, wSize int) (int, error) {
	w := Window{[]rune{}, wSize, map[rune]int{}}
	for i, r := range message {
		w.Add(r)
		if w.IsMarker() {
			return i + 1, nil
		}
	}
	return -1, fmt.Errorf("could not find marker in: %s", message)
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day6/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	for _, message := range parsed {
		var wSize int
		if part == "1" {
			wSize = 4
		} else {
			wSize = 14
		}
		markerPos, err := getMarkerPos(message, wSize)
		if err != nil {
			fmt.Printf("Failed to solve: %v\n", err)
			return 1
		}
		fmt.Printf("Part %s: %d\n", part, markerPos)
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
