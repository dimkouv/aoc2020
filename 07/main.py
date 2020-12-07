import re

graph = {}

for line in open("input.txt").readlines():
    outer_bag = line[:line.find(" ", line.find(" ") + 1)]
    groups = re.findall(r"(\d+) ([a-z]+ [a-z]+)", line)
    graph[outer_bag] = {g[1]: int(g[0]) for g in groups}


def has_path(f, t):
    q = [f]
    while q:
        c = q.pop()
        if t in graph[c]:
            return True
        q += graph[c].keys()
    return False


def sum_weights(target):
    return sum([graph[target][n] + graph[target][n] * sum_weights(n) for n in graph[target]])


print(sum([has_path(b, "shiny gold") for b in graph]))
print(sum_weights("shiny gold"))
