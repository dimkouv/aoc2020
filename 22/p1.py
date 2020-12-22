with open("input.txt") as f:
    lines = f.read().replace("\n", " ").replace("Player 2:", "#").replace("Player 1: ", "")
    p1, p2 = lines.strip().split("  # ")
    p1 = [int(c) for c in p1.split(" ")]
    p2 = [int(c) for c in p2.split(" ")]

while p1 and p2:
    c1 = p1.pop(0)
    c2 = p2.pop(0)

    if c1 > c2:
        p1 += [c1, c2]
    else:
        p2 += [c2, c1]

print(sum([(i+1)*v for i, v in enumerate(reversed(p1+p2))]))
