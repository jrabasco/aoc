package day7

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"sort"
	"strconv"
	"strings"
)

const FIVE = 6
const FOUR = 5
const FULL = 4
const THREE = 3
const TWOPAIR = 2
const PAIR = 1
const HIGH = 0

var valueMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	bid     int
	hand    string
	values  []int
	valuesJ []int
	power   int
	powerJ  int
}

func (h Hand) Less(o Hand) bool {
	if h.power != o.power {
		return h.power < o.power
	}

	for i, val := range h.values {
		if val != o.values[i] {
			return val < o.values[i]
		}
	}

	return false
}

func (h Hand) LessJ(o Hand) bool {
	if h.powerJ != o.powerJ {
		return h.powerJ < o.powerJ
	}

	for i, val := range h.valuesJ {
		if val != o.valuesJ[i] {
			return val < o.valuesJ[i]
		}
	}

	return false
}

func handPower(counts map[rune]int) int {
	// default power level
	res := HIGH
	triplets := 0
	pairs := 0
	for _, count := range counts {
		if count == 5 {
			res = FIVE
			break
		}

		if count == 4 {
			res = FOUR
			break
		}

		if count == 3 {
			triplets += 1
		}

		if count == 2 {
			pairs += 1
		}
	}

	if triplets == 1 {
		if pairs == 1 {
			res = FULL
		} else {
			res = THREE
		}
	} else if pairs != 0 {
		if pairs == 2 {
			res = TWOPAIR
		} else {
			res = PAIR
		}
	}
	return res
}

func parseHand(line string) (Hand, error) {
	res := Hand{}
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return res, fmt.Errorf("malformed line: %s", line)
	}

	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return res, err
	}
	res.bid = bid
	res.hand = parts[0]
	counts := map[rune]int{}
	for _, card := range parts[0] {
		value, ok := valueMap[card]

		if !ok {
			return res, fmt.Errorf("unknown card: %s", string(card))
		}

		res.values = append(res.values, value)

		if card == 'J' {
			res.valuesJ = append(res.valuesJ, 1)
		} else {
			res.valuesJ = append(res.valuesJ, value)
		}

		// add to counts
		if _, exists := counts[card]; !exists {
			counts[card] = 0
		}

		counts[card] += 1
	}

	res.power = handPower(counts)
	res.powerJ = res.power

	jcount := counts['J']

	// if we have 5 Js it's obviously a 5 of a kind situation
	// if we have 4 Js they can take the value of the last card and give a 5 of
	// a kind
	if jcount == 5 || jcount == 4 {
		res.powerJ = FIVE
	}

	// XXJJJ becomes XXXXX if J=X
	if jcount == 3 && len(counts) == 2 {
		res.powerJ = FIVE
	}

	// XYJJJ can become XYYYY or YXXXX but cannot do a 5 of a kind
	if jcount == 3 && len(counts) == 3 {
		res.powerJ = FOUR
	}

	// try all possibilities, can't be bothered to find nicer rules as it's
	// going to be at most 9 attempts (both Js can take any of the values for
	// the other cards)
	if jcount == 2 {
		// we are ignoring all cases where we keep a J as is because if we
		// change only one J we can always trivially improve the power by
		// changing the other J to the same card

		// we're going to use this map for potential J values
		tmpCounts := map[rune]int{}
		for card, count := range counts {
			if card == 'J' {
				continue
			}
			tmpCounts[card] = count
		}

		for cardo, _ := range counts {
			if cardo == 'J' {
				continue
			}
			tmpCounts[cardo] += 1
			for cardi, _ := range counts {
				if cardi == 'J' {
					continue
				}
				tmpCounts[cardi] += 1
				nScore := handPower(tmpCounts)
				if nScore > res.powerJ {
					res.powerJ = nScore
				}
				tmpCounts[cardi] -= 1
			}
			tmpCounts[cardo] -= 1
		}
	}

	// same things as for jcount == 2 but with only one loop
	if jcount == 1 {
		tmpCounts := map[rune]int{}
		for card, count := range counts {
			if card == 'J' {
				continue
			}
			tmpCounts[card] = count
		}

		for card, _ := range counts {
			if card == 'J' {
				continue
			}
			tmpCounts[card] += 1
			nScore := handPower(tmpCounts)
			if nScore > res.powerJ {
				res.powerJ = nScore
			}
			tmpCounts[card] -= 1
		}
	}

	return res, nil
}

func Solution() int {
	hands, err := parse.GetLinesAs[Hand]("day7/input.txt", parseHand)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	sort.Slice(hands, func(i, j int) bool { return hands[i].Less(hands[j]) })
	res1 := 0
	for i, h := range hands {
		res1 += (i + 1) * h.bid
	}
	fmt.Printf("Part 1: %d\n", res1)

	sort.Slice(hands, func(i, j int) bool { return hands[i].LessJ(hands[j]) })
	res2 := 0
	for i, h := range hands {
		res2 += (i + 1) * h.bid
	}
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
