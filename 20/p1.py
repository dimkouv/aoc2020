import numpy as np


def borders(arr2d):
    return [arr2d[0, :], arr2d[-1, :], arr2d[:, 0], arr2d[:, -1]]


def they_have_shared_border(this, other):
    for tb in borders(this):
        for ob in borders(other):
            if np.array_equal(tb, ob) or np.array_equal(tb, ob[::-1]):
                return True
    return False


def is_corner(idx):
    this = tiles[idx]

    n = 0
    for i, other in enumerate(tiles):
        if i == idx:
            continue
        if they_have_shared_border(this, other):
            n += 1
    return n == 2


tiles = []
ids = []
with open("input.txt") as f:
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
print(np.array([ids[i] for i, _ in enumerate(tiles) if is_corner(i)]).prod())
