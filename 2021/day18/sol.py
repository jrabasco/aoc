#!/usr/bin/python3.8
import json
from functools import reduce

numbers = [json.loads(line.strip()) for line in open('input.txt').readlines()]

def split(elm):
    return [elm//2, elm//2 + elm%2]

def explode(elm, lst, i):
    a, b = elm
    carry_left, carry_right = 0, 0
    if i > 0:
        lst[i-1] = add_to_the_right(lst[i-1], a)
    else:
        carry_left = a

    if i < (len(lst)-1):
        lst[i+1] = add_to_the_left(lst[i+1], b)
    else:
        carry_right = b
    lst[i] = 0
    return carry_left, carry_right

def add_to_the_right(elm, val):
    if type(elm) == int:
        return elm + val
    else:
        return elm[:-1] + [add_to_the_right(elm[-1], val)]

def add_to_the_left(elm, val):
    if type(elm) == int:
        return elm + val
    else:
        return [add_to_the_left(elm[0], val)] + elm[1:]

def reduce_split_only(nb):
    res = nb[:]
    for i, elm in enumerate(res):
        if type(elm) == int:
            if elm > 9:
                res[i] = split(elm)
                return res, True
        else:
            res[i], changed = reduce_split_only(elm)
            if changed:
                return res, changed
    return res, False

def reduce_explode_only(nb, depth=0):
    res = nb[:]
    for i, elm in enumerate(res):
        if type(elm) == list:
            if depth >= 3:
                carry_left, carry_right = explode(elm, res, i)
                return res, True, carry_left, carry_right

            res[i], changed, carry_left, carry_right = reduce_explode_only(elm, depth+1)
            if i > 0 and carry_left > 0:
                res[i-1] = add_to_the_right(res[i-1], carry_left)
                carry_left = 0

            if i < (len(res)-1) and carry_right > 0:
                res[i+1] = add_to_the_left(res[i+1], carry_right)
                carry_right = 0

            if changed:
                return res, changed, carry_left, carry_right

    return res, False, 0, 0

def red(nb, depth=0):
    res, changed, _, __ = reduce_explode_only(nb)
    if changed:
        return res, changed
    res, changed = reduce_split_only(nb)
    return res, changed

def magnitude(nb):
    if type(nb) == int:
        return nb
    return 3*magnitude(nb[0]) + 2*magnitude(nb[1])

def add(a, b):
    acc = [a, b]
    acc, changed = red(acc)
    while changed:
        acc, changed = red(acc)
    return acc

p1 = magnitude(reduce(add, numbers))
print(f'Part 1: {p1}')
p2 = max(magnitude(add(left, right)) for i, left in enumerate(numbers) for j, right in enumerate(numbers) if j!=i)
print(f'Part 2: {p2}')
