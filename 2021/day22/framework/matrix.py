from typing import List
from numbers import Number
import math

class Matrix:
    # construct from values given as a list of columns
    def __init__(self, values: List[List[Number]]):
        self._M = len(values)
        self._data = []
        if self._M == 0:
            self._N = 0
            return

        self._N = len(values[0])
        self._data = [[] for _ in range(self._N)]
        # store in row order
        for col in values:
            if len(col) != self._N:
                raise ValueError(f'Inconsistent dimensions, '
                                 f'first column has {self._N} elements, but found a '
                                 f'column with {len(col)} elements.')
            for i, elm in enumerate(col):
                self._data[i].append(elm)

    @property
    def N(self):
        return self._N

    @property
    def M(self):
        return self._M

    def get(self, i: int, j: int) -> Number:
        if i < 0 or j < 0 or i >= self._N or j >= self._M:
            raise IndexError(f'i must be in [0, {self._N}[ and j must be in '
                             f'[0, {self._M}[ but got ({i=}, {j=})')
        return self._data[i][j]

    @property
    def transposed(self):
        return Matrix(self._data)

    def __add__(self, other):
        if not isinstance(other, Matrix):
            raise TypeError(f"unsupported operand type(s) for +: "
                            f"'{self.__class__.__name__}' "
                            f"and '{type(other).__name__}'")
        if self._M != other._M or self._N != other._N:
            raise ValueError(f'Incompatible sizes for addition: '
                             f'{self._N}x{self._M} and {other._N}x{other._M}')

        cols = [[] for _ in range(self._M)]
        for i,row in enumerate(self._data):
            for j, elm in enumerate(row):
                cols[j].append(elm + other._data[i][j])
        if isinstance(self, Vector):
            return Vector(cols[0])
        return Matrix(cols)

    def __sub__(self, other):
        if not isinstance(other, Matrix):
            raise TypeError(f"unsupported operand type(s) for -: "
                            f"'{self.__class__.__name__}' "
                            f"and '{type(other).__name__}'")
        if self._M != other._M or self._N != other._N:
            raise ValueError(f'Incompatible sizes for subtraction: '
                             f'{self._N}x{self._M} and {other._N}x{other._M}')

        cols = [[] for _ in range(self._M)]
        for i,row in enumerate(self._data):
            for j, elm in enumerate(row):
                cols[j].append(elm - other._data[i][j])
        if isinstance(self, Vector):
            return Vector(cols[0])
        return Matrix(cols)

    def __mul__(self, other):
        if isinstance(other, Number):
            cols = [[] for _ in range(self._M)]
            for row in self._data:
                for i, elm in enumerate(row):
                    cols[i].append(elm * other)
            if isinstance(self, Vector):
                return Vector(cols[0])
            return Matrix(cols)

        if isinstance(other, Matrix):
            if self._M != other._N:
                raise ValueError(f'Incompatible sizes for multiplication: '
                                 f'{self._N}x{self._M} and {other._N}x{other._M}')

            cols = [[] for _ in range(other._M)]
            for i in range(self._N):
                for j in range(other._M):
                    s = sum(self._data[i][k] * other._data[k][j] for k in range(self._M))
                    cols[j].append(s)
            if isinstance(other, Vector):
                return Vector(cols[0])
            return Matrix(cols)


        raise TypeError(f"unsupported operand type(s) for *: "
                            f"'{self.__class__.__name__}' "
                        f"and '{type(other).__name__}'")

    def __eq__(self, other) -> bool:
        if not isinstance(other, Matrix):
            return False

        if self._N != other._N or self._M != other._M:
            return False

        for i in range(self._N):
            for j in range(self._M):
                if self._data[i][j] != other._data[i][j]:
                    return False
        return True

    def __hash__(self) -> int:
        return hash(str(self))

    def __repr__(self):
        cols = [[] for _ in range(self._M)]
        for row in self._data:
            for i, elm in enumerate(row):
                cols[i].append(elm)
        cols_str = [ f'[{",".join(str(elm) for elm in col)}]' for col in cols]
        return f'{self.__class__.__name__}({",".join(cols_str)})'


class Vector(Matrix):
    # a vector is a matrix with only one column
    def __init__(self, values: List[Number]):
        super().__init__([values])

    def get(self, i: int) -> Number:
        return super().get(i, 0)

    def cartesian_distance(self, other):
        if not isinstance(other, Vector):
            raise TypeError(f"Cartesian distance is undefined for types 'Vector' and "
                            f"{type(other).__name__}")
        if self._N != other._N:
            raise ValueError(f"Incompatible dimensions : {self._N} vs {other._N}")
        s = sum((self._data[i][0]-other._data[i][0])**2 for i in range(self._N))
        return math.sqrt(s)


    def manhattan_distance(self, other):
        if not isinstance(other, Vector):
            raise TypeError(f"Manhattan distance is undefined for types 'Vector' and "
                            f"{type(other).__name__}")
        if self._N != other._N:
            raise ValueError(f"Incompatible dimensions : {self._N} vs {other._N}")
        return sum(abs(self._data[i][0]-other._data[i][0]) for i in range(self._N))
