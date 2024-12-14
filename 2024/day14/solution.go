package day14

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/grid"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"strconv"
	"strings"
)

const WIDTH = 101
const HEIGHT = 103

type Robot struct {
	p grid.Point
	v grid.Vector
}

func parse2Ints(str string) (int, int, error) {
	parts := strings.Split(str, "=")
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("invalid int pair: %s", str)
	}
	xyParts := strings.Split(parts[1], ",")
	if len(xyParts) < 2 {
		return 0, 0, fmt.Errorf("invalid int pair: %s", str)
	}
	x, err := strconv.Atoi(xyParts[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(xyParts[1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func parseRobot(line string) (Robot, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return Robot{}, fmt.Errorf("invalid robot spec: %s", line)
	}
	pStr := parts[0]
	px, py, err := parse2Ints(pStr)
	if err != nil {
		return Robot{}, err
	}
	vStr := parts[1]
	vx, vy, err := parse2Ints(vStr)
	if err != nil {
		return Robot{}, err
	}
	return Robot{grid.Point{px, py}, grid.Vector{vx, vy}}, nil
}

func printRobots(robots []Robot) {
	res := [][]rune{}
	for i := 0; i < HEIGHT; i++ {
		row := []rune{}
		for j := 0; j < WIDTH; j++ {
			row = append(row, '.')
		}
		res = append(res, row)
	}
	for _, r := range robots {
		res[r.p.Y][r.p.X] = '#'
	}
	for _, row := range res {
		fmt.Println(string(row))
	}
}

func getRobotsAt(time int, robots []Robot) []Robot {
	res := []Robot{}
	for _, robot := range robots {
		finalX := (robot.p.X + time*robot.v.X) % WIDTH
		if finalX < 0 {
			finalX += WIDTH
		}
		finalY := (robot.p.Y + time*robot.v.Y) % HEIGHT
		if finalY < 0 {
			finalY += HEIGHT
		}
		res = append(res, Robot{grid.Point{finalX, finalY}, robot.v})
	}
	return res
}

func getEntropy1At(time int, robots []Robot) int {
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for _, robot := range robots {
		finalX := (robot.p.X + time*robot.v.X) % WIDTH
		if finalX < 0 {
			finalX += WIDTH
		}
		finalY := (robot.p.Y + time*robot.v.Y) % HEIGHT
		if finalY < 0 {
			finalY += HEIGHT
		}
		if finalX < WIDTH/2 {
			if finalY < HEIGHT/2 {
				q1 += 1
			} else if finalY > HEIGHT/2 {
				q2 += 1
			}
		} else if finalX > WIDTH/2 {
			if finalY < HEIGHT/2 {
				q3 += 1
			} else if finalY > HEIGHT/2 {
				q4 += 1
			}
		}
	}
	return q1 * q2 * q3 * q4
}

func getEntropy2At(time int, robots []Robot) int {
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	q5 := 0
	q6 := 0
	q7 := 0
	q8 := 0
	q9 := 0
	for _, robot := range robots {
		finalX := (robot.p.X + time*robot.v.X) % WIDTH
		if finalX < 0 {
			finalX += WIDTH
		}
		finalY := (robot.p.Y + time*robot.v.Y) % HEIGHT
		if finalY < 0 {
			finalY += HEIGHT
		}
		if finalX < WIDTH/3 {
			if finalY < HEIGHT/3 {
				q1 += 1
			} else if finalY > HEIGHT/3 && finalY < (2*HEIGHT)/3 {
				q2 += 1
			} else if finalY > (2*HEIGHT)/3 {
				q3 += 1
			}
		} else if finalX > WIDTH/3 && finalX < (2*WIDTH)/3 {
			if finalY < HEIGHT/3 {
				q4 += 1
			} else if finalY > HEIGHT/3 && finalY < (2*HEIGHT)/3 {
				q5 += 1
			} else if finalY > (2*HEIGHT)/3 {
				q6 += 1
			}
		} else if finalX > (2*WIDTH)/3 {
			if finalY < HEIGHT/3 {
				q7 += 1
			} else if finalY > HEIGHT/3 && finalY < (2*HEIGHT)/3 {
				q8 += 1
			} else if finalY > (2*HEIGHT)/3 {
				q9 += 1
			}
		}
	}
	return q1 * q2 * q3 * q4 * q5 * q6 * q7 * q8 * q9
}

func p1(robots []Robot) int {
	return getEntropy1At(100, robots)
}

func p2(robots []Robot) int {
	minentropy := getEntropy2At(0, robots)
	minidx := 0
	for i := 1; i < 10000; i++ {
		entropy := getEntropy2At(i, robots)
		if entropy < minentropy {
			minentropy = entropy
			minidx = i
		}
	}
	printRobots(getRobotsAt(minidx, robots))
	return minidx
}

func Solution() int {
	robots, err := parse.GetLinesAs[Robot]("day14/input.txt", parseRobot)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(robots))
	fmt.Printf("Part 2: %d\n", p2(robots))
	return 0
}
