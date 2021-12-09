#!/usr/bin/python3.8

lines = [line.strip() for line in open('input.txt').readlines()]

grid = []

for line in lines:
    grid.append([int(dig) for dig in line])

max_i = len(grid) - 1
max_j = len(grid[0]) - 1

res = 0
for i, row in enumerate(grid):
    for j, height in enumerate(row):
        if i < max_i and height >= grid[i+1][j]:
            continue

        if i > 0 and height >= grid[i-1][j]:
            continue

        if j < max_j and height >= grid[i][j+1]:
            continue

        if j > 0 and height >= grid[i][j-1]:
            continue

        res += height + 1

print(f'Part 1: {res}')

# BFS to find the size
def basin_size(x,y):
    visited = {(x,y)}
    queue = [(x,y)]

    while queue:
        i, j = queue.pop(0)
        height = grid[i][j]
        if i < max_i and grid[i+1][j] < 9 and height < grid[i+1][j]:
            queue.append((i+1, j))
            visited.add((i+1, j))

        if i > 0 and grid[i-1][j] < 9 and height < grid[i-1][j]:
            queue.append((i-1, j))
            visited.add((i-1, j))

        if j < max_j and grid[i][j+1] < 9 and height < grid[i][j+1]:
            queue.append((i, j+1))
            visited.add((i, j+1))

        if j > 0 and grid[i][j-1] < 9 and height < grid[i][j-1]:
            queue.append((i, j-1))
            visited.add((i, j-1))

    return len(visited)

sizes = []
for i, row in enumerate(grid):
    for j, height in enumerate(row):
        if i < max_i and height >= grid[i+1][j]:
            continue

        if i > 0 and height >= grid[i-1][j]:
            continue

        if j < max_j and height >= grid[i][j+1]:
            continue

        if j > 0 and height >= grid[i][j-1]:
            continue

        sizes.append(basin_size(i,j))

largest_3 = sorted(sizes, reverse=True)[:3]

print(f'Part 2: {largest_3[0] * largest_3[1] * largest_3[2]}')
