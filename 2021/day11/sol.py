#!/usr/bin/python3.8
import time

lines = [line.strip() for line in open('input.txt').readlines()]

grid = []
N = 10

for line in lines:
    grid.append([int(dig) for dig in line])

def step(g):
    flashes = 0
    # first increment
    for i in range(N):
        for j in range(N):
            g[i][j] += 1

    # then flash
    nflashes = True
    while nflashes:
        nflashes = False
        for i in range(N):
            for j in range(N):
                if g[i][j] > 9:
                    flash(g, i, j)
                    flashes += 1
                    nflashes = True

    return flashes


def flash(g, row, col):
    top = max(row-1, 0)
    bottom = min(row+1, N-1)
    left = max(col-1, 0)
    right = min(col+1, N-1)
    for i in range(top, bottom+1):
        for j in range(left, right+1):
            if g[i][j] != 0:
                g[i][j] += 1
    g[row][col] = 0


def print_grid(g):
    for row in g:
        print(''.join('\033[1;33;40m0\033[0m' if dig == 0 else str(dig) for dig in row))

res1 = 0
res2 = -1
i = 0
while res2 < 0:
    print(f'Step {i}:')
    print()
    print_grid(grid)
    print(13*'\033[F')
    cur_flashes = step(grid)
    i += 1
    if i <= 100:
        res1 += cur_flashes
    if cur_flashes == 100:
        res2 = i
    time.sleep(0.05)


print(f'Step {i}:')
print()
print_grid(grid)
print()
print(f'Part 1: {res1}')
print(f'Part 2: {res2}')
