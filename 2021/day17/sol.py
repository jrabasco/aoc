#!/usr/bin/python3.8

target_str = open('input.txt').read().strip()[13:]
parts = target_str.split(', ')
x_str, y_str = parts[0][2:], parts[1][2:]
xparts = x_str.split('..')
x0, x1 = int(xparts[0]), int(xparts[1])
yparts = y_str.split('..')
y0, y1 = int(yparts[0]), int(yparts[1])

# Get the list of possible starting horizontal speed that will not miss the
# target
def get_valid_start_xs(x0, x1):
    # any x_speed higher than this will never be in the target area
    max_x_speed = x1
    res = []
    for x in range(1, max_x_speed+1):
        cx = 0
        for i in range(x, 0, -1):
            cx += i
            if cx >= x0 and cx <= x1:
                res.append(x)
                break
    return res

def get_max_height(y):
    if y <= 0:
        return 0
    return y*(y+1)//2

def get_height_at_step(starty, step):
    cy = starty
    res = 0
    for i in range(step):
        res += cy
        cy -= 1
    return res, cy

# Assuming startx is a possible horizontal speed that will land us in the
# target, find out which steps are within the target horizontal coordinates
# no max step means the x velocity reaches 0 while still within bounds
# so height is the only limiting factor
def steps_in_bound(startx, x0, x1):
    cx = 0
    cnt = 0
    min_step = None
    for i in range(startx, 0, -1):
        cx += i
        cnt += 1
        if cx >= x0 and min_step is None:
            min_step = cnt
        if cx > x1:
            return min_step, cnt - 1
    return min_step, None

valid_xs = get_valid_start_xs(x0, x1)

possibilities = set()
max_y, max_x = -1, -1
for x in valid_xs:
    min_step, max_step = steps_in_bound(x, x0, x1)
    # assuming y0 and y1 are both negative and y0 < y1, starting with less than
    # y0 already puts the first step too low and starting with more than
    # abs(y0) means that when the probe comes back down to 0 it will have a
    # negative speed so large it' will skip the area
    for y in range(y0, abs(y0) + 1):
        step = min_step
        h, speed = get_height_at_step(y, step)
        while h >= y0 and (max_step is None or step <= max_step):
            if h >= y0 and h <= y1:
                possibilities.add((x,y))
                if y > max_y:
                    max_x, max_y = x, y
                break
            h += speed
            speed -= 1
            step += 1

print(f'Part 1: {get_max_height(max_y)}')
print(f'Part 2: {len(possibilities)}')
