import numpy as np
import itertools

data = np.loadtxt("input.txt")

# part1
for i, j in itertools.combinations([i for i in range(len(data))], 2):
    if i != j and data[i]+data[j] == 2020:
        print(data[i]*data[j])

# part2
for i, j, z in itertools.combinations([i for i in range(len(data))], 3):
    if i != j != z and data[i]+data[j]+data[z] == 2020:
        print(data[i]*data[j]*data[z])
