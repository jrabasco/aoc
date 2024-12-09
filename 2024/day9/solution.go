package day9

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"github.com/jrabasco/aoc/2024/framework/utils"
	"slices"
	"strconv"
)

type File struct {
	id   int
	size int
}

func printDisk(disk []File) {
	for _, f := range disk {
		for i := 0; i < f.size; i++ {
			if f.id >= 0 {
				fmt.Print(f.id)
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}

func checkSum(disk []File) int {
	res := 0
	n := 0
	for _, f := range disk {
		if f.id == -1 {
			n += f.size
			continue
		}
		for i := 0; i < f.size; i++ {
			res += n * f.id
			n++
		}
	}
	return res
}

func p2(diskMap []int) int {
	disk := []File{}

	isFile := true
	curFile := 0
	for _, elm := range diskMap {
		if isFile {
			file := File{curFile, elm}
			disk = append(disk, file)
			curFile++
		}
		if !isFile {
			file := File{-1, elm}
			disk = append(disk, file)
		}
		isFile = !isFile
	}

	seen := utils.NewSet[int]()
	// never want to move 0
	seen.Add(0)
	for i := len(disk) - 1; i >= 0; i-- {
		cFile := disk[i]
		if cFile.id == -1 || seen.Contains(cFile.id) {
			continue
		}
		seen.Add(cFile.id)
		for j := 0; j < len(disk) && j < i; j++ {
			gap := disk[j]
			if gap.id != -1 {
				continue
			}
			if gap.size < cFile.size {
				continue
			}

			disk[i].id = -1
			disk[j].id = cFile.id
			disk[j].size = cFile.size
			if gap.size > cFile.size {
				nGap := File{-1, gap.size - cFile.size}
				disk = slices.Insert(disk, j+1, nGap)
				i++ // this is needed because i needs to keep pointing at the same space
			}
			break
		}
	}
	// printDisk(disk)
	return checkSum(disk)
}

func p1(disk []int) int {
	diskA := []int{}
	isFile := true
	curFile := 0
	for _, elm := range disk {
		if isFile {
			for i := 0; i < elm; i++ {
				diskA = append(diskA, curFile)
			}
			curFile++
		}
		if !isFile {
			for i := 0; i < elm; i++ {
				diskA = append(diskA, -1)
			}
		}
		isFile = !isFile
	}

	for i := len(diskA) - 1; i >= 0; i-- {
		fid := diskA[i]
		diskA[i] = -1

		for j := 0; j < len(diskA); j++ {
			if diskA[j] == -1 {
				diskA[j] = fid
				break
			}
		}
	}

	chcksum := 0

	for i := 0; i < len(diskA) && diskA[i] >= 0; i++ {
		chcksum += i * diskA[i]
	}

	return chcksum
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[[]int]("day9/input.txt", func(lines []string) ([]int, error) {
		if len(lines) != 1 {
			return []int{}, fmt.Errorf("More than one line?")
		}
		res := []int{}
		for _, c := range lines[0] {
			nb, err := strconv.Atoi(string(c))
			if err != nil {
				return res, err
			}
			res = append(res, nb)
		}
		return res, nil
	})
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	fmt.Printf("Part 1: %d\n", p1(parsed))
	fmt.Printf("Part 2: %d\n", p2(parsed))
	return 0
}
