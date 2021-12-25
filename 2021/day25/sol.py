#!/usr/bin/python3.8
from framework import Parser, Grid
from typing import Tuple

p = Parser('input.txt')

G = p.grid()
print(G)

def neighb(g: Grid, direction: str, i:int, j: int) -> Tuple[int, int]:
    if direction == '>':
        return g.right(i, j, wrap=True)

    if direction == 'v':
        return g.down(i, j, wrap=True)

def move(g: Grid, direction: str) -> Tuple[Grid, bool]:
    moved = False
    n_g = []
    for row in g:
        n_g.append(row[:])
    for i, row in enumerate(g):
        for j, elm in enumerate(row):
            if elm == '.':
                continue

            if elm == direction:
                x,y = neighb(g, direction, i, j)
                if g.get(x,y) == '.':
                    moved = True
                    n_g[i][j] = '.'
                    n_g[x][y] = direction
    return Grid(n_g), moved
print()
moved = True
count = 0
while moved:
    count += 1
    G, mvr = move(G, '>')
    G, mvd = move(G, 'v')
    moved = mvr or mvd
print(G)
print(count)
