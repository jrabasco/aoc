package day6

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"strconv"
	"strings"
)

type Op func(int, int) int

type Problem struct {
	numbers []int
	op      Op
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func p1(lines []string) int {
	problems, err := parseProblems(lines)
	if err != nil {
		panic(err)
	}
	res := 0
	for _, p := range problems {
		acc := p.numbers[0]
		for i := 1; i < len(p.numbers); i++ {
			acc = p.op(acc, p.numbers[i])
		}
		res += acc
	}
	return res
}

func p2(lines []string) int {
	res := 0
	// every line in the input has the same length
	nbRunes := [][]rune{}
	for j := len(lines[0]) - 1; j >= 0; j-- {
		nbR := []rune{}
		opR := ' '
		for i := range lines {
			r := rune(lines[i][j])
			if r == ' ' {
				continue
			}
			if r != '+' && r != '*' {
				nbR = append(nbR, r)
				continue
			}
			opR = r
		}
		nbRunes = append(nbRunes, nbR)
		if opR == ' ' {
			continue
		}
		op := add
		if opR == '*' {
			op = mul
		}
		nbInts := []int{}
		for _, nbr := range nbRunes {
			nbS := string(nbr)
			if len(nbS) == 0 {
				continue
			}
			nb, err := strconv.Atoi(nbS)
			if err != nil {
				panic(err)
			}
			nbInts = append(nbInts, nb)
		}
		acc := nbInts[0]
		for k := 1; k < len(nbInts); k++ {
			acc = op(acc, nbInts[k])
		}
		res += acc
		nbRunes = [][]rune{}
		opR = ' '
	}
	return res
}

func parseProblems(lines []string) ([]Problem, error) {
	res := []Problem{}
	first := true
	for _, line := range lines {
		fields := strings.Fields(line)
		for i, f := range fields {
			if first {
				res = append(res, Problem{[]int{}, add})
			}

			if f == "*" || f == "+" {
				if f == "*" {
					res[i].op = mul
				}
				continue
			}

			nb, err := strconv.Atoi(f)
			if err != nil {
				return res, err
			}
			res[i].numbers = append(res[i].numbers, nb)
		}
		first = false
	}
	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLines("day6/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
