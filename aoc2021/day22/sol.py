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
    to_remove = [piece]
    n_space = []
    while space:
        c = space.pop()
        frags = [c]
        found = True
        while found:
            found = False
            i = 0
            l_remove = len(to_remove)
            l_frags = len(frags)
            while i < l_remove and not found:
                rm = to_remove[i]
                j = 0
                while j < l_frags and not found:
                    f = frags[j]
                    if not f.has_intersect(rm):
                        j+=1
                        continue

                    found = True
                    n_remove = rm.minus(f)
                    to_remove = to_remove[:i] + to_remove[i+1:] + n_remove
                    n_frags = f.minus(rm)
                    frags = frags[:j] + frags[j+1:] + n_frags
                    j += 1
                i += 1
        n_space += frags
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
