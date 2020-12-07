#!/usr/bin/python3.7

lines = [ l.strip() for l in open('input.txt').readlines() ]
valid_counts = 0

for line in lines:
    rule, password = line.split(':')
    password = password.strip()
    positions, letter = rule.split(' ')
    positions = [int(i) for i in positions.split('-')]
    matches = [password[position-1] == letter for position in positions]
    if sum(matches) == 1:
        print(positions)
        print(letter)
        print(password)
        print(matches)
        valid_counts+=1

print(valid_counts)
