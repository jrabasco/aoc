#!/usr/bin/python3.8
from grid import InfiniteGrid
from typing import List
from PIL import Image

def to_bin(char: str) -> int:
    if char == '.':
        return '0'
    return '1'

lines = [line.strip() for line in open('input.txt').readlines()]

CONVERTER = list(lines[0])
# skip the empty line
G = InfiniteGrid(lines[2:], '.')
def enhance(g: InfiniteGrid, converter: List[str]):
    n_g = []
    for i in range(-1, g.h+1):
        n_g.append([])
        for j in range(-1, g.w+1):
            offs = int(''.join(to_bin(c) for c in g.square(i,j)), 2)
            n_g[i+1].append(converter[offs])

    default_offs = int(to_bin(g.get(-1,-1))*9, 2)
    n_default = converter[default_offs]
    return InfiniteGrid(n_g, n_default)


def save_image(g: InfiniteGrid, step: int):
    offs = (50 - step)
    width = g.w + offs
    height = g.h + offs
    img = Image.new('RGB', (width*2, height*2), "black")
    pixels = img.load()
    for i in range(width):
        for j in range(height):
            if g.get(j,i) == '#':
                pixels[(2*i+offs)%(width*2),(2*j+offs)%(height*2)] = (255, 255, 255)
                pixels[(2*i+1+offs)%(width*2),(2*j+offs)%(height*2)] = (255, 255, 255)
                pixels[(2*i+offs)%(width*2),(2*j+1+offs)%(height*2)] = (255, 255, 255)
                pixels[(2*i+1+offs)%(width*2),(2*j+1+offs)%(height*2)] = (255, 255, 255)
    img.save(f'imgs/step{step:02}.png')

for i in range(50):
    G = enhance(G, CONVERTER)
    save_image(G, i+1)
    if i == 1:
        print(f'Part 1: {G.count(lambda x: x == "#")}')
print(f'Part 2: {G.count(lambda x: x == "#")}')
