package day15

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"strconv"
	"strings"
)

func HASH(label string) int {
	cur := 0
	for _, r := range label {
		cur = ((cur + int(r)) * 17) % 256
	}
	return cur
}

type Lense struct {
	label string
	focal int
}

type HASHMAP struct {
	boxes [][]Lense
}

func NewHASHMAP() HASHMAP {
	boxes := [][]Lense{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, []Lense{})
	}
	return HASHMAP{boxes}
}

func (h *HASHMAP) Add(label string, focal int) {
	hash := HASH(label)
	found := false
	for i, l := range h.boxes[hash] {
		if l.label == label {
			found = true
			h.boxes[hash][i].focal = focal
			break
		}
	}

	if !found {
		h.boxes[hash] = append(h.boxes[hash], Lense{label, focal})
	}
}

func (h *HASHMAP) Remove(label string) {
	hash := HASH(label)
	found := -1
	for i, l := range h.boxes[hash] {
		if l.label == label {
			found = i
			break
		}
	}

	if found == -1 {
		return
	}

	lbox := len(h.boxes[hash])
	for i := found + 1; i < lbox; i++ {
		h.boxes[hash][i-1] = h.boxes[hash][i]
	}

	h.boxes[hash] = h.boxes[hash][:lbox-1]
}

func parseEqual(instruction string) (string, int, error) {
	parts := strings.Split(instruction, "=")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid = instruction: %s", instruction)
	}

	focal, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}
	return parts[0], focal, nil
}

func parseRemove(instruction string) (string, error) {
	linst := len(instruction)
	if instruction[linst-1] != '-' {
		return "", fmt.Errorf("invalid - instruction: %s", instruction)
	}
	return instruction[:linst-1], nil
}

func (h *HASHMAP) Apply(instruction string) error {
	if strings.ContainsRune(instruction, '=') {
		label, focal, err := parseEqual(instruction)
		if err != nil {
			return err
		}
		h.Add(label, focal)
		return nil
	}
	if strings.ContainsRune(instruction, '-') {
		label, err := parseRemove(instruction)
		if err != nil {
			return err
		}
		h.Remove(label)
		return nil
	}
	return fmt.Errorf("unknown instruction type: %s", instruction)
}

func (h HASHMAP) FocusPower() int {
	res := 0
	for i, b := range h.boxes {
		for j, l := range b {
			res += (i + 1) * (j + 1) * l.focal
		}
	}
	return res
}

func Solution() int {
	parsed, err := parse.GetLines("day15/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	code := strings.Split(parsed[0], ",")
	res := 0
	h := NewHASHMAP()
	for _, inst := range code {
		res += HASH(inst)
		err := h.Apply(inst)
		if err != nil {
			fmt.Printf("Failed to apply '%s': %v\n", inst, err)
			return 1
		}
	}
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", h.FocusPower())
	return 0
}
