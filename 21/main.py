# part1
foods = []

with open("input.txt") as f:
    for row in f.read().split("\n"):
        parts = row.split(" (contains ")
        ingredients = set(parts[0].split(" "))
        allergens = set(parts[1][:-1].split(", "))
        foods.append((ingredients, allergens))

candidates = dict()
for f in foods:
    for alrg in f[1]:
        if alrg in candidates:
            candidates[alrg] = candidates[alrg].intersection(f[0])
        else:
            candidates[alrg] = f[0]

allergenic_ingredients = set().union(*[candidates[c] for c in candidates])
all_ingredients = set().union(*[r[0] for r in foods])
non_allergenic_ingredients = all_ingredients - allergenic_ingredients

# part1
print(sum([len(f[0].intersection(non_allergenic_ingredients)) for f in foods]))

found = {}
while candidates:
    for c in list(candidates.keys()):
        allergen, ingredients = c, candidates[c]
        if len(ingredients) == 1:
            found[allergen] = ingredients.pop()
            if allergen in candidates:
                del candidates[allergen]
        else:
            candidates[c] -= set(found.values())

allergens, ingredients = [], []
for f in found:
    allergens.append(f)
    ingredients.append(found[f])

allergens, ingredients = zip(*sorted(zip(allergens, ingredients)))

# part2
print(",".join(ingredients))
