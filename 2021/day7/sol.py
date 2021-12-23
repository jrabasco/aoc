#!/usr/bin/python3.8
from statistics import median
positions = [int(nb) for nb in open('input.txt').read().split(',')]

# Part 1
pivot = median(positions)
res = 0
for pos in positions:
    res += abs(pos - pivot)
print(f'Median pivot is {pivot}. Answer is: {res}')

# Part 2

def costi(pivot, pos):
    return sum(range(1, abs(pos-pivot)+1))

def cost(pivot, positions):
    return sum(costi(pivot, pos) for pos in positions)

mean = sum(positions)/len(positions)
pivot = int(mean)
res = min(cost(pivot, positions), cost(pivot + 1, positions))
print(f'Mean is {mean}. Answer is: {res}')

# Test code
#print(sum(positions)/len(positions))
#for pivot in range(200):
#    res = 0
#    high = 0
#    low = 0
#    for pos in positions:
#        res += cost(pivot, pos)
#        if pos > pivot:
#            high += 1
#        if pos < pivot:
#            low += 1
#    print(f'{pivot}: {high=} {low=} {res=}')
