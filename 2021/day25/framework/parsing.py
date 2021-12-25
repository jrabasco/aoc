import re
from typing import Callable, Optional, TypeVar, Iterable, Union
from .grid import Grid

SEPARATORS = [',', ' ']

class Parser:
    def __init__(self, path: str):
        self._lines = [line.strip() for line in open(path).readlines()]

    T = TypeVar('T')
    def lines(self,
              t: Optional[Callable[[str], T]]=None,
              f: Optional[Callable[[T], bool]]=None) -> Iterable[Union[str, T]]:
        if t is None:
            t = lambda x : x
        if f is None:
            f = lambda x : True
        for line in self._lines:
            v = t(line[:])
            if f(v):
                yield v
    def grid(self, conv: Optional[Callable[[str], T]]=None) -> Grid:
        return Grid(self.lines(), conv)

    def words_by_line(self):
        """
        Splits each line into sections (words) using some common separators
        """
        for line in self._lines:
            words = [line[:]]
            seps = SEPARATORS[:]
            while seps:
                s = seps.pop()
                words = tuple(
                    w
                    for word in words
                    for w in word.split(s)
                )
            yield words

    def ints_by_line(self):
        """
        Finds all ints in every line and returns them
        """
        r = re.compile(r'[-+]?\d+')
        for line in self._lines:
            matches = r.findall(line)
            yield tuple(int(x) for x in matches)
