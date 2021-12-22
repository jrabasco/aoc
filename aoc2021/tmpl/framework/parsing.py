import re
SEPARATORS = [',', ' ']

class Parser:
    def __init__(self, path):
        self._lines = [line.strip() for line in open(path).readlines()]

    def lines(self):
        for line in self._lines:
            yield line[:]

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
