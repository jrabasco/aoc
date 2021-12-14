#!/usr/bin/python3.8
from collections import Counter, defaultdict

lines = [line.strip() for line in open('input.txt').readlines()]

poly_in = lines[0]

lines = lines[2:]
RULES = dict()

for line in lines:
    parts = line.split(' -> ')
    RULES[parts[0]] = parts[1]

def step(rules, poly):
    res_poly = ""
    for elm in poly:
        if len(res_poly) == 0:
            res_poly += elm
            continue

        pair = (res_poly[-1] + elm)
        if pair in rules:
            res_poly += rules[pair]

        res_poly += elm
    return res_poly

def get_sorted_counts(poly):
    counts = Counter(poly)
    return sorted([v for k, v in counts.items()])

cur_poly = poly_in[:]
for _ in range(10):
    cur_poly = step(RULES, cur_poly)

res1 = get_sorted_counts(cur_poly)
print(f'Part 1 (with step): {res1[-1] - res1[0]}')

poly_as_pairs = defaultdict(int)
prev = ""
for elm in poly_in:
    if prev == "":
        prev = elm
        continue
    poly_as_pairs[prev + elm] += 1
    prev = elm

def step_opt(rules, poly_pairs):
    res = defaultdict(int)
    for pair, count in poly_pairs.items():
        if pair in rules:
            p1, p2 = pair[0], pair[1]
            mid = rules[pair]
            res[p1 + mid] += count
            res[mid + p2] += count
    return res

def get_sorted_counts_opt(poly, poly_pairs):
    counts_raw = defaultdict(int)
    first, last = poly[0], poly[-1]
    for pair, v in poly_pairs.items():
        p1, p2 = pair[0], pair[1]
        counts_raw[p1] += v
        counts_raw[p2] += v

    counts = dict()
    for k, v in counts_raw.items():
        if k == first or k == last:
            counts[k] = (v + 1)//2
        else:
            counts[k] = v//2

    return sorted([v for k, v in counts.items()])

cur_poly_pairs = poly_as_pairs.copy()
for _ in range(10):
    cur_poly_pairs = step_opt(RULES, cur_poly_pairs)

res2 = get_sorted_counts_opt(poly_in, cur_poly_pairs)
print(f'Part 1 (with step_opt): {res2[-1] - res2[0]}')


for _ in range(10, 40):
    cur_poly_pairs = step_opt(RULES, cur_poly_pairs)

res3 = get_sorted_counts_opt(poly_in, cur_poly_pairs)
print(f'Part 2 (with step_opt): {res3[-1] - res3[0]}')
