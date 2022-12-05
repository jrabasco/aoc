package day2

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strings"
)

type Move int64

const (
	Rock Move = iota
	Paper
	Scissors
	Fail
)

func (m Move) String() string {
	switch m {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return "Fail"
	}
}

var ring = []Move{Rock, Scissors, Paper}
var idx = map[Move]int{
	Rock:     0,
	Scissors: 1,
	Paper:    2,
}

func (mine Move) Beats(theirs Move) bool {
	loserIdx := (idx[mine] + 1) % 3
	return theirs == ring[loserIdx]
}

var scores = map[Move]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
	Fail:     0,
}

type Round struct {
	mine   Move
	theirs Move
}

func (r Round) String() string {
	return fmt.Sprintf("%s-%s", r.mine, r.theirs)
}

func (r Round) Score() int {
	if r.mine == r.theirs {
		return 3 + scores[r.mine]
	} else if r.mine.Beats(r.theirs) {
		return 6 + scores[r.mine]
	} else {
		return scores[r.mine]
	}
}

type Strategy []Round

func (s Strategy) Score() int {
	res := 0
	for _, r := range s {
		score := r.Score()
		res += score
	}
	return res
}

func sToMove(c string) (Move, error) {
	if c == "A" || c == "X" {
		return Rock, nil
	}

	if c == "B" || c == "Y" {
		return Paper, nil
	}

	if c == "C" || c == "Z" {
		return Scissors, nil
	}
	return Fail, fmt.Errorf("invalid move: %s", c)
}

func lToRound(line string, tInfer func(Move, string) (Move, error)) (Round, error) {
	parts := strings.Fields(line)

	if len(parts) != 2 {
		return Round{Fail, Fail}, fmt.Errorf("invalid round: %s", line)
	}

	theirs, err := sToMove(parts[0])
	if err != nil {
		return Round{Fail, Fail}, err
	}

	mine, err := tInfer(theirs, parts[1])
	if err != nil {
		return Round{Fail, Fail}, err
	}

	return Round{mine, theirs}, nil
}

func lToRound1(line string) (Round, error) {
	return lToRound(line, func(m Move, s string) (Move, error) {
		return sToMove(s)
	})
}

func inferMine(theirs Move, wld string) (Move, error) {
	// losing
	if wld == "X" {
		loserIdx := (idx[theirs] + 1) % 3
		return ring[loserIdx], nil
	}

	// draw
	if wld == "Y" {
		return theirs, nil
	}

	// win
	if wld == "Z" {
		winnerIdx := (idx[theirs] + 2) % 3
		return ring[winnerIdx], nil
	}

	return Fail, fmt.Errorf("Impossible w/l/d value: %s", wld)
}

func lToRound2(line string) (Round, error) {
	return lToRound(line, inferMine)
}

func Solution() int {
	var strategy Strategy
	strategy, err := parse.GetLinesAs[Round]("day2/input.txt", lToRound1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", strategy.Score())

	var strategy2 Strategy
	strategy2, err = parse.GetLinesAs[Round]("day2/input.txt", lToRound2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", strategy2.Score())

	return 0
}
