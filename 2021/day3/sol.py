#!/usr/bin/python3.8

lines = [line.strip() for line in open('input.txt').readlines()]

total = len(lines)
counts = len(lines[0])*[0]

for line in lines:
    for i, digit in enumerate(line):
        counts[i] += int(digit)

gamma = 0
epsilon = 0
mpow = len(counts) - 1
for i, count in enumerate(counts):
    if count > total/2:
        gamma += 2**(mpow - i)
    else:
        epsilon += 2**(mpow - i)

print(f'Part 1: {gamma * epsilon}')


oxy_lines = lines[:]

for i in range(len(oxy_lines[0])):
    p0 = []
    p1 = []
    for line in oxy_lines:
        if line[i] == "0":
            p0.append(line)
        else:
            p1.append(line)

    if len(p1) >= len(p0):
        oxy_lines = p1
    else:
        oxy_lines = p0

    if len(oxy_lines) == 1:
        break
oxy_num = int(oxy_lines[0], 2)

co2_lines = lines[:]

for i in range(len(co2_lines[0])):
    p0 = []
    p1 = []
    for line in co2_lines:
        if line[i] == "0":
            p0.append(line)
        else:
            p1.append(line)

    if len(p1) < len(p0):
        co2_lines = p1
    else:
        co2_lines = p0

    if len(co2_lines) == 1:
        break

co2_num = int(co2_lines[0], 2)

print(f'Part 2: {oxy_num * co2_num}')
