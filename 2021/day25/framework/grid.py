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

    def down(self, x: int, y: int, wrap: bool=False)-> Tuple[int,int]:
        if x >= self._max_x:
            if wrap:
                x = 0
            else:
                raise ValueError(f'cannot go down from {x}')
        else:
            x = x+1
        return (x, y)

    def up(self, x: int, y: int, wrap: bool=False)-> Tuple[int,int]:
        if x == 0:
            if wrap:
                x = self._max_x
            else:
                raise ValueError(f'cannot go up from {x}')
        else:
            x = x-1
        return (x, y)

    def right(self, x: int, y: int, wrap: bool=False)-> Tuple[int,int]:
        if y >= self._max_y:
            if wrap:
                y = 0
            else:
                raise ValueError(f'cannot go right from {y}')
        else:
            y = y + 1
        return (x, y)

    def left(self, x: int, y: int, wrap: bool=False)-> Tuple[int,int]:
        if y == 0:
            if wrap:
                y = self._max_y
            else:
                raise ValueError(f'cannot go left from {y}')
        else:
            y = y - 1
        return (x, y)

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

    def __str__(self):
        return '\n'.join(
                    ''.join(str(item) for item in row)
                    for row in self.rows()
                )

    def __repr__(self):
        return '\n'.join(
                    ''.join(repr(item) for item in row)
                    for row in self.rows()
                )

class InfiniteGrid:
    def __init__(self,
                 lines: Iterable[Iterable[S]],
                 default: T,
                 conv: Optional[Callable[[S], T]] = None):
        # that is if T and S are the same
        if conv is None:
            conv = lambda x: x
        self._grid = [
            [ conv(item) for item in line ]
            for line in lines
        ]
        self._h = len(self._grid)
        self._w = 0
        self._max_x = None
        self._max_y = None
        if self._h > 0:
            self._w = len(self._grid[0])
            self._max_x = self._h - 1
        if self._w > 0:
            self._max_y = self._w - 1
        self._min_x = 0
        self._min_y = 0
        self._default = default


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
        if x > self._max_x or x < 0 or y > self._max_y or y < 0:
            return self._default
        return self._grid[x][y]

    def square(self, x: int, y: int) -> Iterable[Tuple[int, int]]:
        for i in range(x-1, x+2):
            for j in range(y-1, y+2):
                yield self.get(i, j)

    def count(self, p: Callable[[T], bool]) -> int:
        return sum(
            1
            for row in self._grid
            for elm in row
            if p(elm)
        )


    def __str__(self):
        res = []
        for i in range(-1, self._h + 1):
            res.append(
                ''.join(str(self.get(i,j)) for j in range(-1, self._w+1))
            )
        return '\n'.join(res)

