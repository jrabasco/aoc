package day7

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"strconv"
	"strings"
)

type Equation struct {
	goal int
	vals []int
}

func parseEquation(line string) (Equation, error) {
	res := Equation{0, []int{}}
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return res, fmt.Errorf("invalid line :%s", line)
	}

	goal, err := strconv.Atoi(parts[0])
	if err != nil {
		return res, fmt.Errorf("error parsing goal: %v", err)
	}
	res.goal = goal

	valsStr := strings.Fields(parts[1])
	for _, valStr := range valsStr {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return res, fmt.Errorf("error parsing value %s : %v", valStr, err)
		}
		res.vals = append(res.vals, val)
	}
	return res, nil
}

type op func(a, b int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	acc := b
	for acc > 0 {
		a *= 10
		acc /= 10
	}
	return a + b
}

func satisfyable(goal int, cur int, numbers []int, ops []op) bool {
	if cur > goal {
		return false
	}

	if len(numbers) == 0 {
		return goal == cur
	}

	head := numbers[0]
	tail := numbers[1:]
	for _, o := range ops {
		if satisfyable(goal, o(cur, head), tail, ops) {
			return true
		}
	}
	return false
}

func px(eqs []Equation, ops []op) int {
	res := 0
	for _, eq := range eqs {
		if satisfyable(eq.goal, eq.vals[0], eq.vals[1:], ops) {
			res += eq.goal
		}
	}
	return res
}

func p1(eqs []Equation) int {
	ops := []op{add, mul}
	return px(eqs, ops)
}

func p2(eqs []Equation) int {
	ops := []op{add, mul, concat}
	return px(eqs, ops)
}

func Solution() int {
	parsed, err := parse.GetLinesAs[Equation]("day7/input.txt", parseEquation)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %v\n", p1(parsed))
	fmt.Printf("Part 2: %v\n", p2(parsed))
	return 0
}
