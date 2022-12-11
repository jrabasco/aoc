package day11

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
	"strings"
)

type Monkey struct {
	id      int
	items   []int
	op      func(int) int
	test    func(int) int
	divisor int
	insp    int
}

func (m *Monkey) PopItem() (int, error) {
	if len(m.items) < 1 {
		return -1, fmt.Errorf("not enough items")
	}
	item := m.items[0]
	m.items = m.items[1:]
	return item, nil
}

func parseMonkeyID(line string) (int, error) {
	// parse monkey ID
	pts := strings.Fields(line)
	if len(pts) != 2 {
		return 0, fmt.Errorf("invalid monkey line: %s", line)
	}

	idS := strings.TrimSuffix(pts[1], ":")
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, fmt.Errorf("could not parse ID: %v", err)
	}

	return id, nil
}

func parseMonkeyItems(line string) ([]int, error) {
	ptItems := strings.Split(line, ": ")
	if len(ptItems) != 2 {
		return nil, fmt.Errorf("invalid items line: %s", line)
	}
	itemsStr := strings.Split(ptItems[1], ", ")
	items := []int{}

	for _, itemS := range itemsStr {
		item, err := strconv.Atoi(itemS)
		if err != nil {
			return nil, fmt.Errorf("could not parse item: %v", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func parseBinOp(op string) (func(int, int) int, error) {
	switch op {
	case "*":
		return func(a int, b int) int { return a * b }, nil
	case "+":
		return func(a int, b int) int { return a + b }, nil
	case "-":
		return func(a int, b int) int { return a - b }, nil
	case "/":
		return func(a int, b int) int { return a / b }, nil
	default:
		return func(a int, b int) int { return a }, fmt.Errorf("invalid binary operator: %s", op)
	}
}

func parseMonkeyOp(line string) (func(int) int, error) {
	id := func(e int) int { return e }
	pts := strings.Split(line, "new = old ")
	if len(pts) != 2 {
		return id, fmt.Errorf("imparsable op: %s", line)
	}

	opsPts := strings.Fields(pts[1])
	if len(opsPts) != 2 {
		return id, fmt.Errorf("imparsable op: %s", line)
	}

	binOp, err := parseBinOp(opsPts[0])
	if err != nil {
		return id, err
	}

	if opsPts[1] == "old" {
		return func(e int) int { return binOp(e, e) }, nil
	}

	operand, err := strconv.Atoi(opsPts[1])
	if err != nil {
		return id, fmt.Errorf("invalid operand: %s", opsPts[1])
	}

	return func(e int) int { return binOp(e, operand) }, nil
}

func parseMonkeyTest(lines []string) (func(int) int, int, error) {
	id := func(e int) int { return e }
	if len(lines) != 3 {
		return id, 0, fmt.Errorf("not enough lines for test")
	}

	pts1 := strings.Fields(lines[0])
	divS := pts1[len(pts1)-1]
	divisor, err := strconv.Atoi(divS)
	if err != nil {
		return id, 0, fmt.Errorf("invalid divisor: %s", divS)
	}

	pts2 := strings.Fields(lines[1])
	tTrueS := pts2[len(pts2)-1]
	tTrue, err := strconv.Atoi(tTrueS)
	if err != nil {
		return id, 0, fmt.Errorf("invalid divisor: %s", tTrueS)
	}

	pts3 := strings.Fields(lines[2])
	tFalseS := pts3[len(pts3)-1]
	tFalse, err := strconv.Atoi(tFalseS)
	if err != nil {
		return id, 0, fmt.Errorf("invalid divisor: %s", tFalseS)
	}

	return func(e int) int {
		if e%divisor == 0 {
			return tTrue
		} else {
			return tFalse
		}
	}, divisor, nil
}

func parseMonkey(lines []string) (Monkey, error) {
	var monkey Monkey
	if len(lines) != 6 {
		return monkey, fmt.Errorf("not enough lines to parse a monkey: %d", len(lines))
	}

	id, err := parseMonkeyID(lines[0])
	if err != nil {
		return monkey, err
	}
	monkey.id = id

	items, err := parseMonkeyItems(lines[1])
	if err != nil {
		return monkey, err
	}
	monkey.items = items

	op, err := parseMonkeyOp(lines[2])
	if err != nil {
		return monkey, err
	}
	monkey.op = op

	test, divisor, err := parseMonkeyTest(lines[3:])
	if err != nil {
		return monkey, err
	}
	monkey.test = test
	monkey.divisor = divisor

	return monkey, nil
}

func round(monkeys *[]Monkey, divisor int, relax bool) {
	for i := 0; i < len(*monkeys); i++ {
		monkey := &(*monkeys)[i]
		for len(monkey.items) > 0 {
			// can ignore error as per loop condition
			item, _ := monkey.PopItem()
			item = monkey.op(item)
			monkey.insp += 1
			if relax {
				item /= 3
			} else {
				item = item % divisor
			}
			target := monkey.test(item)
			(*monkeys)[target].items = append((*monkeys)[target].items, item)
		}
	}
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day11/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	monkeys := []Monkey{}
	for i := 0; i < len(parsed); i++ {
		monkey, err := parseMonkey(parsed[i : i+6])
		if err != nil {
			fmt.Printf("Failed to parse monkey from %d: %v\n", i, err)
			return 1
		}

		monkeys = append(monkeys, monkey)
		i += 6
	}

	// Since all divisors are prime and
	// where to throw the ball is only based on divisibility
	// we can perform all operations modulo the product of the
	// divisors to avoid large numbers
	genDiv := 1

	for _, monkey := range monkeys {
		genDiv *= monkey.divisor
	}

	relax := false
	rounds := 10000
	if part == "1" {
		relax = true
		rounds = 20
	}

	for i := 0; i < rounds; i++ {
		round(&monkeys, genDiv, relax)
	}

	maxs := [2]int{0, 0}

	for _, monkey := range monkeys {
		if monkey.insp > maxs[0] {
			tmp := maxs[0]
			maxs[0] = monkey.insp
			maxs[1] = tmp
		} else if monkey.insp > maxs[1] {
			maxs[1] = monkey.insp
		}
	}
	fmt.Printf("Part %s: %v\n", part, maxs[0]*maxs[1])
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
