#!/usr/bin/python3.8
from collections import defaultdict

lines = [line.strip() for line in open('input.txt').readlines()]

G = defaultdict(set)
SMALLS = set()

for line in lines:
    parts = line.split('-')
    start, end = parts[0], parts[1]
    G[start].add(end)
    G[end].add(start)
    if start.lower() == start:
        SMALLS.add(start)

    if end.lower() == end:
        SMALLS.add(end)


def get_all_paths1(G, start, end, cur_path=None):
    if cur_path is None:
        cur_path = []
    cur_path += [start]
    if start == end:
        return [cur_path]

    all_paths = []
    for nb in G[start]:
        if nb in SMALLS and nb in cur_path:
            continue
        all_paths += get_all_paths1(G, nb, end, cur_path[:])
    return all_paths

part1_all_paths = get_all_paths1(G, 'start', 'end')
print(f'Part 1: {len(part1_all_paths)}')

RESTRICTED = {'start', 'end'}
def get_all_paths2(G, start, end, restricted=None, cur_path=None):
    if restricted is None:
        restricted = RESTRICTED.copy()
    if cur_path is None:
        cur_path = []
    cur_path += [start]
    if start == end:
        return [cur_path]

    all_paths = []
    for nb in G[start]:
        # copy here as to n0t impact other neighbours with new
        # restrictions
        nrest = restricted.copy()
        if nb in SMALLS and nb in cur_path:
            if nb in nrest:
                continue
            else:
                nrest = nrest.union(SMALLS)
        all_paths += get_all_paths2(G, nb, end, nrest, cur_path[:])
    return all_paths

part2_all_paths = get_all_paths2(G, 'start', 'end')

print(f'Part 2: {len(part2_all_paths)}')
