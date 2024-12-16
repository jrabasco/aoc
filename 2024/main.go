package main

import (
	"flag"
	"fmt"
	"github.com/jrabasco/aoc/2024/day1"
	"github.com/jrabasco/aoc/2024/day10"
	"github.com/jrabasco/aoc/2024/day11"
	"github.com/jrabasco/aoc/2024/day12"
	"github.com/jrabasco/aoc/2024/day13"
	"github.com/jrabasco/aoc/2024/day14"
	"github.com/jrabasco/aoc/2024/day15"
	"github.com/jrabasco/aoc/2024/day16"
	"github.com/jrabasco/aoc/2024/day2"
	"github.com/jrabasco/aoc/2024/day3"
	"github.com/jrabasco/aoc/2024/day4"
	"github.com/jrabasco/aoc/2024/day5"
	"github.com/jrabasco/aoc/2024/day6"
	"github.com/jrabasco/aoc/2024/day7"
	"github.com/jrabasco/aoc/2024/day8"
	"github.com/jrabasco/aoc/2024/day9"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"os"
	"runtime/pprof"
	"time"
)

type Commands map[string]func() int

var cmds = Commands{
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
	"day14": day14.Solution,
	"day15": day15.Solution,
	"day16": day16.Solution,
}

var tests = Commands{
	"grid":   grid.BasicTest,
	"stack":  utils.TestStack,
	"queue":  utils.TestQueue,
	"pqueue": utils.TestPriorityQueue,
	"heap":   utils.TestHeap,
}

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var day = flag.Int("day", 0, "run solution for one day, default to all when not specified")
	var test = flag.Bool("test", false, "run all tests (supercedes -day)")
	var doTime = flag.Bool("time", false, "time whatever is executed by runnint it 10 times")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Printf("could not create profile file: %v\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	loops := 1
	start := time.Now()
	if *doTime {
		loops = 10
	}

	for l := 0; l < loops; l++ {
		if *test {
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
		} else if *day == 0 {
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
		} else if fn, ok := cmds[fmt.Sprintf("day%d", *day)]; ok {
			retVal := fn()
			if retVal != 0 {
				os.Exit(retVal)
			}
		} else {
			fmt.Printf("Invalid day or test: %s\n", *day)
			os.Exit(1)
		}
	}
	elapsed := time.Since(start)
	if *doTime {
		fmt.Printf("Average time: %s\n", elapsed/time.Duration(loops))
	}
}
