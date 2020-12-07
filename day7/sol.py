#!/usr/bin/python3.7

lines = [ l.strip() for l in open('input.txt').readlines() ]

rules = {}
backwards_rules = {}

for line in lines:
    container, contents = line.split(" bags contain ")
    rules[container] = []
    if 'no other bags' in contents:
        continue

    contents_arr = contents.split(', ')
    for elm in contents_arr:
        qty, bag1, bag2, garbage  = elm.split(' ')
        bag = ' '.join((bag1, bag2))
        if bag not in backwards_rules:
            backwards_rules[bag] = set()

        rules[container].append((bag, int(qty)))
        backwards_rules[bag].add((container, int(qty)))

def possible_containers(bag, acc=set()):
    direct_containers = { bag for bag, qty in backwards_rules.get(bag, []) }
    acc |= direct_containers
    for c in direct_containers:
        acc |= possible_containers(c)
    return acc

def count_contents(bag):
    direct = rules[bag]
    direct_sum = sum(qty for bag, qty in direct)
    return direct_sum + sum(qty * count_contents(bag) for bag, qty in direct)

print('Part 1:')
containers = possible_containers('shiny gold')
print(len(containers))

print('Part 2:')

print(count_contents('shiny gold'))
