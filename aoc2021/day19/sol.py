#!/usr/bin/python3.8
from matrix import Matrix, Vector
from typing import Set, Tuple, Optional, Dict, List
from collections import deque
from dataclasses import dataclass

@dataclass
class Scanner:
    tag: int
    beacons: Set[Vector]
    pos: Optional[Vector]

bases_lst = [
    [[ 1, 0, 0], [ 0, 1, 0], [ 0, 0, 1]],
    [[ 0, 1, 0], [-1, 0, 0], [ 0, 0, 1]],
    [[-1, 0, 0], [ 0,-1, 0], [ 0, 0, 1]],
    [[ 0,-1, 0], [ 1, 0, 0], [ 0, 0, 1]],
    [[ 0, 0, 1], [ 0, 1, 0], [-1, 0, 0]],
    [[-1, 0, 0], [ 0, 1, 0], [ 0, 0,-1]],
    [[ 0, 0,-1], [ 0, 1, 0], [ 1, 0, 0]],
    [[ 1, 0, 0], [ 0, 0, 1], [ 0,-1, 0]],
    [[ 1, 0, 0], [ 0,-1, 0], [ 0, 0,-1]],
    [[ 1, 0, 0], [ 0, 0,-1], [ 0, 1, 0]],
    [[-1, 0, 0], [ 0, 0, 1], [ 0, 1, 0]],
    [[-1, 0, 0], [ 0, 0,-1], [ 0,-1, 0]],
    [[ 0, 1, 0], [ 1, 0, 0], [ 0, 0,-1]],
    [[ 0, 1, 0], [ 0, 0, 1], [ 1, 0, 0]],
    [[ 0, 1, 0], [ 0, 0,-1], [-1, 0, 0]],
    [[ 0,-1, 0], [-1, 0, 0], [ 0, 0,-1]],
    [[ 0,-1, 0], [ 0, 0, 1], [-1, 0, 0]],
    [[ 0,-1, 0], [ 0, 0,-1], [ 1, 0, 0]],
    [[ 0, 0, 1], [ 0,-1, 0], [ 1, 0, 0]],
    [[ 0, 0, 1], [ 1, 0, 0], [ 0, 1, 0]],
    [[ 0, 0, 1], [-1, 0, 0], [ 0,-1, 0]],
    [[ 0, 0,-1], [ 0,-1, 0], [-1, 0, 0]],
    [[ 0, 0,-1], [ 1, 0, 0], [ 0,-1, 0]],
    [[ 0, 0,-1], [-1, 0, 0], [ 0, 1, 0]]
]

# first base is the one we assume for the first scanner
BASE0 = Matrix(bases_lst[0])
bases = {BASE0}
for base in bases_lst[1:]:
    bases.add(Matrix(base))

lines = [line.strip() for line in open('input.txt').readlines()]
scanners = []
offs = -1
for line in lines:
    if line == '':
        continue
    if '---' in line:
        offs += 1
        scanners.append(Scanner(offs,set(),None))
        continue
    beacon = Vector([int(nb) for nb in line.split(',')])
    scanners[offs].beacons.add(beacon)

ALL_BEACONS = scanners[0].beacons.copy()
POS0 = Vector([0,0,0])
scanners[0].pos = POS0
known = deque([scanners[0]])

print('Rotating scanners...')
ROTATED_SCANNERS = dict()
for sc in scanners[1:]:
    ROTATED_SCANNERS[sc.tag] = []
    for base in bases:
        if base == BASE0:
            ROTATED_SCANNERS[sc.tag].append(sc)
            continue
        r_beacons = {base*b for b in sc.beacons}
        ROTATED_SCANNERS[sc.tag].append(Scanner(sc.tag, r_beacons, None))


unknown = []
for sc in scanners[1:]:
    unknown.append(sc)

def rotate(y: Vector, A: Matrix) -> Vector:
    return A*y

def find_match_and_transform(rotated_scanners: Dict[int, List[Scanner]],
                   known_scanner: Scanner,
                   unknown_scanner: Set[Vector]) -> Tuple[Matrix, Scanner]:
    for r_scan in rotated_scanners[unknown_scanner.tag]:
        for v0 in known_scanner.beacons:
            for v1 in r_scan.beacons:
                # assume v1 == v0
                sc_pos = v0 - v1
                new_beacs = {vec + sc_pos for vec in r_scan.beacons}
                if len(new_beacs.intersection(known_scanner.beacons)) >= 12:
                    return base, Scanner(unknown_scanner.tag, new_beacs, sc_pos)

while unknown:
    print(f'Unknown: {len(unknown)}, known: {len(known)}')
    sc = known.popleft()
    known.append(sc)
    still_unknown = []
    for usc in unknown:
        print(f'Trying {sc.tag} vs {usc.tag}')
        res = find_match_and_transform(ROTATED_SCANNERS, sc, usc)
        if res is None:
            still_unknown.append(usc)
            continue

        _ , ksc = res
        for v in ksc.beacons:
            ALL_BEACONS.add(v)
        known.appendleft(ksc)

    unknown = still_unknown

for sc in known:
    print(sc.tag, sc.pos)

print(f'Part 1: {len(ALL_BEACONS)}')

distances = [
    a.pos.manhattan_distance(b.pos)
    for a in known
    for b in known
    if a.tag != b.tag
]
print(f'Part 2: {max(distances)}')
