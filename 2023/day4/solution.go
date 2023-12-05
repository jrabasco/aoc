package day4

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	winning utils.Set[int]
	have    utils.Set[int]
}

func (c Card) wins() int {
	return len(c.winning.Intersect(c.have))
}

func (c Card) value() int {
	ln := c.wins()
	if ln == 0 {
		return 0
	}

	return int(math.Pow(2, float64(ln-1)))
}

func parseCard(line string) (Card, error) {
	card := Card{utils.Set[int]{}, utils.Set[int]{}}
	discard := strings.Split(line, ": ")
	if len(discard) != 2 {
		return card, fmt.Errorf("malformed line: %s", line)
	}

	parts := strings.Split(discard[1], " | ")
	if len(parts) != 2 {
		return card, fmt.Errorf("malformed line: %s", line)
	}

	winningStr := strings.Split(parts[0], " ")
	for _, nbStr := range winningStr {
		if nbStr == "" {
			continue
		}
		nb, err := strconv.Atoi(nbStr)
		if err != nil {
			return card, err
		}
		card.winning.Add(nb)
	}

	haveStr := strings.Split(parts[1], " ")
	for _, nbStr := range haveStr {
		if nbStr == "" {
			continue
		}
		nb, err := strconv.Atoi(nbStr)
		if err != nil {
			return card, err
		}
		card.have.Add(nb)
	}

	return card, nil
}

func Solution() int {
	cards, err := parse.GetLinesAs[Card]("day4/input.txt", parseCard)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	res := 0
	for _, card := range cards {
		res += card.value()
	}
	fmt.Printf("Part 1: %d\n", res)
	res = 0
	copies := make([]int, len(cards))
	for i, card := range cards {
		wins := card.wins()
		adds := copies[i] + 1
		for j := i + 1; j <= i+wins; j++ {
			copies[j] += adds
		}
	}

	for _, cp := range copies {
		res += cp + 1
	}
	fmt.Printf("Part 2: %d\n", res)
	return 0
}
