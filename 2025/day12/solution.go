package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jrabasco/aoc/2025/framework/parse"
)

type Shape struct {
	id   int
	grid [][]bool // true = occupied (#), false = empty (.)
	area int      // number of occupied cells
}

type Area struct {
	width  int
	height int
	shapes []int // indices of shapes and their counts (shape_id, count pairs)
}

type Input struct {
	shapes []Shape
	areas  []Area
}

func parseInput(lines []string) (*Input, error) {
	result := &Input{
		shapes: []Shape{},
		areas:  []Area{},
	}

	i := 0
	// Parse shapes
	for i < len(lines) {
		line := lines[i]

		// Check if we've reached the areas section
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			break
		}

		// Skip empty lines
		if line == "" {
			i++
			continue
		}

		// Parse shape header (e.g., "0:")
		if strings.HasSuffix(line, ":") {
			shapeID, err := strconv.Atoi(strings.TrimSuffix(line, ":"))
			if err != nil {
				return nil, fmt.Errorf("invalid shape ID: %s", line)
			}

			// Parse the grid
			grid := [][]bool{}
			i++
			area := 0
			for i < len(lines) && lines[i] != "" {
				row := []bool{}
				for _, ch := range lines[i] {
					row = append(row, ch == '#')
					if ch == '#' {
						area++
					}
				}
				grid = append(grid, row)
				i++
			}

			result.shapes = append(result.shapes, Shape{
				id:   shapeID,
				grid: grid,
				area: area,
			})
		} else {
			i++
		}
	}

	// Parse areas
	for i < len(lines) {
		line := lines[i]
		if line == "" {
			i++
			continue
		}

		// Parse area (e.g., "4x4: 0 0 0 0 2 0")
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			i++
			continue
		}

		// Parse dimensions
		dims := strings.Split(parts[0], "x")
		if len(dims) != 2 {
			return nil, fmt.Errorf("invalid area dimensions: %s", parts[0])
		}

		width, err := strconv.Atoi(dims[0])
		if err != nil {
			return nil, fmt.Errorf("invalid width: %s", dims[0])
		}

		height, err := strconv.Atoi(dims[1])
		if err != nil {
			return nil, fmt.Errorf("invalid height: %s", dims[1])
		}

		// Parse shape counts
		shapeCounts := []int{}
		countStrs := strings.Fields(parts[1])
		for _, countStr := range countStrs {
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return nil, fmt.Errorf("invalid shape count: %s", countStr)
			}
			shapeCounts = append(shapeCounts, count)
		}

		result.areas = append(result.areas, Area{
			width:  width,
			height: height,
			shapes: shapeCounts,
		})

		i++
	}

	return result, nil
}

func p1(input *Input) int {
	count := 0

	for _, area := range input.areas {
		// Calculate total area needed for all shapes
		totalShapeArea := 0
		for shapeIdx, shapeCount := range area.shapes {
			if shapeIdx < len(input.shapes) {
				totalShapeArea += input.shapes[shapeIdx].area * shapeCount
			}
		}

		// Calculate available area
		availableArea := area.width * area.height

		// Check if all shapes can fit
		if totalShapeArea <= availableArea {
			count++
		}
	}

	return count
}

func Solution() int {
	parsed, err := parse.GetLinesAsOne[*Input]("day12/input.txt", parseInput)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	fmt.Printf("Solution: %d\n", p1(parsed))
	return 0
}
