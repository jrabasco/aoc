package main

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/day1"
	"github.com/jrabasco/aoc/2022/day10"
	"github.com/jrabasco/aoc/2022/day11"
	"github.com/jrabasco/aoc/2022/day12"
	"github.com/jrabasco/aoc/2022/day13"
	"github.com/jrabasco/aoc/2022/day2"
	"github.com/jrabasco/aoc/2022/day3"
	"github.com/jrabasco/aoc/2022/day4"
	"github.com/jrabasco/aoc/2022/day5"
	"github.com/jrabasco/aoc/2022/day6"
	"github.com/jrabasco/aoc/2022/day7"
	"github.com/jrabasco/aoc/2022/day8"
	"github.com/jrabasco/aoc/2022/day9"
	"github.com/jrabasco/aoc/2022/framework/grid"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"os"
)

type Command map[string]func(string) int

var cmds = Command{
	"grid":  grid.BasicTest,
	"stack": utils.TestStack,
	"queue": utils.TestQueue,
	"day1":  day1.Solution,
	"day2":  day2.Solution,
	"day3":  day3.Solution,
	"day4":  day4.Solution,
	"day5":  day5.Solution,
	"day6":  day6.Solution,
	"day7":  day7.Solution,
	"day8":  day8.Solution,
	"day9":  day9.Solution,
	"day10": day10.Solution,
	"day11": day11.Solution,
	"day12": day12.Solution,
	"day13": day13.Solution,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments, please specify a day or test.")
		os.Exit(1)
	}

	cmd := os.Args[1]
	part := "all"

	if len(os.Args) > 2 {
		part = os.Args[2]
	}

	if fn, ok := cmds[cmd]; ok {
		os.Exit(fn(part))
	} else {
		fmt.Printf("Invalid day or test: %s\n", cmd)
		os.Exit(1)
	}
}
