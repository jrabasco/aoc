#!/usr/bin/python3.8
from framework import Parser
from collections import deque
# RESEARCH
REGISTERS = {
    'w': 0,
    'x': 0,
    'y': 0,
    'z': 0
}

INPUT = deque([5, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1])

def inp(a):
    REGISTERS[a] = INPUT.popleft()

def parse(b):
    if b in REGISTERS:
        return REGISTERS[b]
    return int(b)

def add(a, b):
    REGISTERS[a] = REGISTERS[a] + parse(b)

def sub(a, b):
    REGISTERS[a] = REGISTERS[a] - parse(b)

def mul(a, b):
    REGISTERS[a] = REGISTERS[a] * parse(b)

def div(a, b):
    REGISTERS[a] = REGISTERS[a] // parse(b)

def mod(a, b):
    REGISTERS[a] = REGISTERS[a] % parse(b)

def eql(a, b):
    REGISTERS[a] = int(REGISTERS[a] == parse(b))

def neq(a, b):
    REGISTERS[a] = int(REGISTERS[a] != parse(b))

def mov(a, b):
    REGISTERS[a] = parse(b)


FUNCTIONS = {
    'inp': inp,
    'add': add,
    'mul': mul,
    'div': div,
    'mod': mod,
    'eql': eql,
    'mov': mov,
    'sub': sub,
    'neq': neq
}

REGISTERS2 = {
    'w': '0',
    'x': '0',
    'y': '0',
    'z': '0'
}

CONDITIONS = []

INPUT2 = deque(['a1', 'a2', 'a3', 'a4', 'a5', 'a6', 'a7', 'a8', 'a9', 'a10', 'a11', 'a12', 'a13', 'a14'])

def inp2(a):
    REGISTERS2[a] = INPUT2.popleft()

def parse2(b, paren=True):
    if b in REGISTERS2:
        if paren:
            return format_reg(b)
        return f'{REGISTERS2[b]}'
    return b

def format_reg(a):
    if ' ' not in REGISTERS2[a]:
        return REGISTERS2[a]
    return f'({REGISTERS2[a]})'

def add2(a, b):
    REGISTERS2[a] = f'{REGISTERS2[a]} + {parse2(b, paren=False)}'

def mul2(a, b):
    REGISTERS2[a] = f'{format_reg(a)} * {parse2(b)}'

def div2(a, b):
    REGISTERS2[a] = f'{format_reg(a)} // {parse2(b)}'

def mod2(a, b):
    REGISTERS2[a] = f'{format_reg(a)} % {parse2(b)}'

def eql2(a, b):
    REGISTERS2[a] = f'int({format_reg(a)} == {parse2(b)})'
    CONDITIONS.append(REGISTERS2[a])

def neq2(a,b):
    REGISTERS2[a] = f'int({format_reg(a)} != {parse2(b)})'
    CONDITIONS.append(REGISTERS2[a])

def mov2(a,b):
    REGISTERS2[a] = f'{parse2(b, paren=False)}'

def sub2(a, b):
    REGISTERS2[a] = f'{REGISTERS2[a]} - {parse2(b)}'

FUNCTIONS2 = {
    'inp': inp2,
    'add': add2,
    'mul': mul2,
    'div': div2,
    'mod': mod2,
    'eql': eql2,
    'mov': mov2,
    'sub': sub2,
    'neq': neq2
}

def run_program(path,min_line, max_line):
    p = Parser(path)
    for words in p.words_by_line():
        line_no = int(words[0])
        op = words[1]
        ops = words[2:]
        if line_no < min_line:
            continue
        if line_no > max_line:
            break
        FUNCTIONS[op](*ops)

def get_expr(path, min_line, max_line):
    p = Parser(path)
    for words in p.words_by_line():
        line_no = int(words[0])
        op = words[1]
        ops = words[2:]
        if line_no < min_line:
            continue
        if line_no > max_line:
            break
        FUNCTIONS2[op](*ops)

def reset(inputs, state):
    global INPUT, REGISTERS
    INPUT = deque(inputs)
    REGISTERS = {
        'w': state[0],
        'x': state[1],
        'y': state[2],
        'z': state[3]
    }

res = []

def compile():
    # manually simplified program
    get_expr('input_opt.txt', 0, 251)
    with open('compiled.py', 'w+') as f:
        f.write('def z(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14):\n')
        f.write(f'    return {REGISTERS2["z"]}\n')

#from compiled import z


# Worked out from input.work.txt
#a5 == a4 - 1
#a6 == a3 - 4
#a8 == a7 + 8
#a10 == a9 + 4
#a12 == a11 + 3
#a13 == a2 + 1
#a14 == a1 - 2

def find_max():
    # a1 - 2 > 0 -> a1 > 2
    for a1 in range(9, 2, -1):
        # a2 + 1 <= 9 -> a2 <= 8
        for a2 in range(8, 0, -1):
            # a3 - 4 > 0 -> a3 > 4
            for a3 in range(9, 4, -1):
                # a4 - 1 > 0 -> a4 > 1
                for a4 in range(9, 1, -1):
                    # a7 + 8 <= 9 -> a7 <= 1
                    for a7 in range(1, 0, -1):
                        # a9 + 4 <= 9 -> a9 <= 5
                        for a9 in range(5, 0, -1):
                            # a11 + 3 <= 9 -> a11 <= 6
                            for a11 in range(6, 0, -1):
                                a5 = a4 - 1
                                a6 = a3 - 4
                                a8 = a7 + 8
                                a10 = a9 + 4
                                a12 = a11 + 3
                                a13 = a2 + 1
                                a14 = a1 - 2
                                return a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14

def find_min():
    # a1 - 2 >= 1 -> a1 >= 3
    for a1 in range(3, 10):
        # a2 + 1 < 10 -> a2 < 9
        for a2 in range(1, 9):
            # a3 - 4 >= 1 -> a3 >= 5
            for a3 in range(5, 10):
                # a4 - 1 >= 1 -> a4 >= 2
                for a4 in range(2, 10):
                    # a7 + 8 < 10 -> a7 < 2
                    for a7 in range(1, 2):
                        # a9 + 4 < 10 -> a9 < 6
                        for a9 in range(1, 6):
                            # a11 + 3 < 10 -> a11 < 7
                            for a11 in range(1, 7):
                                a5 = a4 - 1
                                a6 = a3 - 4
                                a8 = a7 + 8
                                a10 = a9 + 4
                                a12 = a11 + 3
                                a13 = a2 + 1
                                a14 = a1 - 2
                                return a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14
print(f'Part 1: {"".join(str(a) for a in find_max())}')
print(f'Part 2: {"".join(str(a) for a in find_min())}')


