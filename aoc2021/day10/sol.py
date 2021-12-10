#!/usr/bin/python3.8

lines = [line.strip() for line in open('input.txt').readlines()]

closing = { ')', ']', '}', '>' }
scores1 = {
    ')' : 3,
    ']' : 57,
    '}' : 1197,
    '>' : 25137
}

scores2 = {
    ')' : 1,
    ']' : 2,
    '}' : 3,
    '>' : 4
}

matches = {
    '(' : ')',
    '[' : ']',
    '{' : '}',
    '<' : '>'
}

res1 = 0
p2 = []
for line in lines:
    st = []
    corrupted = False
    for c in line:
        if c in matches:
            st.append(c)

        if c in closing:
            last_open = st.pop()
            if c != matches[last_open]:
                res1 += scores1[c]
                corrupted = True
                break

    if corrupted:
        continue

    res2 = 0
    while st:
        last_open = st.pop()
        needed = matches[last_open]
        res2 *= 5
        res2 += scores2[needed]
    p2.append(res2)

p2.sort()

print(f'Part 1: {res1}')
print(f'Part 2: {p2[len(p2)//2]}')
