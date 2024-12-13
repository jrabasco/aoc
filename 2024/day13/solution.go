package day13

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"math"
	"strconv"
	"strings"
)

func parseButton(line string) (int, int) {
	parts := strings.Split(line, ": ")
	nbparts := strings.Split(parts[1], ", ")
	xparts := strings.Split(nbparts[0], "+")
	yparts := strings.Split(nbparts[1], "+")
	x, err := strconv.Atoi(xparts[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yparts[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

func parsePrize(line string) (int, int) {
	parts := strings.Split(line, ": ")
	nbparts := strings.Split(parts[1], ", ")
	xparts := strings.Split(nbparts[0], "=")
	yparts := strings.Split(nbparts[1], "=")
	x, err := strconv.Atoi(xparts[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yparts[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

type Machine struct {
	a0 int
	b0 int
	a1 int
	b1 int
	c0 int
	c1 int
}

func parseMachine(lines []string, i int) Machine {
	a0, a1 := parseButton(lines[i])
	b0, b1 := parseButton(lines[i+1])
	c0, c1 := parsePrize(lines[i+2])
	return Machine{a0, b0, a1, b1, c0, c1}
}

func parseAll(lines []string) []Machine {
	res := []Machine{}
	for i := 0; i < len(lines)-2; i += 4 {
		res = append(res, parseMachine(lines, i))
	}
	return res
}

func solveMachine(m Machine) (int, int, bool) {
	xf := 0.0
	yf := 0.0
	x := 0
	y := 0
	a0 := float64(m.a0)
	a1 := float64(m.a1)
	b0 := float64(m.b0)
	b1 := float64(m.b1)
	c0 := float64(m.c0)
	c1 := float64(m.c1)
	if a0 == 0.0 {
		return x, y, false
	}

	ytop := (c1 - (a1*c0)/a0)
	ybot := (b1 - (a1*b0)/a0)
	if ybot == 0.0 {
		return x, y, false
	}

	yf = ytop / ybot
	if yf <= 0 {
		return x, y, false
	}

	xtop := c0 - b0*yf

	xf = xtop / a0
	if xf <= 0 {
		return x, y, false
	}

	// this is necessary because of float precision
	x = int(math.Round(xf))
	y = int(math.Round(yf))
	if m.a0*x+m.b0*y != m.c0 || m.a1*x+m.b1*y != m.c1 {
		return x, y, false
	}
	return x, y, true
}

func p1(machines []Machine) int {
	res := 0
	for _, m := range machines {
		x, y, possible := solveMachine(m)
		if possible {
			res += x*3 + y
		}
	}
	return res
}

func p2(machines []Machine) int {
	res := 0
	for _, m := range machines {
		m.c0 += 10000000000000
		m.c1 += 10000000000000
		x, y, possible := solveMachine(m)
		if possible {
			res += x*3 + y
		}
	}
	return res
}

func Solution() int {
	parsed, err := parse.GetLines("day13/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	machines := parseAll(parsed)
	fmt.Printf("Part 1: %d\n", p1(machines))
	fmt.Printf("Part 2: %d\n", p2(machines))
	return 0
}
