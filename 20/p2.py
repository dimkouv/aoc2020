#!/usr/bin/env python3

import itertools
import numpy as np
import fileinput

# part 2 runs in two steps, first run p2_img.py to generate the image
# and then run p2.py to count sea monsters.
#
# ./p2_img.py input.txt | ./p2.py

img = []

for line in fileinput.input():
    if line != "":
        img.append([1 if c == "#" else 0 for c in line if line != ""])

img = np.array(img)

n = img.shape[0]
dc, dr = 0, 0
for i in range(0, n, 10):
    img = np.delete(img, i - dc, 1)
    dc += 1
    img = np.delete(img, i - dr, 0)
    dr += 1
    img = np.delete(img, i - dc + 9, 1)
    dc += 1
    img = np.delete(img, i + 9 - dr, 0)
    dr += 1

sea_monster = [
    "                  # ",
    "#    ##    ##    ###",
    " #  #  #  #  #  #   "
]

sm_pos = []
for i, row in enumerate(sea_monster):
    for j, c in enumerate(row):
        if c == "#":
            sm_pos.append((i, j - 18))

# transforms = [rot90 rot90 rot90 flipH flipV]
transformations_list = []
for tl in itertools.product([True, False], repeat=5):
    for t in itertools.permutations(np.nonzero(tl)[0]):
        transformations_list.append(t)


def apply(transformations, arr):
    cp = arr.copy()
    for t in transformations:
        if t in (0, 1, 2):
            cp = np.rot90(cp)
        elif t == 3:
            cp = np.flip(cp, axis=0)
        elif t == 4:
            cp = np.flip(cp, axis=1)
    return cp


for t1 in transformations_list:
    t_img = apply(t1, img)
    num_sea_monsters = 0
    for i in range(t_img.shape[0]):
        for j in range(t_img.shape[1]):
            found = True
            for p in sm_pos:
                o_i, o_j = p
                try:
                    if t_img[i + o_i, j + o_j] != 1:
                        found = False
                        break
                except IndexError:
                    found = False
                    break

            if found:
                num_sea_monsters += 1

    if num_sea_monsters > 0:
        print("num sea monster:", num_sea_monsters)
        occupied_by_sea_monsters = num_sea_monsters * 15
        print(np.count_nonzero(img == 1) - occupied_by_sea_monsters)
        break
