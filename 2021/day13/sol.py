#!/usr/bin/python3.8

lines = [line.strip() for line in open('input.txt').readlines()]

coords = []
folds = []
for line in lines:
    if len(line) == 0:
        continue
    if 'fold' in line:
        # remove 'fold along '
        line = line[11:]
        parts = line.split('=')
        folds.append((parts[0], int(parts[1])))
        continue

    parts = line.split(',')
    coords.append((int(parts[0]), int(parts[1])))

max_x = max(elm[0] for elm in coords) + 1
max_y = max(elm[1] for elm in coords) + 1

grid = [[0 for _ in range(max_x)] for _ in range(max_y) ]

for x,y in coords:
    grid[y][x] = 1

first = True
for fold in folds:
    axis, fld = fold
    if axis == 'x':
        # fold is on row fld, fld+1 is then the first column to move
        for i in range(max_y):
            target = fld - (i - fld)
            for j in range(fld+1, max_x):
                target = fld - (j - fld)
                old_hole = grid[i][target]
                new_hole = grid[i][j]
                grid[i][target] = max(old_hole, new_hole)
            grid[i] = grid[i][:fld]
        max_x = fld

    if axis == 'y':
        # fold is on row fld, fld+1 is then the first row to move
        for i in range(fld+1, max_y):
            target = fld - (i - fld)
            for j in range(max_x):
                old_hole = grid[target][j]
                new_hole = grid[i][j]
                grid[target][j] = max(old_hole, new_hole)
        grid = grid[:fld]
        max_y = fld

    if first:
        first = False
        res1 = sum(sum(row) for row in grid)
        print(f'Part 1: {res1}')

print('Part 2:')
conv = ['.', '#']
for row in grid:
    print(''.join(conv[elm] for elm in row))
