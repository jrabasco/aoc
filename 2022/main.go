package main

import (
    "os"
    "fmt"
    "github.com/jrabasco/aoc/2022/framework/grid"
    "github.com/jrabasco/aoc/2022/day1"
    "github.com/jrabasco/aoc/2022/day2"
)

type Command map[string]func() int

var cmds = Command{
    "grid": grid.BasicTest,
    "day1": day1.Solution,
    "day2": day2.Solution,
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Not enough arguments, please specify a day or test.")
        os.Exit(1)
    }

    cmd := os.Args[1]

    if fn, ok := cmds[cmd]; ok {
        os.Exit(fn())
    } else {
        fmt.Printf("Invalid day or test: %s\n", cmd)
        os.Exit(1)
    }
}
