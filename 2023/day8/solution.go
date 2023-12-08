package day8

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strings"
)

type Adj struct {
	left  string
	right string
}

type Map map[string]Adj

type Document struct {
	instructions []rune
	mp           Map
}

// Targets the ZZZ node
func (d Document) getOut1(from string) int {
	lInt := len(d.instructions)
	cur := from
	count := 0
	for cur != "ZZZ" {
		inst := d.instructions[count%lInt]
		adj, _ := d.mp[cur]
		if inst == 'L' {
			cur = adj.left
		}

		if inst == 'R' {
			cur = adj.right
		}
		count += 1
	}
	return count
}

// Targets any node that ends in Z
func (d Document) getOut2(from string) int {
	lInt := len(d.instructions)
	cur := from
	count := 0
	for cur[2] != 'Z' {
		inst := d.instructions[count%lInt]
		adj, _ := d.mp[cur]
		if inst == 'L' {
			cur = adj.left
		}

		if inst == 'R' {
			cur = adj.right
		}
		count += 1
	}
	return count
}

func (d Document) possibleStarts() []string {
	res := []string{}
	for start, _ := range d.mp {
		if start[2] == 'A' {
			res = append(res, start)
		}
	}
	return res
}

func parseDocument(lines []string) (Document, error) {
	doc := Document{}

	if len(lines) < 3 {
		return doc, fmt.Errorf("malformed document: %v", lines)
	}

	for _, r := range lines[0] {
		if r != 'L' && r != 'R' {
			return doc, fmt.Errorf("unknown instruction: %s", string(r))
		}
		doc.instructions = append(doc.instructions, r)
	}

	doc.mp = Map{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]

		fromTo := strings.Split(line, " = ")

		if len(fromTo) != 2 {
			return doc, fmt.Errorf("malformed line: %s", line)
		}

		leftRight := strings.Split(fromTo[1], ", ")
		if len(leftRight) != 2 {
			return doc, fmt.Errorf("malformed left,right pair: %s", fromTo[1])
		}

		if len(leftRight[0]) != 4 || len(leftRight[1]) != 4 {
			return doc, fmt.Errorf("malformed left,right pair: %s", fromTo[1])
		}

		adj := Adj{leftRight[0][1:], leftRight[1][:3]}
		doc.mp[fromTo[0]] = adj
	}
	return doc, nil
}

func Solution() int {
	doc, err := parse.GetLinesAsOne[Document]("day8/input.txt", parseDocument)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", doc.getOut1("AAA"))

	// Data analysis shows that for all the A nodes if it takes N step to reach
	// a Z node, it then takes N steps to get back to the same Z node and it
	// carries on forever. So the anser is the Lowest Common Multiple of all
	// the steps required to get to a Z node for each start.
	nbs := []int{}
	for _, st := range doc.possibleStarts() {
		nbs = append(nbs, doc.getOut2(st))
	}
	fmt.Printf("Part 2: %d\n", utils.LCM(nbs...))
	return 0
}
