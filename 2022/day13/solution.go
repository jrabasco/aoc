package day13

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"sort"
	"strconv"
	"strings"
)

type Packet interface {
	// returns 1 if self < p, -1 if p < self, 0 otherwise
	Compare(p Packet) int
	MakeDivider()
	IsDivider() bool
	String() string
}

type I struct {
	data      int
	isDivider bool
}

func NewI(data int) *I {
	return &I{data, false}
}

func (i I) Compare(p Packet) int {
	switch pack := p.(type) {
	case *I:
		if i.data < pack.data {
			return 1
		} else if i.data > pack.data {
			return -1
		} else {
			return 0
		}
	case *L:
		lI := NewL([]Packet{&i})
		return lI.Compare(pack)
	default:
		return 0
	}
}

func (i I) String() string {
	return fmt.Sprintf("%d", i.data)
}

func (i *I) MakeDivider() {
	i.isDivider = true
}

func (i I) IsDivider() bool {
	return i.isDivider
}

type L struct {
	data      []Packet
	isDivider bool
}

func NewL(data []Packet) *L {
	return &L{data, false}
}

func (l L) Compare(p Packet) int {
	switch pack := p.(type) {
	case *I:
		lP := NewL([]Packet{pack})
		return l.Compare(lP)
	case *L:
		for i, e := range l.data {
			if i >= len(pack.data) {
				// l is longer than pack
				return -1
			}

			res := e.Compare(pack.data[i])
			if res != 0 {
				return res
			}
		}
		// check if pack had more elements left
		if len(pack.data) > len(l.data) {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func (l L) String() string {
	resP := []string{}
	for _, p := range l.data {
		pS := p.String()
		if p.IsDivider() {
			pS += "*"
		}
		resP = append(resP, pS)
	}
	res := "["
	res += strings.Join(resP, ",")
	res += "]"
	return res
}

func (l *L) MakeDivider() {
	l.isDivider = true
}

func (l L) IsDivider() bool {
	return l.isDivider
}

func (l *L) Sort() {
	sort.Slice(l.data, func(i, j int) bool {
		return l.data[i].Compare(l.data[j]) == 1
	})
}

type Block struct {
	left  Packet
	right Packet
}

func NewBlock(left, right Packet) Block {
	return Block{left, right}
}

func (b Block) IsOrdered() bool {
	return b.left.Compare(b.right) == 1
}

func parsePacket(line string) (Packet, int, error) {
	var zero *I
	l := len(line)
	if l < 2 {
		return zero, 0, fmt.Errorf("empty line is not a packet")
	}

	if line[0] != '[' {
		return zero, 0, fmt.Errorf("invalid packet: %s", line)
	}

	var acc []Packet
	for i := 1; i < l; i++ {
		c := line[i]
		if c == ',' || c == ' ' {
			continue
		}
		if c == ']' {
			return NewL(acc), i, nil
		}

		if c != '[' {
			nbS := ""
			for c != ',' && c != ']' && i < l {
				nbS += string(c)
				i++
				c = line[i]
			}
			// avoid discarding the character coming right after the number
			i--
			nb, err := strconv.Atoi(nbS)
			if err != nil {
				return zero, i, fmt.Errorf("could not parse number: %v", err)
			}
			acc = append(acc, NewI(nb))
			continue
		}

		if c == '[' {
			child, delta, err := parsePacket(line[i:])
			if err != nil {
				return zero, i + delta, err
			}
			acc = append(acc, child)
			i += delta
		}

	}

	return NewL(acc), l, nil
}

func getBlocks(lines []string) ([]Block, error) {
	acc := []Block{}
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}

		if i > len(lines)-2 {
			return acc, fmt.Errorf("not enough lines to form a block")
		}

		left, _, err := parsePacket(lines[i])
		if err != nil {
			return acc, err
		}

		right, _, err := parsePacket(lines[i+1])
		if err != nil {
			return acc, err
		}
		i++

		acc = append(acc, NewBlock(left, right))
	}

	return acc, nil
}

func getTransmission(lines []string) (L, error) {
	var zero L
	filtered := []string{}
	for _, l := range lines {
		if len(l) != 0 {
			filtered = append(filtered, l)
		}
	}
	asOne := fmt.Sprintf("[%s]", strings.Join(filtered, ","))
	p, _, err := parsePacket(asOne)
	if err != nil {
		return zero, fmt.Errorf("could not parse transmission: %v", err)
	}

	transm, ok := p.(*L)
	if !ok {
		return zero, fmt.Errorf("transmission did not return a list")
	}
	div1, _, _ := parsePacket("[[2]]")
	div1.MakeDivider()
	transm.data = append(transm.data, div1)
	div2, _, _ := parsePacket("[[6]]")
	div2.MakeDivider()
	transm.data = append(transm.data, div2)
	return *transm, nil
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day13/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	res := 0

	if part == "1" {
		blocks, err := getBlocks(parsed)
		if err != nil {
			fmt.Printf("Could not parse blocks: %v\n", err)
			return 1
		}
		for i, b := range blocks {
			if b.IsOrdered() {
				res += i + 1
			}
		}
	} else {
		transm, err := getTransmission(parsed)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		transm.Sort()

		res = 1
		for i, p := range transm.data {
			if p.IsDivider() {
				res *= i + 1
			}
		}
	}

	fmt.Printf("Part %s: %d\n", part, res)
	return 0
}

func Solution(part string) int {
	if part != "1" && part != "2" {
		p1 := solvePart("1")
		if p1 != 0 {
			return p1
		}
		return solvePart("2")
	} else {
		return solvePart(part)
	}
}
