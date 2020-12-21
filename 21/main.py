# part1
foods = []

with open("input.txt") as f:
    for row in f.read().split("\n"):
        parts = row.split(" (contains ")
        ingredients = set(parts[0].split(" "))
        allergens = set(parts[1][:-1].split(", "))
        foods.append((ingredients, allergens))

all_ingredients = set()
for r in foods:
    all_ingredients.update(r[0])

candidates = dict()
for f in foods:
    for ingr in f[1]:
        if ingr in candidates:
            candidates[ingr] = candidates[ingr].intersection(f[0])
            pass
        else:
            candidates[ingr] = f[:][0]

allergenic_ingredients = set()
for c in candidates:
    allergenic_ingredients.update(candidates[c])

non_allergenic_ingredients = all_ingredients - allergenic_ingredients
cnt = 0
for f in foods:
    cnt += len(f[0].intersection(non_allergenic_ingredients))
print(cnt)

# part2
found = {}
while candidates:
    for c in candidates:
        allergen, ingredients = c, candidates[c]
        if len(ingredients) == 1:
            found[allergen] = ingredients.pop()
        candidates[c] -= set(found.values())

    for f in found:
        if f in candidates:
            del candidates[f]

allergens, ingredients = [], []
for f in found:
    allergens.append(f)
    ingredients.append(found[f])

allergens, ingredients = zip(*sorted(zip(allergens, ingredients)))
print(",".join(ingredients))
