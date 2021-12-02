#!/usr/bin/python3.8

lines = [line.strip().split(' ') for line in open('input.txt').readlines()]

h = 0
d = 0

for line in lines:
    move = line[0]
    amount = int(line[1])

    if move == 'forward':
        h += amount

    if move == 'down':
        d += amount

    if move == 'up':
        d -= amount


print(f'Part 1: {d*h}')

h = 0
d = 0
a = 0
for line in lines:
    move = line[0]
    amount = int(line[1])

    if move == 'forward':
        h += amount
        d += a*amount

    if move == 'down':
        a += amount

    if move == 'up':
        a -= amount

print(f'Part 2: {d*h}')
