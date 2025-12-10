package day10

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
	"strconv"
	"strings"
)

type Problem struct {
	target  int
	buttons []int
}

func p1(problems []Problem) int {
	res := 0
	for _, pb := range problems {
		mx := utils.IntPow(2, len(pb.buttons))
		combs := []int{}
		for i := 0; i < mx; i++ {
			combs = append(combs, i)
		}
		minCnt := mx + 1
		for _, comb := range combs {
			curV := 0
			mask := comb
			for _, bt := range pb.buttons {
				if mask%2 == 1 {
					curV = curV ^ bt
				}
				mask /= 2
			}
			if curV == pb.target {
				cnt := utils.BitCount(comb)
				if cnt < minCnt {
					minCnt = cnt
				}
			}
		}
		fmt.Println(minCnt)
		res += minCnt
	}
	return res
}

func p2(problems []Problem) int {
	return 0
}

func parseProblem(line string) (Problem, error) {
	res := Problem{}
	parts := strings.Split(line, " ")
	if len(parts) < 3 {
		return res, fmt.Errorf("invalid problem: %s", line)
	}

	lights := parts[0]
	ll := len(lights)
	if ll < 3 || lights[0] != '[' || lights[ll-1] != ']' {
		return res, fmt.Errorf("invalid lights: %s", lights)
	}

	target := 0
	pow := 1
	for i := ll - 2; i > 0; i-- {
		if lights[i] == '#' {
			target += pow
		}
		pow *= 2
	}
	res.target = target

	for i := 1; i < len(parts)-1; i++ {
		bt := parts[i]
		lb := len(bt)
		if lb < 3 || bt[0] != '(' || bt[lb-1] != ')' {
			return res, fmt.Errorf("invalid button: %s", bt)
		}
		bt = bt[1 : lb-1]
		pbt := strings.Split(bt, ",")
		btint := 0
		for _, v := range pbt {
			exp, err := strconv.Atoi(v)
			if err != nil {
				return res, err
			}
			btint += utils.IntPow(2, exp)
		}
		res.buttons = append(res.buttons, btint)
	}
	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[Problem]("day10/input_test.txt", parseProblem)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
