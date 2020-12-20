#!/usr/bin/env python3

import numpy as np
import itertools
import sys


def borders(arr2d):
    return [arr2d[0, :], arr2d[-1, :], arr2d[:, 0], arr2d[:, -1]]


def they_have_shared_border(this, other):
    for tb in borders(this):
        for ob in borders(other):
            if np.array_equal(tb, ob) or np.array_equal(tb, ob[::-1]):
                return True
    return False


def tiles_with_shared_border(idx, override_tiles=None):
    this = tiles[idx]

    if override_tiles is None:
        rng = list(range(len(tiles)))
    else:
        rng = override_tiles

    s = set()
    for i in rng:
        other = tiles[i]
        if i == idx:
            continue
        if they_have_shared_border(this, other):
            s.add(i)
    return s


def is_corner(idx):
    return len(tiles_with_shared_border(idx)) == 2


tiles = []
ids = []
with open(sys.argv[1]) as f:
    tile = []
    for row in f.read().split("\n"):
        if row == "":
            tiles.append(tile)
            tile = []
        elif row.startswith("Tile"):
            ids.append(int(row.split(" ")[1][:-1]))
        else:
            tile.append([0 if c == "." else 1 for c in row])
    tiles.append(tile)

tiles = np.array(tiles)

border = set()
corners = set()
for i, other in enumerate(tiles):
    num_neibs = len(tiles_with_shared_border(i))
    if num_neibs == 2:
        corners.add(i)
        border.add(i)
    elif num_neibs == 3:
        border.add(i)

curr = corners.pop()
corners.add(curr)
ordered_corners, ordered_border = [curr], [curr]

while len(ordered_corners) < len(corners):
    for b in border:
        if b == curr or b in ordered_border:
            continue

        neibs = tiles_with_shared_border(b, list(border))
        if len(neibs) <= 2 and curr in neibs:
            curr = b
            if curr in corners:
                ordered_corners.append(curr)
            ordered_border.append(curr)
            break

n = int(tiles.shape[0] ** (1 / 2))
ordered_images = np.full((n, n), -1)
ordered_images[0, 0] = ordered_corners[0]
ordered_images[0, -1] = ordered_corners[1]
ordered_images[-1, -1] = ordered_corners[2]
ordered_images[-1, 0] = ordered_corners[3]

found = set()
curr = ordered_images[0, 0]
found.add(curr)
for j in range(1, ordered_images.shape[1]):
    ordered_images[0, j] = ordered_border[j]
    found.add(ordered_border[j])

for i in range(1, ordered_images.shape[0]):
    for b in border:
        if b in found:
            continue
        if ordered_images[i - 1, 0] in tiles_with_shared_border(b):
            ordered_images[i, 0] = b
            found.add(b)
            break

    for j in range(1, ordered_images.shape[1]):
        for t in range(len(tiles)):
            if t in found:
                continue
            left = ordered_images[i, j - 1]
            top = ordered_images[i - 1, j]
            shared = tiles_with_shared_border(t)
            if left in shared and top in shared:
                ordered_images[i, j] = t
                found.add(t)
                break

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


n = int(tiles.shape[0] ** (1 / 2))
image = np.zeros((n, n, tiles.shape[1], tiles.shape[2]))
for t1 in transformations_list:
    for t2 in transformations_list:
        ti0 = apply(t1, tiles[ordered_images[0, 0]])
        ti1 = apply(t2, tiles[ordered_images[0, 1]])
        if np.array_equal(ti0[:, -1], ti1[:, 0]):
            image[0, 0] = ti0
            image[0, 1] = ti1
            break

for j in range(2, image.shape[1]):
    i0 = image[0, j - 1]
    for t1 in transformations_list:
        ti1 = apply(t1, tiles[ordered_images[0, j]])
        if np.array_equal(i0[:, -1], ti1[:, 0]):
            image[0, j] = ti1
            break

for i in range(1, image.shape[0]):
    for j in range(0, image.shape[1]):
        i0 = image[i - 1, j]
        for t1 in transformations_list:
            ti1 = apply(t1, tiles[ordered_images[i, j]])
            if np.array_equal(i0[-1, :], ti1[0, :]):
                image[i, j] = ti1
                break

for img_i in range(image.shape[1]):
    for i in range(image.shape[3]):
        for img_j in range(image.shape[1]):
            for j in range(image.shape[2]):
                s = "#" if image[img_i, img_j, i, j] == 1 else "."
                print(s, end="")
        print("")
