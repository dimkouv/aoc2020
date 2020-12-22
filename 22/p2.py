with open("input.txt") as f:
    lines = f.read().replace("\n", " ").replace("Player 2:", "#").replace("Player 1: ", "")
    p1, p2 = lines.strip().split("  # ")
    p1 = [int(c) for c in p1.split(" ")]
    p2 = [int(c) for c in p2.split(" ")]


def recursive_combat(p1, p2, depth=1):
    history = set()
    k = (tuple(p1), tuple(p2))

    while p1 and p2:
        if k in history:
            return 1
        history.add(k)

        c1 = p1.pop(0)
        c2 = p2.pop(0)

        if c1 <= len(p1) and c2 <= len(p2):
            winner = recursive_combat(p1[:c1][:], p2[:c2][:], depth=depth+1)
        else:
            winner = 1 if c1 > c2 else 2

        if winner == 1:
            p1 += [c1, c2]
        else:
            p2 += [c2, c1]

        k = (tuple(p1), tuple(p2))

    if depth == 1:
        return p1, p2
    else:
        return 1 if p1 else 2


P1, P2 = recursive_combat(p1, p2)
print(sum([(i+1)*v for i, v in enumerate(reversed(P1+P2))]))
