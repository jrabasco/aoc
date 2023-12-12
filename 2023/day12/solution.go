package day12

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"strconv"
	"strings"
)

type Record struct {
	status         string
	groups         []int
	unfoldedStatus string
	unfoldedGroups []int
}

func parseRecord(line string) (Record, error) {
	rec := Record{}
	stGr := strings.Fields(line)
	if len(stGr) != 2 {
		return rec, fmt.Errorf("invalid line: %s", line)
	}

	rec.status = stGr[0]

	groupsStr := strings.Split(stGr[1], ",")

	rec.groups = []int{}
	for _, grStr := range groupsStr {
		gr, err := strconv.Atoi(grStr)

		if err != nil {
			return rec, err
		}

		rec.groups = append(rec.groups, gr)
	}
	rec.unfoldedStatus = rec.status
	rec.unfoldedGroups = rec.groups
	for i := 0; i < 4; i++ {
		rec.unfoldedStatus += "?" + rec.status
		rec.unfoldedGroups = append(rec.unfoldedGroups, rec.groups...)
	}
	return rec, nil
}

type Key struct {
	i int
	j int
}

type Cache map[Key]int

func recurse(i, j int, status string, groups []int, cache Cache) int {
	if i >= len(status) {
		if j >= len(groups) {
			// done with all groups
			return 1
		}
		// at least one group not completed
		return 0
	}

	key := Key{i, j}
	if res, ok := cache[key]; ok {
		return res
	}

	c := status[i]

	// ignore char and carry on
	if c == '.' {
		return recurse(i+1, j, status, groups, cache)
	}

	res := 0
	// try to put a dot and move on
	if c == '?' {
		res += recurse(i+1, j, status, groups, cache)
	}

	// here we examine the case where ? becomes # or where we find a #
	if j < len(groups) {
		// try and make a run of # that is the size of the next expected group
		count := 0
		for k := i; k < len(status); k++ {
			// - If count is larger than the group size then it's not going to
			//   work (means we found many # in a row)
			// - If we find a . then it's the end of the run
			// - If the count is the group size and the next is a ? then we can
			//   make it a group (assume a .). This would also work if the next
			//   is a . but that's covered by the previous condition
			if count > groups[j] || status[k] == '.' || count == groups[j] && status[k] == '?' {
				break
			}
			count += 1
		}

		// if we managed to create a group of the right size
		// this can happen if:
		// - we reached a . or a ? and happen to have the right size
		// - we reached the end of the status
		if count == groups[j] {
			if i+count < len(status) {
				// not at the end, meaning the next one is a . or a ? which
				// we should skip, we also record the fact that we found
				// a group by going to group j+1
				res += recurse(i+count+1, j+1, status, groups, cache)
			} else {
				// we reached the end! we can perform the same check as the
				// start of this function or simply do a recursive call
				res += recurse(i+count, j+1, status, groups, cache)
			}
		}
	}
	cache[key] = res
	return res
}

func Solution() int {
	records, err := parse.GetLinesAs[Record]("day12/input.txt", parseRecord)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	res := 0
	res2 := 0
	for _, rec := range records {
		cache := Cache{}
		res += recurse(0, 0, rec.status, rec.groups, cache)
		cache2 := Cache{}
		res2 += recurse(0, 0, rec.unfoldedStatus, rec.unfoldedGroups, cache2)
	}
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", res2)
	return 0
}
