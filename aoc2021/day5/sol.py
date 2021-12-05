#!/usr/bin/python3.8
from dataclasses import dataclass

lines = [line.strip() for line in open('input.txt').readlines()]

@dataclass
class Point:
    x: int
    y: int

@dataclass
class Segment:
    start: Point
    end: Point
    def is_diagonal(self):
        return self.start.x != self.end.x and self.start.y != self.end.y

class SparseGrid:
    def __init__(self):
        self._data = {}

    def mark_point(self, p: Point):
        if (p.x, p.y) not in self._data:
            self._data[(p.x, p.y)] = 0
        self._data[(p.x, p.y)] += 1

    def mark_segment(self, s: Segment):
        xstep = int((s.end.x - s.start.x)/abs(s.end.x - s.start.x)) if s.start.x != s.end.x else 0
        ystep = int((s.end.y - s.start.y)/abs(s.end.y - s.start.y)) if s.start.y != s.end.y else 0

        it = Point(s.start.x, s.start.y)
        while it != s.end:
            self.mark_point(it)
            it.x += xstep
            it.y += ystep
        self.mark_point(it)

    def result(self):
        return sum(1 for val in self._data.values() if val > 1)

segments = []
for line in lines:
    parts = line.split(" -> ")
    p1 = Point(*[ int(c) for c in parts[0].split(',') ])
    p2 = Point(*[ int(c) for c in parts[1].split(',') ])
    segments.append(Segment(p1, p2))

g1 = SparseGrid()
for s in segments:
    if not s.is_diagonal():
        g1.mark_segment(s)

print(f'Part 1: {g1.result()}')

g2 = SparseGrid()
for s in segments:
    g2.mark_segment(s)

print(f'Part 2: {g2.result()}')
