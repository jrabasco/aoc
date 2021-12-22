from __future__ import annotations
from dataclasses import dataclass
from typing import Optional, List

@dataclass
class Point:
    x: int
    y: int
    z: int

class Cuboid:
    def __init__(self, x0: int, x1: int, y0: int, y1: int, z0: int, z1: int):
        if x1 < x0 or y1 < y0 or z1 < z0:
            raise ValueError('coordinates need to be given in ascending order')
        self._from = Point(x0, y0, z0)
        self._to = Point(x1, y1, z1)

    def has_intersect(self, other: Cuboid) -> bool:
        return not (self._from.x > other._to.x
           or self._from.y > other._to.y
           or self._from.z > other._to.z
           or self._to.x < other._from.x
           or self._to.y < other._from.y
           or self._to.z < other._from.z)

    def intersect(self, other: Cuboid) -> Optional[Cuboid]:
        if not self.has_intersect(other):
            return None

        fromx = max(self._from.x, other._from.x)
        fromy = max(self._from.y, other._from.y)
        fromz = max(self._from.z, other._from.z)

        tox = min(self._to.x, other._to.x)
        toy = min(self._to.y, other._to.y)
        toz = min(self._to.z, other._to.z)

        return Cuboid(fromx, tox, fromy, toy, fromz, toz)

    def minus(self, other: Cuboid) -> List[Cuboid]:
        """
        Returns a list of cuboids that are left after
        intersecting
        """
        if not self.has_intersect(other):
            return [self]
        res = []
        if self._from.x < other._from.x:
            res.append(Cuboid(self._from.x, other._from.x-1,
                              self._from.y, self._to.y,
                              self._from.z, self._to.z))
        if self._to.x > other._to.x:
            res.append(Cuboid(other._to.x+1, self._to.x,
                              self._from.y, self._to.y,
                              self._from.z, self._to.z))
        fromx = max(self._from.x, other._from.x)
        tox = min(self._to.x, other._to.x)

        if self._from.y < other._from.y:
            res.append(Cuboid(fromx, tox,
                              self._from.y, other._from.y-1,
                              self._from.z, self._to.z))
        if self._to.y > other._to.y:
            res.append(Cuboid(fromx, tox,
                              other._to.y+1, self._to.y,
                              self._from.z, self._to.z))
        fromy = max(self._from.y, other._from.y)
        toy = min(self._to.y, other._to.y)

        if self._from.z < other._from.z:
            res.append(Cuboid(fromx, tox,
                              fromy, toy,
                              self._from.z, other._from.z-1))
        if self._to.z > other._to.z:
            res.append(Cuboid(fromx, tox,
                              fromy, toy,
                              other._to.z+1, self._to.z))
        return res

    def volume(self) -> int:
        xlength = self._to.x - self._from.x + 1
        ylength = self._to.y - self._from.y + 1
        zlength = self._to.z - self._from.z + 1

        return xlength*ylength*zlength

    def __eq__(self, other) -> bool:
        if not isinstance(other, Cuboid):
            return False
        return self._from == other._from and self._to == other._to

    def __repr__(self) -> str:
        return (f'Cuboid(x0={self._from.x}, x1={self._to.x},'
                f'y0={self._from.y}, y1={self._to.y},'
                f'z0={self._from.z}, z1={self._to.z})')

