#!/usr/bin/python3.8
from dataclasses import dataclass
from typing import List

@dataclass
class Cell:
    value: int
    marked: bool

class Board:
    def __init__(self, cells: List[List[Cell]]):
        self._row_cells = cells
        self._col_cells = []
        for i in range(5):
            self._col_cells.append([])
        for row in self._row_cells:
            for i, cell in enumerate(row):
                self._col_cells[i].append(cell)

    def has_won(self) -> bool:
        row_win = any(all(cell.marked for cell in row) for row in self._row_cells)
        col_win = any(all(cell.marked for cell in col) for col in self._col_cells)
        return row_win or col_win

    def __str__(self) -> str:
        lines = ['rows:']
        for row in self._row_cells:
            lines.append(f"{', '.join(map(str, row))}")
        lines.append('columns:')
        for col in self._col_cells:
            lines.append(f"{', '.join(map(str, col))}")
        return "\n".join(lines)

    def score(self, last: int) -> int:
        return sum(sum(cell.value for cell in row if not cell.marked) for row in self._row_cells) * last

    def mark(self, called: int):
        for row in self._row_cells:
            for cell in row:
                if cell.value == called:
                    cell.marked = True

lines = [line.strip() for line in open('input.txt').readlines()]

first_line = lines[0]
numbers = [int(nb) for nb in first_line.split(",")]
lines = [ line for line in lines[2:] if line != "" ]
boards = []

while len(lines) > 0:
    board_lines = lines[:5]
    lines = lines[5:]
    cells = [ [ Cell(int(val), False) for val in line.split(" ") if val != ""] for line in board_lines]
    boards.append(Board(cells))

for nb in numbers:
    done = False
    for board in boards:
        board.mark(nb)
        if board.has_won():
            done = True
            print(f'Part 1: {board.score(nb)}')
            break
    if done:
        break

for nb in numbers:
    for board in boards:
        board.mark(nb)

    if len(boards) == 1 and boards[0].has_won():
        print(f'Part 2: {board.score(nb)}')
        break

    boards = [board for board in boards if not board.has_won()]
