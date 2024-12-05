package day5

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"sort"
	"strconv"
	"strings"
)

// Maps a page number to the page number they prevent
// i.e. if the rule says X|Y then when we see Y we must
// not see X again so the rules would contain Y -> X
type Rules map[int]utils.Set[int]

func px(lines []string) (int, int) {
	rules := Rules{}
	i := 0
	for ; lines[i] != ""; i++ {
		parts := strings.Split(lines[i], "|")
		X, _ := strconv.Atoi(parts[0])
		Y, _ := strconv.Atoi(parts[1])

		if _, found := rules[Y]; !found {
			rules[Y] = utils.NewSet[int]()
		}
		s := rules[Y]
		s.Add(X)
		rules[Y] = s
	}
	i++

	manuals := [][]int{}

	for ; i < len(lines); i++ {
		pages := []int{}
		parts := strings.Split(lines[i], ",")
		for _, part := range parts {
			p, _ := strconv.Atoi(part)
			pages = append(pages, p)
		}
		manuals = append(manuals, pages)
	}

	res1 := 0
	res2 := 0
	for _, pages := range manuals {
		ok := true
		banned := utils.NewSet[int]()
		for _, page := range pages {
			if banned.Contains(page) {
				ok = false
				break
			}
			if b, found := rules[page]; found {
				banned.AddSet(b)
			}
		}

		l := len(pages)
		if ok {
			res1 += pages[l/2]
		} else {
			sort.Slice(pages, func(i, j int) bool {
				X := pages[i]
				Y := pages[j]
				if set, found := rules[Y]; found {
					return set.Contains(X)
				}
				return false
			})
			res2 += pages[l/2]
		}
	}

	return res1, res2
}

func Solution() int {
	parsed, err := parse.GetLines("day5/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	p1, p2 := px(parsed)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	return 0
}
