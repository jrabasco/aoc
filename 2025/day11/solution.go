package day11

import (
	"fmt"
	"strings"

	"github.com/jrabasco/aoc/2025/framework/parse"
)

type Graph struct {
	deviceToIdx map[string]int
	idxToDevice []string
	adjMatrix   [][]bool
	numDevices  int
}

func parseGraph(lines []string) (*Graph, error) {
	// First pass: collect all unique device names
	deviceSet := make(map[string]bool)
	connections := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		source := parts[0]
		deviceSet[source] = true

		if parts[1] != "" {
			targets := strings.Fields(parts[1])
			for _, target := range targets {
				deviceSet[target] = true
			}
			connections[source] = targets
		}
	}

	// Create device index mappings
	idxToDevice := make([]string, 0, len(deviceSet))
	for device := range deviceSet {
		idxToDevice = append(idxToDevice, device)
	}

	deviceToIdx := make(map[string]int)
	for i, device := range idxToDevice {
		deviceToIdx[device] = i
	}

	// Build adjacency matrix
	n := len(idxToDevice)
	adjMatrix := make([][]bool, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]bool, n)
	}

	for source, targets := range connections {
		sourceIdx := deviceToIdx[source]
		for _, target := range targets {
			targetIdx := deviceToIdx[target]
			adjMatrix[sourceIdx][targetIdx] = true
		}
	}

	return &Graph{
		deviceToIdx: deviceToIdx,
		idxToDevice: idxToDevice,
		adjMatrix:   adjMatrix,
		numDevices:  n,
	}, nil
}

type PathState struct {
	node         int
	visitedNode1 bool
	visitedNode2 bool
}

func countPathsWithNodes(graph *Graph, current, target, node1, node2 int, visitedNode1, visitedNode2 bool, memo map[PathState]int) int {
	// Update visited flags if we're at node1 or node2 (if they are specified, i.e., >= 0)
	if node1 >= 0 && current == node1 {
		visitedNode1 = true
	}
	if node2 >= 0 && current == node2 {
		visitedNode2 = true
	}

	// If we reached the target
	if current == target {
		// If both nodes are specified (>= 0), only count if we visited both
		// If either is < 0 (unspecified), don't check that condition
		requireNode1 := node1 >= 0
		requireNode2 := node2 >= 0

		if (!requireNode1 || visitedNode1) && (!requireNode2 || visitedNode2) {
			return 1
		}
		return 0
	}

	// Check memo
	state := PathState{current, visitedNode1, visitedNode2}
	if count, exists := memo[state]; exists {
		return count
	}

	totalPaths := 0
	// Explore all neighbors
	for neighbor := 0; neighbor < graph.numDevices; neighbor++ {
		if graph.adjMatrix[current][neighbor] {
			totalPaths += countPathsWithNodes(graph, neighbor, target, node1, node2, visitedNode1, visitedNode2, memo)
		}
	}

	memo[state] = totalPaths
	return totalPaths
}

func p1(graph *Graph) int {
	// Find "you" and "out" device indices
	youIdx, youExists := graph.deviceToIdx["you"]
	outIdx, outExists := graph.deviceToIdx["out"]

	if !youExists || !outExists {
		return 0
	}

	// Count all paths from "you" to "out" with no required intermediate nodes
	memo := make(map[PathState]int)
	return countPathsWithNodes(graph, youIdx, outIdx, -1, -1, false, false, memo)
}

func p2(graph *Graph) int {
	// Find "svr", "out", "dac", and "fft" device indices
	svrIdx, svrExists := graph.deviceToIdx["svr"]
	outIdx, outExists := graph.deviceToIdx["out"]
	dacIdx, dacExists := graph.deviceToIdx["dac"]
	fftIdx, fftExists := graph.deviceToIdx["fft"]

	if !svrExists || !outExists || !dacExists || !fftExists {
		return 0
	}

	// Count paths that go through both "dac" and "fft" in any order
	memo := make(map[PathState]int)
	return countPathsWithNodes(graph, svrIdx, outIdx, dacIdx, fftIdx, false, false, memo)
}

func Solution() int {
	graph, err := parse.GetLinesAsOne[*Graph]("day11/input.txt", parseGraph)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	fmt.Printf("Part 1: %d\n", p1(graph))
	fmt.Printf("Part 2: %d\n", p2(graph))
	return 0
}
