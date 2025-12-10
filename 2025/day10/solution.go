package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/jrabasco/aoc/2025/framework/parse"
	"github.com/jrabasco/aoc/2025/framework/utils"
)

type Problem struct {
	target   int
	buttons  []int
	joltages []int // For Part 2
}

func p1(problems []Problem) int {
	res := 0
	for _, pb := range problems {
		mx := utils.IntPow(2, len(pb.buttons))
		minCnt := mx + 1
		for comb := 0; comb < mx; comb++ {
			curV := 0
			for i, bt := range pb.buttons {
				// Check if bit i is set in comb
				if (comb & (1 << i)) != 0 {
					curV = curV ^ bt
				}
			}
			if curV == pb.target {
				cnt := utils.BitCount(comb)
				if cnt < minCnt {
					minCnt = cnt
				}
			}
		}
		res += minCnt
	}
	return res
}

func p2(problems []Problem) int {
	res := 0
	for _, pb := range problems {
		if len(pb.joltages) == 0 {
			continue
		}

		numIndicators := len(pb.joltages)
		minPresses := solvePart2WithLP(pb.buttons, pb.joltages, numIndicators)
		res += minPresses
	}
	return res
}

func solvePart2WithLP(buttons []int, targets []int, numIndicators int) int {
	lp := golp.NewLP(0, len(buttons))

	// Set objective: minimize sum of button presses
	objCoeffs := make([]float64, len(buttons))
	for i := range objCoeffs {
		objCoeffs[i] = 1.0
	}
	lp.SetObjFn(objCoeffs)

	// Add constraints for each indicator
	for indicatorIdx := 0; indicatorIdx < numIndicators; indicatorIdx++ {
		// Build constraint: sum of button effects = target
		entries := []golp.Entry{}
		for btnIdx, button := range buttons {
			bitPos := numIndicators - 1 - indicatorIdx
			if (button & (1 << bitPos)) != 0 {
				entries = append(entries, golp.Entry{Col: btnIdx, Val: 1.0})
			}
		}
		lp.AddConstraintSparse(entries, golp.EQ, float64(targets[indicatorIdx]))
	}

	// Set variable bounds (non-negative integers)
	for i := 0; i < len(buttons); i++ {
		lp.SetInt(i, true) // Make variable integer
	}

	// Solve
	lp.Solve()

	// Get the objective value
	obj := lp.Objective()

	// Verify solution is valid (non-negative)
	if obj < 0 {
		return -1
	}

	return int(obj + 0.5) // Round to nearest integer
}

func stateKey(levels []int) string {
	var sb strings.Builder
	for i, v := range levels {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

func statesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

	// Count actual light positions
	numLights := ll - 2

	target := 0
	pow := 1
	for i := ll - 2; i > 0; i-- {
		if lights[i] == '#' {
			target += pow
		}
		pow *= 2
	}
	res.target = target

	// Find the joltages (last element in curly braces)
	lastPart := parts[len(parts)-1]
	joltageStart := -1

	// Parse buttons (everything between lights and joltages)
	endIdx := len(parts) - 1
	if len(lastPart) > 0 && lastPart[0] == '{' {
		joltageStart = len(parts) - 1
		endIdx = len(parts) - 1
	}

	for i := 1; i < endIdx; i++ {
		bt := parts[i]
		lb := len(bt)
		if lb < 3 || bt[0] != '(' || bt[lb-1] != ')' {
			return res, fmt.Errorf("invalid button: %s", bt)
		}
		bt = bt[1 : lb-1]
		pbt := strings.Split(bt, ",")
		btint := 0
		for _, v := range pbt {
			pos, err := strconv.Atoi(v)
			if err != nil {
				return res, err
			}
			// Convert position from left to bit position
			// Position 0 from left = bit (numLights-1)
			bitPos := numLights - 1 - pos
			btint += utils.IntPow(2, bitPos)
		}
		res.buttons = append(res.buttons, btint)
	}

	// Parse joltages if present
	if joltageStart != -1 {
		joltStr := parts[joltageStart]
		if len(joltStr) < 3 || joltStr[0] != '{' || joltStr[len(joltStr)-1] != '}' {
			return res, fmt.Errorf("invalid joltages: %s", joltStr)
		}
		joltStr = joltStr[1 : len(joltStr)-1]
		joltParts := strings.Split(joltStr, ",")
		for _, v := range joltParts {
			val, err := strconv.Atoi(v)
			if err != nil {
				return res, err
			}
			res.joltages = append(res.joltages, val)
		}
	}

	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[Problem]("day10/input.txt", parseProblem)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
