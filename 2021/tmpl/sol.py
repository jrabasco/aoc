#!/usr/bin/python3.8
from framework import Parser

p = Parser('input.txt')

for line in p.lines():
    print(line)
