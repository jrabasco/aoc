from typing import Iterable, TypeVar, Callable, Optional, Tuple

T = TypeVar('T')
S = TypeVar('S')
class Grid:
    def __init__(self,
                 lines: Iterable[Iterable[S]],
                 conv: Optional[Callable[[S], T]] = None):
        # that is if T and S are the same
        if conv is None:
            conv = lambda x: x
        self._grid = [ [ conv(item) for item in line ] for line in lines]
        self._h = len(self._grid)
        self._w = 0
        self._max_x = None
        self._max_y = None
        if self._h > 0:
            self._w = len(self._grid[0])
            self._max_x = self._h - 1
        if self._w > 0:
            self._max_y = self._w - 1


    @property
    def h(self) -> int:
        return self._h

    @property
    def w(self) -> int:
        return self._w

    @property
    def max_x(self) -> Optional[int]:
        return self._max_x

    @property
    def max_y(self) -> Optional[int]:
        return self._max_y

    def get(self, x: int, y: int) -> T:
        return self._grid[x][y]

    def neighbours(self, x: int, y: int) -> Iterable[Tuple[int, int]]:
        if self._h == 0 or self._w == 0:
            return []
        if x < self._max_x:
            yield (x+1, y)
        if x > 0:
            yield (x-1, y)
        if y < self._max_y:
            yield (x, y+1)
        if y > 0:
            yield (x, y-1)

    def rows(self):
        for row in self._grid:
            yield row

    def row(self, i):
        return self._grid[i]

    def columns(self):
        for j in range(self._max_y):
            yield [ row[j] for row in self._grid ]

    def __iter__(self):
        return self.rows()

    def __repr__(self):
        return '\n'.join(
                    ''.join(repr(item) for item in row)
                    for row in self.rows()
                )


