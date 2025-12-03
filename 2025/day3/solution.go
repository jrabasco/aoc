package day3

import (
	"fmt"
	"github.com/jrabasco/aoc/2025/framework/parse"
)

func p1(banks [][]int) int {
	tot := 0
	for _, bank := range banks {
		lmax := 0
		lidx := -1
		// the left number can be at most the one before last
		for i := 0; i < len(bank)-1; i++ {
			if bank[i] > lmax {
				lmax = bank[i]
				lidx = i
			}
		}
		rmax := 0
		// the right number can be at most the one after the prev max
		for i := lidx + 1; i < len(bank); i++ {
			if bank[i] > rmax {
				rmax = bank[i]
			}
		}
		tot += lmax*10 + rmax
	}
	return tot
}

func p2(banks [][]int) int {
	// we generalise the idea from p1 for 12 numbers to find
	tot := 0
	for _, bank := range banks {
		offset := 0
		res_array := []int{}
		for n := 0; n < 12; n++ {
			// the exploration space is from offset (which is one after the
			// last found number) to the end of the array minus how many number
			// we have to find
			mx := 0
			for i := offset; i < len(bank)-(12-n-1); i++ {
				if bank[i] > mx {
					mx = bank[i]
					offset = i + 1
				}
			}
			res_array = append(res_array, mx)
		}
		exp := 1
		for i := len(res_array) - 1; i >= 0; i-- {
			tot += res_array[i] * exp
			exp *= 10
		}
	}
	return tot
}

func strToIntArray(line string) ([]int, error) {
	runes := []rune(line)
	res := []int{}
	for _, r := range runes {
		if r < '0' || r > '9' {
			return res, fmt.Errorf("not a valid digit: %s", r)
		}
		res = append(res, int(r-'0'))
	}
	return res, nil
}

func Solution() int {
	parsed, err := parse.GetLinesAs[[]int]("day3/input.txt", strToIntArray)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
