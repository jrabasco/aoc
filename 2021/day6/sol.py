#!/usr/bin/python3.8

fishes_list = [ int(nb) for nb in open('input.txt').read().strip().split(',') ]
fishes = [0]*9
for f in fishes_list:
    fishes[f] += 1

# Simple straightforward too slow solution for part 1
for _ in range(80):
    for i in range(len(fishes_list)):
        if fishes_list[i] == 0:
            fishes_list[i] = 6
            fishes_list.append(8)
        else:
            fishes_list[i] -= 1

print(f'Part 1: {len(fishes_list)}')

# optimised solution for part 2
for _ in range(256):
    reset = fishes[0]
    for f in range(1, len(fishes)):
        fishes[f-1] = fishes[f]
    fishes[8] = reset
    fishes[6] += reset

print(f'Part 2: {sum(fishes)}')
