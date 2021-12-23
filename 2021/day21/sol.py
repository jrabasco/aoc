#!/usr/bin/python3.8
from typing import Tuple

lines = [line.strip() for line in open('input.txt').readlines()]

s0, s1 = int(lines[0].split(': ')[1]), int(lines[1].split(': ')[1])

def move(pos: int, rolls: int) -> int:
    # there is no 0 so we do -1 mod 10 + 1
    return (pos + rolls - 1) % 10 + 1


# starts on position s0, s1 and returns players' final scores and number of
# rolls
def deterministic_die(s0: int, s1: int) -> Tuple[Tuple[int, int], int]:
    roll_count = 0
    c_dice = 1
    c_player = 0
    pos = [s0, s1]
    scores = [0, 0]
    while scores[0] < 1000 and scores[1] < 1000:
        # rolling 3 consecutive numbers in a row produces 3*initial + 3 in total
        # roll numbers
        pos[c_player] = move(pos[c_player], 3*c_dice + 3)
        scores[c_player] += pos[c_player]
        roll_count += 3
        # there is no 0 so we do -1 mod 100 + 1
        c_dice = (c_dice + 2) % 100 + 1

        c_player = (c_player + 1)%2
    return tuple(scores), roll_count

scores, roll_count = deterministic_die(s0, s1)

print(f'Part 1: {min(scores) * roll_count}')

# This map will count how many ways there are to get a particular roll
ROLLS_MAP = {}
for i in range(1, 4):
    for j in range(1, 4):
        for k in range(1, 4):
            tot = i + j + k
            if tot not in ROLLS_MAP:
                ROLLS_MAP[tot] = 0
            ROLLS_MAP[tot] += 1

# Memoised method that returns the number of universes where player 0 wins
# and number of universes where player 1 wins given starting positions and
# scores
WIN_STATS = {}
def dirac_die(s0: int, s1: int, score0: int, score1: int):
    global WIN_STATS
    if (s0, s1, score0, score1) in WIN_STATS:
        return WIN_STATS[s0, s1, score0, score1]

    w0 = 0
    w1 = 0
    for rolls, cnt in ROLLS_MAP.items():
        n_s0 = move(s0, rolls)
        if n_s0 + score0 >= 21:
            w0 += cnt
        else:
            n_w1, n_w0 = dirac_die(s1, n_s0, score1, score0 + n_s0)
            w0 += cnt * n_w0
            w1 += cnt * n_w1

    WIN_STATS[s0, s1, score0, score1] = w0, w1
    return w0, w1


w0, w1 = dirac_die(s0, s1, 0, 0)
print(f'Part 2: {max(w0, w1)}')
