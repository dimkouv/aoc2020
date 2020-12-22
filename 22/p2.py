with open("input.txt") as f:
    deck1, deck2 = list(map(lambda x: list(map(int, x.split(":\n")[1].split("\n"))), f.read().strip().split("\n\n")))

def recursive_combat(deck1, deck2):
    history = set()

    while deck1 and deck2:
        k = (tuple(deck1), tuple(deck2))
        if k in history:
            return 1

        history.add(k)
        card1, card2 = deck1.pop(0), deck2.pop(0)

        if card1 <= len(deck1) and card2 <= len(deck2):
            winner = recursive_combat(deck1[:card1], deck2[:card2])
        else:
            winner = 1 if card1 > card2 else 2

        deck1 += [card1, card2] if winner == 1 else []
        deck2 += [card2, card1] if winner == 2 else []

    return 1 if deck1 else 2

recursive_combat(deck1, deck2)
print(sum([(i+1) * v for i, v in enumerate(reversed(deck1 + deck2))]))
