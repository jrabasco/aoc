package day5

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"strconv"
	"strings"
)

type CrateStack = utils.Stack[string]

func splitCratesMoves(lines []string) ([]string, []string) {
	var crateLines []string
	var moveLines []string
	moves := false
	for _, line := range lines {
		if moves {
			moveLines = append(moveLines, line)
			continue
		}

		if line == "" {
			moves = true
			continue
		}
		crateLines = append(crateLines, line)
	}

	l := len(crateLines)
	crateLines = crateLines[:l-1]
	return crateLines, moveLines
}

func getStacks(crateLines []string) []CrateStack {
	var stacks []CrateStack
	if len(crateLines) == 0 || len(crateLines[0]) == 0 {
		return stacks
	}
	depth := len(crateLines)
	// figure out number of stacks
	l := len(crateLines[0])
	// we know l is 3n + n - 1 = 4n - 1 (3 chars per stack + spaces in between)
	// so n is (l + 1)/4
	n := (l + 1) / 4
	for i := 0; i < n; i++ {
		stacks = append(stacks, utils.NewStack[string]())
	}
	for i := depth - 1; i >= 0; i-- {
		line := crateLines[i]
		for j := 0; j < n; j++ {
			crate := line[4*j : 4*j+3]
			if crate != "   " {
				stacks[j].Push(crate[1:2])
			}
		}
	}
	return stacks
}

type Move struct {
	from int
	to   int
	qty  int
}

func getMoves(moveLines []string) ([]Move, error) {
	var moves []Move
	for _, moveStr := range moveLines {
		tokens := strings.Fields(moveStr)
		if len(tokens) != 6 {
			return nil, fmt.Errorf("invalid line: %s", moveStr)
		}
		qty, err := strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			return nil, err
		}

		from, err := strconv.ParseInt(tokens[3], 10, 64)
		if err != nil {
			return nil, err
		}

		to, err := strconv.ParseInt(tokens[5], 10, 64)
		if err != nil {
			return nil, err
		}

		moves = append(moves, Move{int(from), int(to), int(qty)})
	}
	return moves, nil
}

func applyMoves1(stacks *[]CrateStack, moves []Move) error {
	stackN := len(*stacks)
	for _, move := range moves {
		if move.from > stackN {
			return fmt.Errorf("invalid from position: %d", move.from)
		}

		if move.to > stackN {
			return fmt.Errorf("invalid to position: %d", move.to)
		}
		for i := 0; i < move.qty; i++ {
			// the stacks are 1-indexed
			crate, err := (*stacks)[move.from-1].Pop()
			if err != nil {
				return err
			}
			(*stacks)[move.to-1].Push(crate)
		}
	}
	return nil
}

func applyMoves2(stacks *[]CrateStack, moves []Move) error {
	stackN := len(*stacks)
	for _, move := range moves {
		if move.from > stackN {
			return fmt.Errorf("invalid from position: %d", move.from)
		}

		if move.to > stackN {
			return fmt.Errorf("invalid to position: %d", move.to)
		}

		// temporary stack to move crates to
		// will take care of re-ordering
		var mCrates CrateStack
		for i := 0; i < move.qty; i++ {
			// the stacks are 1-indexed
			crate, err := (*stacks)[move.from-1].Pop()
			if err != nil {
				return err
			}
			mCrates.Push(crate)
		}

		for !mCrates.Empty() {
			crate, _ := mCrates.Pop()
			(*stacks)[move.to-1].Push(crate)
		}
	}
	return nil
}

func extractTops(stacks *[]CrateStack) (string, error) {
	tops := ""
	for _, s := range *stacks {
		top, err := s.Peek()
		if err != nil {
			return "", fmt.Errorf("could not peek: %v", err)
		}
		tops += top
	}
	return tops, nil
}

func Solution() int {
	parsed, err := parse.GetLines("day5/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	crateLines, moveLines := splitCratesMoves(parsed)
	stacks1 := getStacks(crateLines)
	// need a copy
	stacks2 := getStacks(crateLines)
	moves, err := getMoves(moveLines)
	if err != nil {
		fmt.Printf("Error parsing moves: %v\n", err)
		return 1
	}

	err = applyMoves1(&stacks1, moves)
	if err != nil {
		fmt.Printf("Error applying moves: %v\n", err)
		return 1
	}

	err = applyMoves2(&stacks2, moves)
	if err != nil {
		fmt.Printf("Error applying moves: %v\n", err)
		return 1
	}

	part1, err := extractTops(&stacks1)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Printf("Part 1: %s\n", part1)

	part2, err := extractTops(&stacks2)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Printf("Part 2: %s\n", part2)
	return 0
}
