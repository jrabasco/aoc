package main

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/day1"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"os"
)

type Command map[string]func(string) int

var cmds = Command{
	"grid":  grid.BasicTest,
	"stack": utils.TestStack,
	"queue": utils.TestQueue,
	"day1":  day1.Solution,
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
