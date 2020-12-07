#!/usr/bin/python3.7

lines = [ l.strip() for l in open('input.txt').readlines() ]

groups = [[]]
for line in lines:
    if line == '':
        groups.append([])
    else:
        groups[-1].append(line)


print('Part 1:')
count = 0
for group in groups:
    yes = set()
    for answer in group:
        yes |= set(answer)
    count += len(yes)

print(count)
print('Part 2:')
count = 0
for group in groups:
    yes = set('abcdefghijklmnopqrstuvwxyz')
    for answer in group:
        yes &= set(answer)
    count += len(yes)
print(count)
