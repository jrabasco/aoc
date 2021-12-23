#!/usr/bin/python3.8
from framework import Parser

GOALS = {
    'A': 2,
    'B': 4,
    'C': 6,
    'D': 8
}

ROOMS = {
    2: 'A',
    4: 'B',
    6: 'C',
    8: 'D'
}

ENERGY = {
    'A': 1,
    'B': 10,
    'C': 100,
    'D': 1000
}


# Valid positions
BOARD = ['.', '.', '', '.', '', '.', '', '.', '', '.', '.']
p = Parser('input.txt')
def parse_line(line):
    return [ p for p in line.split('#') if p != '' ]
for line in p.lines(t=parse_line, f=lambda l: len(l) > 0):
    if line[0][0] == '.':
        continue
    for letter, pos in zip(line, [2,4,6,8]):
        BOARD[pos] += letter

# useful for debugging
def print_board(board, cost):
    print('#############')
    ln = '#'
    ln += ''.join(c if i not in ROOMS else '.' for i,c in enumerate(board))
    ln += f'# {cost}'
    print(ln)
    ln = '###'
    ln += '#'.join(board[r][0] for r in ROOMS)
    ln += '###'
    print(ln)
    for i in range(1, len(board[2])):
        ln = '  #'
        ln += '#'.join(board[r][i] for r in ROOMS)
        ln += '#  '
        print(ln)
    print('  #########')
    print()

# useful for debugging
def print_path(goal, prevs, costs):
    prev = prevs[goal]
    if prev is not None:
        print_path(prev, prevs, costs)
    print_board(goal, costs[goal])

# assuming start has a letter in it and it can get out of that position
def can_move(board, start, end):
    step = (end - start)//abs(end - start)
    for p in range(start+step, end+step, step):
        # skip those as they are not "real" tiles
        if p in ROOMS:
            continue
        if board[p] != '.':
            return False
    return True

def has_a_space(board, pos):
    if board[pos][0] == '.':
        return True


# assuming the move is valid
# returns the new board and cost
def move(board, start, end):
    nb = board[:]
    amph = get_amph(board, start)
    steps = abs(end-start)
    if start in ROOMS:
        depth = nb[start].count('.')
        steps += depth + 1
        nb[start] = nb[start][:depth] + '.' + nb[start][depth+1:]
    else:
        nb[start] = '.'

    if end in ROOMS:
        depth = nb[end].count('.')
        steps += depth
        nb[end] = nb[end][:depth-1] + amph + nb[end][depth:]
    else:
        nb[end] = amph

    return nb, steps * ENERGY[amph]

def room_is_friendly(board, pos, amph):
    return all( a == amph or a == '.' for a in board[pos] )

def get_amph(board, pos):
    for a in board[pos]:
        if a != '.':
            return a

def is_empty(board, pos):
    return board[pos].count('.') == len(board[pos])

def possible_moves(board, pos):
    if is_empty(board, pos):
        return []

    amph = get_amph(board, pos) # get the first amphibian in the room

    if pos == GOALS[amph] and room_is_friendly(board, pos, amph):
        return []

    res = []
    for j in range(len(board)):
        if j == pos:
            continue

        if j in ROOMS:
            if j == GOALS[amph] and room_is_friendly(board, j, amph) and can_move(board, pos, j):
                return [move(board, pos, j)]
        else:
            if pos not in ROOMS:
                # avoids going back and forth in the corridor
                continue
            valid_pos = can_move(board, pos, j)

            if valid_pos:
                new_board, cost = move(board, pos, j)
                res.append((new_board, cost))
    return res


def get_moves(board):
    res = []
    for i in range(len(board)):
        res.append(possible_moves(board, i))
    return [ m for moves in res for m in moves ]

def solve(board):
    queue = [board]
    costs = {tuple(board) : 0}
    prevs = {tuple(board) : None}
    i = 0
    progress = ['.', '..', '...']
    p_pt = 0
    while queue:
        i += 1
        if i % 2000 == 0:
            p_pt = (p_pt + 1)%3
            print('                                             ',end='\r')
            print(f'{len(queue)} states in queue{progress[p_pt]}', end='\r')
        st = queue.pop()
        cost = costs[tuple(st)]
        for n_st, n_cost in get_moves(st):
            tup = tuple(n_st)
            old_cost = costs.get(tup, 9999999999999999999999999)
            if old_cost > cost + n_cost:
                costs[tup] = cost + n_cost
                prevs[tup] = tuple(st)
                queue.append(n_st)
    print()
    return costs, prevs

MAX_IN_ROOM = max(len(s) for s in BOARD)
GOAL_BOARD = tuple(
    ROOMS[p]*MAX_IN_ROOM if p in ROOMS else '.' for p in range(len(BOARD))
)
costs, prevs = solve(BOARD)
print(f'Part 1: {costs[GOAL_BOARD]}')
#print_path(GOAL_BOARD, prevs, costs)

BOARD2 = BOARD[:]
BOARD2[2] = BOARD2[2][0] + "DD" + BOARD2[2][1]
BOARD2[4] = BOARD2[4][0] + "CB" + BOARD2[4][1]
BOARD2[6] = BOARD2[6][0] + "BA" + BOARD2[6][1]
BOARD2[8] = BOARD2[8][0] + "AC" + BOARD2[8][1]
MAX_IN_ROOM2 = max(len(s) for s in BOARD2)
GOAL_BOARD2 = tuple(
    ROOMS[p]*MAX_IN_ROOM2 if p in ROOMS else '.' for p in range(len(BOARD2))
)
costs, prevs = solve(BOARD2)
print(f'Part 2: {costs[GOAL_BOARD2]}')
#print_path(GOAL_BOARD2, prevs, costs)
