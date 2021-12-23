#!/usr/bin/python3.8
from framework import Parser, Cuboid
from typing import List, Dict

p = Parser('input.txt')

words_and_ints = zip(p.words_by_line(), p.ints_by_line())


def volume(space: Dict[Cuboid, int]) -> int:
    """
    The value in the dict shows the "intensity" which takes into account
    the adjustments made necessary by intersections.
    """
    return sum(val*c.volume() for c, val in space.items())

part1 = True
space = {}
first = True
for words, ints in words_and_ints:
    if first:
        first = False
        space[Cuboid(*ints)] = 1
        continue
    max_abs = max(abs(x) for x in ints)
    if max_abs > 50 and part1:
        print(f'Part 1: {volume(space)}')
        part1 = False
    onoff, _, __, ___ = words
    piece = Cuboid(*ints)
    for cuboid, val in space.copy().items():
        inter = cuboid.intersect(piece)
        # remove the intersection:
        # - val positive, turning on:
        #   we need to reduce the light level at the intersection to compensate
        #   for the volume being added twice -> -= val
        # - val negative, turning on:
        #   this area was being compensated for being counted twice, we need to
        #   negate this as we will compensate once for each times the current
        #   piece will intersect with the pieces that originally produced this
        #   intersection -> -= val
        # - val positive, turning off:
        #   this is simply adding the opposite of val to turn off this
        #   intersection with the cuboid -> -= val
        # - val negative, turning off:
        #   this area was being compensated for being counted twice, we need
        #   to negate this as we will turn off the lights for each times the
        #   current piece will intersect with one of the pieces that created
        #   the intersection -> -= val

        if inter is not None:
            if inter not in space:
                space[inter] = 0
            space[inter] -= val
    # After all the adjustments for the intersection, simply add some light
    #  value to the piece
    if onoff == "on":
        if piece not in space:
            space[piece] = 0
        space[piece] += 1

print(f'Part 2: {volume(space)}')
