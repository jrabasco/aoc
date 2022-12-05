package main

import (
    "os"
    "fmt"
    "github.com/jrabasco/aoc/2022/day1"
)

type Solutions map[string]func() int

var solutions = Solutions{
    "day1": day1.Solution,
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Not enough arguments, please specify a day.")
        os.Exit(1)
    }

    day := os.Args[1]

    if fn, ok := solutions[day]; ok {
        os.Exit(fn())
    } else {
        fmt.Printf("Invalid day: %s\n", day)
        os.Exit(1)
    }
}
