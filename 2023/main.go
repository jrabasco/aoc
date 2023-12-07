package main

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/day1"
	"github.com/jrabasco/aoc/2023/day2"
	"github.com/jrabasco/aoc/2023/day3"
	"github.com/jrabasco/aoc/2023/day4"
	"github.com/jrabasco/aoc/2023/day5"
	"github.com/jrabasco/aoc/2023/day6"
	"github.com/jrabasco/aoc/2023/day7"
	"github.com/jrabasco/aoc/2023/framework/grid"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"os"
)

type Commands map[string]func() int

var cmds = Commands{
	"day1": day1.Solution,
	"day2": day2.Solution,
	"day3": day3.Solution,
	"day4": day4.Solution,
	"day5": day5.Solution,
	"day6": day6.Solution,
	"day7": day7.Solution,
}

var tests = Commands{
	"grid":  grid.BasicTest,
	"stack": utils.TestStack,
	"queue": utils.TestQueue,
}

func main() {
	cmd := "all"
	if len(os.Args) < 2 {
		fmt.Println("Running all solutions...")
	} else {
		cmd = os.Args[1]
	}

	if cmd == "test" {
		first := true
		for name, t := range tests {
			if !first {
				fmt.Println()
			}
			fmt.Printf("Running %s test...\n", name)
			res := t()
			if res != 0 {
				fmt.Println("NOT OK")
				os.Exit(res)
			}
			first = false
			fmt.Println("OK")
		}
	} else if cmd == "all" {
		first := true
		for day, sol := range cmds {
			if !first {
				fmt.Println()
			}
			fmt.Printf("Running %s:\n", day)
			res := sol()
			if res != 0 {
				fmt.Println("Failed.")
				os.Exit(res)
			}
			first = false
		}
	} else if fn, ok := cmds[cmd]; ok {
		os.Exit(fn())
	} else {
		fmt.Printf("Invalid day or test: %s\n", cmd)
		os.Exit(1)
	}
}
