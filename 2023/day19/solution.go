package day19

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strconv"
	"strings"
)

type Subject int

const (
	X Subject = iota
	M
	A
	S
)

type Part []int

func NewPart() Part {
	return Part{0, 0, 0, 0}
}

// undefined behaviour when part is empty is lacking
// some characteristic
// always construct parts with the parsePart function
// or the NewPart function
func (p Part) Rating() int {
	return p[X] + p[M] + p[A] + p[S]
}

func parsePart(line string) (Part, error) {
	res := NewPart()
	// at least the {}
	if len(line) < 2 {
		return res, fmt.Errorf("malformed part: %s", line)
	}

	// remove {}
	line = line[1 : len(line)-1]

	parts := strings.Split(line, ",")
	for _, p := range parts {
		pparts := strings.Split(p, "=")
		if len(pparts) != 2 {
			return res, fmt.Errorf("malformed subpart: %s", p)
		}
		val, err := strconv.Atoi(pparts[1])
		if err != nil {
			return res, err
		}

		switch pparts[0] {
		case "x":
			res[X] = val
		case "m":
			res[M] = val
		case "a":
			res[A] = val
		case "s":
			res[S] = val
		default:
			return res, fmt.Errorf("unknown part attribute: %s", pparts[0])
		}
	}
	if len(res) != 4 {
		return res, fmt.Errorf("part is missing some attribute: %v", res)
	}
	return res, nil
}

type Kind int

const (
	LT Kind = iota
	GT
	DEFAULT
)

type Predicate struct {
	level   int
	kind    Kind
	subject Subject
	res     string
}

func (p Predicate) Apply(part Part) string {
	if p.kind == DEFAULT {
		return p.res
	}
	comp := func(a, b int) bool { return a < b }
	if p.kind == GT {
		comp = func(a, b int) bool { return a > b }
	}
	if comp(part[p.subject], p.level) {
		return p.res
	}
	return ""
}

func parsePredicate(predStr string) (Predicate, error) {
	res := Predicate{}
	predParts := strings.Split(predStr, ":")
	if len(predParts) != 2 {
		return res, fmt.Errorf("malformed predicate string: %s", predStr)
	}

	res.res = predParts[1]
	parts := []string{}
	if strings.Contains(predParts[0], "<") {
		parts = strings.Split(predParts[0], "<")
		res.kind = LT
	} else {
		parts = strings.Split(predParts[0], ">")
		res.kind = GT
	}

	if len(parts) != 2 {
		return res, fmt.Errorf("invalid predicate string: %s", predStr)
	}

	op, err := strconv.Atoi(parts[1])
	if err != nil {
		return res, err
	}
	res.level = op
	switch parts[0] {
	case "x":
		res.subject = X
	case "m":
		res.subject = M
	case "a":
		res.subject = A
	case "s":
		res.subject = S
	default:
		return res, fmt.Errorf("unknown variable in predicate: %s", parts[0])
	}

	return res, nil
}

func defaultPredicate(label string) Predicate {
	return Predicate{0, DEFAULT, X, label}
}

type Workflow struct {
	label      string
	predicates []Predicate
}

func (w Workflow) Apply(p Part) string {
	res := ""
	for i := 0; i < len(w.predicates) && res == ""; i++ {
		res = w.predicates[i].Apply(p)
	}
	return res
}

func parseWorkflow(line string) (Workflow, error) {
	res := Workflow{}
	labelRest := strings.Split(line, "{")
	if len(labelRest) != 2 {
		return res, fmt.Errorf("invalid workflow: %s", line)
	}

	res.label = labelRest[0]
	// needs at least the closing } and the default state
	if len(labelRest[1]) < 2 {
		return res, fmt.Errorf("invalid rest of workflow: %s", line)
	}

	insides := labelRest[1][:len(labelRest[1])-1]

	preds := strings.Split(insides, ",")
	last := len(preds) - 1

	for i := 0; i < last; i++ {
		pred, err := parsePredicate(preds[i])
		if err != nil {
			return res, err
		}
		res.predicates = append(res.predicates, pred)
	}

	res.predicates = append(res.predicates, defaultPredicate(preds[last]))

	return res, nil
}

type RangeSet []utils.Range

func NewRangeSet() RangeSet {
	return RangeSet{
		utils.NewRange(1, 4000),
		utils.NewRange(1, 4000),
		utils.NewRange(1, 4000),
		utils.NewRange(1, 4000),
	}
}

func (rs RangeSet) Combinations() int {
	return rs[X].Len() * rs[M].Len() * rs[A].Len() * rs[S].Len()
}

func (rs RangeSet) CloneWith(mut func(*RangeSet)) RangeSet {
	nrs := NewRangeSet()
	nrs[X] = rs[X]
	nrs[M] = rs[M]
	nrs[A] = rs[A]
	nrs[S] = rs[S]
	mut(&nrs)
	return nrs
}

func findValidRangeSets(wfs map[string]Workflow, curLabel string, curRangeSet RangeSet) []RangeSet {
	if curLabel == "A" {
		return []RangeSet{curRangeSet}
	}

	if curLabel == "R" {
		return []RangeSet{}
	}

	wf, _ := wfs[curLabel]
	res := []RangeSet{}
	for _, p := range wf.predicates {
		switch p.kind {
		case LT:
			if curRangeSet[p.subject].End() < p.level {
				newSets := findValidRangeSets(wfs, p.res, curRangeSet)
				res = append(res, newSets...)
			} else if curRangeSet[p.subject].Start() < p.level {
				nrs := curRangeSet.CloneWith(func(rs *RangeSet) {
					(*rs)[p.subject].MoveEnd(p.level - 1)
				})
				newSets := findValidRangeSets(wfs, p.res, nrs)
				res = append(res, newSets...)

				curRangeSet = curRangeSet.CloneWith(func(rs *RangeSet) {
					(*rs)[p.subject].MoveStart(p.level)
				})
			}
		case GT:
			if curRangeSet[p.subject].Start() > p.level {
				newSets := findValidRangeSets(wfs, p.res, curRangeSet)
				res = append(res, newSets...)
			} else if curRangeSet[p.subject].End() > p.level {
				nrs := curRangeSet.CloneWith(func(rs *RangeSet) {
					(*rs)[p.subject].MoveStart(p.level + 1)
				})
				newSets := findValidRangeSets(wfs, p.res, nrs)
				res = append(res, newSets...)

				curRangeSet = curRangeSet.CloneWith(func(rs *RangeSet) {
					(*rs)[p.subject].MoveEnd(p.level)
				})
			}
		default:
			newSets := findValidRangeSets(wfs, p.res, curRangeSet)
			res = append(res, newSets...)
		}
	}

	return res
}

func Solution() int {
	lines, err := parse.GetLines("day19/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	workflows := map[string]Workflow{}
	i := 0
	for ; i < len(lines) && lines[i] != ""; i++ {
		wf, err := parseWorkflow(lines[i])
		if err != nil {
			fmt.Printf("Failed to parse workflow: %v\n", err)
			return 1
		}
		workflows[wf.label] = wf
	}

	// skip empty line
	i++
	parts := []Part{}
	for ; i < len(lines); i++ {
		part, err := parsePart(lines[i])
		if err != nil {
			fmt.Printf("Failed to parse part: %v\n", err)
			return 1
		}
		parts = append(parts, part)
	}

	res1 := 0
	for _, part := range parts {
		res := "in"
		for res != "A" && res != "R" {
			res = workflows[res].Apply(part)
		}

		if res == "A" {
			res1 += part.Rating()
		}
	}

	fmt.Printf("Part 1: %d\n", res1)

	crs := NewRangeSet()
	vrs := findValidRangeSets(workflows, "in", crs)
	res2 := 0
	for _, rs := range vrs {
		res2 += rs.Combinations()
	}
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
