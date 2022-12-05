package main

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/day1"
	"github.com/jrabasco/aoc/2022/day2"
	"github.com/jrabasco/aoc/2022/day3"
	"github.com/jrabasco/aoc/2022/day4"
	"github.com/jrabasco/aoc/2022/day5"
	"github.com/jrabasco/aoc/2022/framework/grid"
	"github.com/jrabasco/aoc/2022/framework/utils"
	"os"
)

type Command map[string]func(string) int

var cmds = Command{
	"grid":  grid.BasicTest,
	"stack": utils.TestStack,
	"day1":  day1.Solution,
	"day2":  day2.Solution,
	"day3":  day3.Solution,
	"day4":  day4.Solution,
	"day5":  day5.Solution,
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
