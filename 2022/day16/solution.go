package day16

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"strconv"
	"strings"
)

type Target struct {
	valve *Valve
	cost  int
}

type Valve struct {
	tag   string
	rate  int
	paths map[string]*Target
}

func NewValve(tag string, rate int) Valve {
	return Valve{tag, rate, map[string]*Target{}}
}

type Map map[string]*Valve

func NewMap() Map {
	return Map{}
}

func parseValve(line string) (Valve, error) {
	parts := strings.Split(line, ";")
	if len(parts) != 2 {
		return NewValve("", 0), fmt.Errorf("could not parse valve from line \"%s\"", line)
	}
	valveStr := parts[0]
	tag := valveStr[6:8]
	rateStr := valveStr[23:]
	rate, err := strconv.Atoi(rateStr)

	if err != nil {
		return NewValve("", 0), fmt.Errorf("invalide rate %s: %v", rateStr, err)
	}
	return NewValve(tag, rate), nil
}

func parseDests(line string) (string, []string, error) {
	parts := strings.Split(line, "; ")
	if len(parts) != 2 {
		return "", []string{}, fmt.Errorf("could not parse dests from line \"%s\"", line)
	}
	tag := parts[0][6:8]
	destsWords := strings.Fields(parts[1])

	blabla := true
	dests := []string{}

	for _, w := range destsWords {
		if w == "valves" || w == "valve" {
			blabla = false
			continue
		}

		if blabla {
			continue
		}

		if len(w) == 3 {
			w = w[:2]
		}
		dests = append(dests, w)
	}

	return tag, dests, nil
}

func addPathToValve(v *Valve, to *Valve, cost int) {
	if t, ok := v.paths[to.tag]; ok {
		if cost < t.cost {
			t.cost = cost
		}
		return
	}
	v.paths[to.tag] = &Target{to, cost}
}

func addPath(m *Map, from, to string) {
	fromV := (*m)[from]
	toV := (*m)[to]

	addPathToValve(fromV, toV, 1)

	for _, t := range toV.paths {
		if t.valve.tag == from {
			continue
		}
		addPathToValve(fromV, t.valve, 1+t.cost)
	}

	for k, v := range *m {
		if k == to {
			continue
		}
		nCost := -1

		for _, t := range v.paths {
			if t.valve.tag == from {
				nCost = t.cost + 1
			}
		}

		if nCost == -1 {
			continue
		}

		addPathToValve(v, toV, nCost)
	}
}

func parseMap(lines []string) (Map, error) {
	m := NewMap()
	for _, line := range lines {
		v, err := parseValve(line)

		if err != nil {
			return m, err
		}

		m[v.tag] = &v
	}

	for _, line := range lines {
		tag, dests, err := parseDests(line)

		if err != nil {
			return m, err
		}
		for _, dest := range dests {
			addPath(&m, tag, dest)
		}
	}
	return m, nil
}

func findBestPath(m *Map, viValves []*Valve, cur *Valve, visited utils.Set[string], rate, minutes, acc int) (int, utils.Set[string]) {
	if minutes == 0 {
		return acc, visited
	}

	reachable := []*Valve{}
	for _, v := range viValves {
		if visited.Contains(v.tag) {
			continue
		}
		if p, ok := cur.paths[v.tag]; ok {
			if p.cost <= minutes {
				reachable = append(reachable, v)
			}
		}
	}

	if len(reachable) == 0 {
		return acc + minutes*rate, visited
	}

	max := 0
	nVis := visited
	for _, v := range reachable {
		visited.Add(v.tag)
		nrate := rate + v.rate
		t := cur.paths[v.tag]
		// cost to go and open
		spentMins := t.cost + 1
		nMinutes := minutes - spentMins
		nAcc := acc + spentMins*rate
		att, vis := findBestPath(m, viValves, v, visited, nrate, nMinutes, nAcc)
		if att > max {
			max = att
			nVis = vis
		}
		visited.Remove(v.tag)
	}

	nVis.Add(cur.tag)

	return max, nVis
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day16/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	m, err := parseMap(parsed)
	if err != nil {
		fmt.Printf("Failed to create map: %v\n", err)
		return 1
	}

	viValves := []*Valve{}
	for _, v := range m {
		//		fmt.Printf("%s: %d|", k, v.rate)
		//		paths := []string{}
		//		for _, p := range v.paths {
		//			paths = append(paths, fmt.Sprintf("(%s:%d)", p.valve.tag, p.cost))
		//		}
		//		fmt.Println(strings.Join(paths, ","))

		if v.rate > 0 {
			viValves = append(viValves, v)
		}
	}

	if part == "1" {

		res, vis := findBestPath(&m, viValves, m["AA"], utils.Set[string]{}, 0, 30, 0)
		fmt.Println(vis)
		fmt.Printf("Part %s: %v\n", part, res)
	} else {
		res, vis := findBestPath(&m, viValves, m["AA"], utils.Set[string]{}, 0, 30, 0)
		fmt.Println(vis)
		fmt.Printf("Part %s: %v\n", part, res)
	}
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
