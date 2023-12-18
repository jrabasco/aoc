package day18

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strconv"
	"strings"
)

type Instruction struct {
	dir    utils.Direction
	amount int
	colour string
}

func (inst Instruction) Apply(p grid.Point) grid.Point {
	return p.Move(inst.dir, inst.amount)
}

func (inst Instruction) Convert() (Instruction, error) {
	res := Instruction{}
	if len(inst.colour) != 6 {
		return res, fmt.Errorf("invalid colour: %s", inst.colour)
	}
	switch inst.colour[5] {
	case '0':
		res.dir = utils.RIGHT
	case '1':
		res.dir = utils.DOWN
	case '2':
		res.dir = utils.LEFT
	case '3':
		res.dir = utils.UP
	default:
		return res, fmt.Errorf("invalid direction: %s", string(inst.colour[5]))
	}

	am64, err := strconv.ParseInt(inst.colour[:5], 16, 64)
	if err != nil {
		return res, err
	}
	// I'm running this on a 64 bit arch...
	res.amount = int(am64)
	return res, nil
}

func parseInstruction(line string) (Instruction, error) {
	inst := Instruction{}
	flds := strings.Fields(line)
	if len(flds) != 3 {
		return inst, fmt.Errorf("invalid instruction line: %s", line)
	}

	switch flds[0] {
	case "R":
		inst.dir = utils.RIGHT
	case "D":
		inst.dir = utils.DOWN
	case "L":
		inst.dir = utils.LEFT
	case "U":
		inst.dir = utils.UP
	default:
		return inst, fmt.Errorf("invalid direction: %s", flds[0])
	}

	amnt, err := strconv.Atoi(flds[1])
	if err != nil {
		return inst, err
	}
	inst.amount = amnt

	if len(flds[2]) != 9 {
		return inst, fmt.Errorf("invalid colour: %s", flds[2])
	}
	inst.colour = flds[2][2:8]
	return inst, nil
}

// Why this works is a bit perplexing... If you search for "area of a
// polygon from its coordinates" you will find what's called the
// "Shoelace formula" : https://www.mathopenref.com/coordpolygonarea.html.
// Using that formula results in finding something that's basically missing
// half of the border. This is because we have a border of non-negligible
// size and the formula ends up "excluding" the bits that are on the
// outside but not facing the origin. Adding a correction of border/2 + 1
// fixes that.
func area(insts []Instruction) int {
	cur := grid.Point{0, 0}
	doubleArea := 0
	borderLen := 0
	for _, inst := range insts {
		next := inst.Apply(cur)
		doubleArea += cur.X*next.Y - next.X*cur.Y
		borderLen += inst.amount
		cur = next
	}
	return utils.Abs(doubleArea)/2 + borderLen/2 + 1
}

func Solution() int {
	insts, err := parse.GetLinesAs[Instruction]("day18/input.txt", parseInstruction)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", area(insts))
	ninsts := []Instruction{}
	for _, inst := range insts {
		ninst, err := inst.Convert()
		if err != nil {
			fmt.Printf("Failed to convert %s: %v\n", inst, err)
			return 1
		}
		ninsts = append(ninsts, ninst)
	}
	fmt.Printf("Part 2: %d\n", area(ninsts))
	return 0
}
