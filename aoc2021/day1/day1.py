#!/usr/bin/python3.8

numbers = [int(line.strip()) for line in open('input.txt').readlines()]
# Readable solution:
# last = None
# res1 = 0
# for nb in numbers:
#     if last is None:
#         last = nb
#         continue
#     if nb > last:
#         res1 += 1
#     last = nb
#
# print(f'Part 1 : {res1}')

# Alternative solution:
res1 = sum([ 1 for (a,b) in zip(numbers, numbers[1:]) if b > a])
print(f'Part 1 : {res1}')


# Readable solution:
# last_w_sum = None
# window = []
# res2 = 0
# for nb in numbers:
#     if len(window) < 3:
#         window.append(nb)
#         continue
#
#     if last_w_sum is None:
#         last_w_sum = sum(window)
#
#     window = window[1:] + [nb]
#     new_w_sum = sum(window)
#     if new_w_sum > last_w_sum:
#         res2 += 1
#     last_w_sum = new_w_sum
#
# print(f'Part 2 : {res2}')

# Alternative solution:
list_of_sums = [sum(x) for x in zip(numbers, numbers[1:], numbers[2:])]
res2 = sum([ 1 for (a,b) in zip(list_of_sums, list_of_sums[1:]) if b > a])
print(f'Part 2 : {res2}')
