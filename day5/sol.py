#!/usr/bin/python3.7

lines = [ l.strip() for l in open('input.txt').readlines() ]

bin_conv = {
    'F': '0',
    'B': '1',
    'L': '0',
    'R': '1'
}

seats = { int(''.join(bin_conv[c] for c in line), 2) for line in lines }

for seat in seats:
    if (seat + 2) in seats and (seat + 1) not in seats:
        print(seat + 1)
    if (seat - 2) in seats and (seat - 1) not in seats:
        print(seat - 1)
