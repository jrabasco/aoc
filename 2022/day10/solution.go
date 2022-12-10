package day10

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
	"strings"
)

type Noop struct{}

type AddX struct {
	x int
}

type Instruction interface{}

type CRT [6][40]rune

func NewCRT() CRT {
	return [6][40]rune{
		[40]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[40]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[40]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[40]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[40]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[40]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
}

func (crt *CRT) PutPixel(pos int) error {
	row := pos / 40
	if row < 0 || row > 5 {
		return fmt.Errorf("invalid position: %d", pos)
	}
	col := crt.GetCol(pos)
	if col < 0 {
		return fmt.Errorf("invalid position: %d", pos)
	}
	crt[row][col] = '#'
	return nil
}

func (crt CRT) GetCol(pos int) int {
	return pos % 40
}

func (crt CRT) String() string {
	parts := []string{}
	for i, l := range crt {
		parts = append(parts, "")
		for _, r := range l {
			parts[i] += string(r)
		}
	}
	return strings.Join(parts, "\n")
}

type CPU struct {
	register int
	cycle    int
	program  []Instruction
	cur      Instruction
	crt      CRT
}

func NewCPU() CPU {
	return CPU{1, 1, []Instruction{}, Noop{}, NewCRT()}
}

func (cpu *CPU) LoadInstructions(lines []string) error {
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 1 {
			if parts[0] != "noop" {
				return fmt.Errorf("invalid instruction: '%s'", line)
			}
			cpu.program = append(cpu.program, Noop{})
			continue
		}

		if len(parts) != 2 || parts[0] != "addx" {
			return fmt.Errorf("invalid instruction: '%s'", line)
		}

		x, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid instruction (invalid int): '%s'", line)
		}

		cpu.program = append(cpu.program, AddX{x})
	}

	return nil
}

func (cpu *CPU) Tick() error {
	// CRT counts from 0
	c := cpu.crt.GetCol(cpu.cycle - 1)
	if cpu.register-1 <= c && c <= cpu.register+1 {
		err := cpu.crt.PutPixel(cpu.cycle - 1)
		if err != nil {
			return err
		}
	}
	// are we still finishing the previous instruction?
	switch cur := cpu.cur.(type) {
	case Noop:
		break
	case AddX:
		cpu.cycle += 1
		cpu.register += cur.x
		cpu.cur = Noop{}
		return nil
	default:
		return fmt.Errorf("invalid instruction type : %T", cur)
	}

	cpu.cycle += 1
	cpu.cur = cpu.program[0]
	cpu.program = cpu.program[1:]
	return nil
}

func (cpu CPU) Done() bool {
	return len(cpu.program) == 0
}

func (cpu CPU) SignalStrength() int {
	return cpu.register * cpu.cycle
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day10/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}
	cpu := NewCPU()
	err = cpu.LoadInstructions(parsed)
	if err != nil {
		fmt.Printf("Could not load instructions: %v\n", err)
		return 1
	}

	strengths := []int{}
	for !cpu.Done() {
		if (cpu.cycle-20)%40 == 0 {
			strengths = append(strengths, cpu.SignalStrength())
		}
		cpu.Tick()
	}
	res := 0
	for _, s := range strengths {
		res += s
	}
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: \n%s\n", cpu.crt)
	return 0
}

func Solution(part string) int {
	return solvePart(part)
}
