#!/usr/bin/python3.8
from framework import Parser, Cuboid
from typing import List

p = Parser('input.txt')

words_and_ints = zip(p.words_by_line(), p.ints_by_line())
space = []


def volume(space: List[Cuboid]):
    return sum(c.volume() for c in space)

def add_piece(space: List[Cuboid], piece: Cuboid) -> List[Cuboid]:
    to_add = [piece]
    for c in space:
        to_add = [f for nc in to_add for f in nc.minus(c)]
    return space + to_add

def remove_piece(space: List[Cuboid], piece: Cuboid) -> List[Cuboid]:
    n_space = []
    for c in space:
        if piece.has_intersect(c):
            n_space += c.minus(piece)
            continue
        n_space.append(c)
    return n_space

part1 = True
for words, ints in words_and_ints:
    max_abs = max(abs(x) for x in ints)
    if max_abs > 50 and part1:
        print(f'Part 1: {volume(space)}')
        part1 = False
    onoff, _, __, ___ = words
    piece = Cuboid(*ints)
    if onoff == 'on':
        space = add_piece(space, piece)

    if onoff == 'off':
        space = remove_piece(space, piece)

print(f'Part 2: {volume(space)}')
