#!/usr/bin/python3.8
import heapq
from dataclasses import dataclass

from grid import Grid

lines = [line.strip() for line in open('input.txt').readlines()]

GRID = Grid(lines, conv=int)

def min_total_risk(grid, x, y):
    # Dijkstra edition
    visited = set()
    costs = dict()
    costs[(0,0)] = 0
    q = [(0, 0, 0)]
    while q:
        cost, cx, cy = heapq.heappop(q)
        for nb in grid.neighbours(cx, cy):
            if nb in visited:
                continue
            nbx, nby = nb
            nbcost = grid.get(nbx, nby)
            if (nbx, nby) not in costs or (nbcost + costs[(cx, cy)]) < costs[nb]:
                costs[nb] = (nbcost + costs[(cx, cy)])
                heapq.heappush(q, (costs[nb], nbx, nby))
        visited.add((cx, cy))
        if (x,y) in visited:
            return costs[(x,y)]
        current = None
    return visited

def make_large_grid(grid):
    res = [[0 for _ in range(grid.h*5)] for __ in range(grid.w*5)]
    for offx in range(5):
        for offy in range(5):
            for i in range(grid.h):
                for j in range(grid.w):
                    res[i + grid.h*offx][j + grid.w*offy] = (grid.get(i,j) + offx + offy - 1)%9 + 1
    return Grid(res)


print(f'Part 1: {min_total_risk(GRID, GRID.max_x, GRID.max_y)}')

LARGE_GRID = make_large_grid(GRID)

print(f'Part 2: {min_total_risk(LARGE_GRID, LARGE_GRID.max_x, LARGE_GRID.max_y)}')
