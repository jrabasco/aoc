package day7

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
	"strings"
)

type INode struct {
	name     string
	size     int
	children []*INode
	parent   *INode
}

func NewINode(name string, size int) INode {
	return INode{name, size, []*INode{}, nil}
}

func (in *INode) AddChild(name string, size int) {
	nin := NewINode(name, size)
	nin.parent = in
	in.children = append(in.children, &nin)
}

func (in INode) IsDir() bool {
	return len(in.children) > 0
}

func (in INode) Stringd(depth int) string {
	res := ""
	for i := 0; i < depth; i++ {
		res += " "
	}

	res += fmt.Sprintf("- %s", in.name)

	if len(in.children) > 0 {
		res += fmt.Sprintf(" (dir, size=%d)\n", in.size)
	} else {
		res += fmt.Sprintf(" (file, size=%d)\n", in.size)
	}

	for _, c := range in.children {
		res += c.Stringd(depth + 1)
	}
	return res
}

func (in INode) String() string {
	return in.Stringd(0)
}

func handleCD(cmd string, cur *INode, root *INode) (*INode, error) {
	fields := strings.Fields(cmd)
	if len(fields) != 3 {
		return nil, fmt.Errorf("invalid cd command %s", cmd)
	}
	target := fields[2]

	if target == "/" {
		return root, nil
	}

	if target == ".." {
		if cur.parent == nil {
			return nil, fmt.Errorf("cannot go to .. without parent")
		}
		return cur.parent, nil
	}

	for _, child := range cur.children {
		if child.name == target {
			return child, nil
		}
	}

	return nil, fmt.Errorf("dir not found: %s", target)
}

func handleLS(lines *[]string, start int, cur *INode) (int, error) {
	var curL int
	for curL = start; curL < len(*lines); curL++ {
		line := (*lines)[curL]
		if line[0] == '$' {
			break
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return 0, fmt.Errorf("invalid file line: %s", line)
		}

		if fields[0] == "dir" {
			cur.AddChild(fields[1], 0)
		} else {
			size, err := strconv.ParseInt(fields[0], 10, 64)
			if err != nil {
				return 0, fmt.Errorf("could not parse size: %v", err)
			}
			cur.AddChild(fields[1], int(size))
		}
	}

	return curL - 1, nil
}

func genFS(lines *[]string) (*INode, error) {
	root := NewINode("/", 0)
	cur := &root
	for i := 0; i < len(*lines); i++ {
		line := (*lines)[i]
		if strings.Contains(line, "cd") {
			next, err := handleCD(line, cur, &root)
			if err != nil {
				return nil, fmt.Errorf("failed to process command '%s', got error: %v", line, err)
			}
			cur = next
			continue
		} else if strings.Contains(line, "ls") {
			last, err := handleLS(lines, i+1, cur)
			if err != nil {
				return nil, fmt.Errorf("failed to process command '%s', got error: %v", line, err)
			}
			i = last
			continue
		}
		return nil, fmt.Errorf("landed on a non-command line: %s\n", line)
	}
	return &root, nil
}

func (in INode) Walk(work func(INode) error) error {
	queue := in.children

	for len(queue) > 0 {
		l := len(queue)
		cin := queue[l-1]
		queue = queue[:l-1]
		err := work(*cin)
		if err != nil {
			return err
		}
		queue = append(queue, cin.children...)
	}

	return nil
}

func (in *INode) ComputeSizes() {
	if in.size != 0 {
		return
	}

	for _, c := range in.children {
		c.ComputeSizes()
		in.size += c.size
	}
}

func solvePart(part string) int {
	parsed, err := parse.GetLines("day7/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	root, err := genFS(&parsed)
	if err != nil {
		fmt.Printf("Generating FS failed: %v\n", err)
		return 1
	}
	root.ComputeSizes()
	res := 0
	if part == "1" {
		root.Walk(func(in INode) error {
			if in.IsDir() && in.size <= 100000 {
				res += in.size
			}
			return nil
		})
	} else {
		target := 30000000
		disk := 70000000
		tot := root.size
		res = root.size
		root.Walk(func(in INode) error {
			if in.IsDir() && in.size < res {
				if (disk - tot + in.size) >= target {
					res = in.size
				}
			}
			return nil
		})
	}
	fmt.Printf("Part %s: %v\n", part, res)
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
