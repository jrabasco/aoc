#!/usr/bin/python3.7

lines = [ l.strip() for l in open('input.txt').readlines() ]

area = []
empty = '.'
tree = '#'

for line in lines:
    area.append(list(line))

width = len(area[0])
last = len(area) - 1

def count_trees(right, down):
    posx = 0
    posy = 0

    trees = 0
    while posy < last:
        posy += down
        posx = (posx + right) % width
        if area[posy][posx] == tree:
            trees += 1
    return trees

slopes = [(1,1), (3,1), (5,1), (7,1), (1, 2)]
results = [count_trees(right, down) for right, down in slopes]

answer = 1
for r in results:
    answer *= r

print(answer)
