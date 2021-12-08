#!/usr/bin/python3.8
lines = [line.strip() for line in open('input.txt').readlines()]

count = 0
for line in lines:
    parts = line.split(' | ')
    patterns, output = parts[0].split(' '), parts[1].split(' ')

    count += sum(1 for _ in output if len(_) == 2 or len(_) == 4 or len(_) == 3 or len(_) == 7)

print(f'Part 1: {count}')

res = 0
for line in lines:
    parts = line.split(' | ')
    patterns, output = parts[0].split(' '), parts[1].split(' ')

    lengths = {}
    for p in patterns:
        if len(p) not in lengths:
            lengths[len(p)] = []
        # make them into frozensets for easy diff/union operations
        lengths[len(p)].append(frozenset(p))

    # one, four, seven and eight are the ones who have unique lenghts
    one = lengths[2][0]
    four = lengths[4][0]
    seven = lengths[3][0]
    eight = lengths[7][0]
    two_three_five = lengths[5]
    zero_six_nine = lengths[6]

    #
    #  --           --
    #    |   |  |     |
    #  --  -  --  -     =
    # |         |     |   |
    #  --                  --
    #
    two = [ nb for nb in two_three_five if len(nb - four - seven) == 2 ][0]
    #
    #  --     --
    #    |      |
    #  --  -  --  =
    #    |   |        |
    #  --     --
    #
    three = [ nb for nb in two_three_five if len(nb - two) == 1 ][0]
    # five is the last
    five = [ nb for nb in two_three_five if nb != two and nb != three ][0]

    #
    #  --            --
    #    |   |  |   |  |
    #  --  &  --  =  --
    #    |      |      |
    #  --            --
    #
    nine = three.union(four)

    #
    #  --     --
    #    |      |
    #  --  -  --  =
    # |         |   |
    #  --     --
    #
    #  --            --
    # |             |
    #  --  &      =  --
    #    |   |      |  |
    #  --            --
    six = five.union(two - three)
    # zero is the last
    zero = [ nb for nb in zero_six_nine if nb != nine and nb != six ][0]

    alphabet = {}

    alphabet[zero] = '0'
    alphabet[one] = '1'
    alphabet[two] = '2'
    alphabet[three] = '3'
    alphabet[four] = '4'
    alphabet[five] = '5'
    alphabet[six] = '6'
    alphabet[seven] = '7'
    alphabet[eight] = '8'
    alphabet[nine] = '9'

    res += int(''.join(alphabet[frozenset(dig)] for dig in output))

print(f'Part 2: {res}')
