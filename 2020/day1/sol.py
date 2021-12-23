#!/usr/bin/python3.7

lines = [ l.strip() for l in open('input.txt').readlines() ]

numbers = { int(l) for l in lines }

for n in numbers:
    if (2020 - n) in numbers:
        print(n * (2020-n))

for n in numbers:
    for m in numbers:
        if n != m:
            if (2020-n-m) in numbers:
                print(n *  m * (2020-n-m))

