import numpy as np

# Solution based on Chinese Remainder Theorem
# explanation: https://www.youtube.com/watch?v=zIFehsBHB8o

with open("input.txt") as f:
    inp = f.readlines()[1].replace("\n", "").split(",")
    buses = np.array([int(v) for v in inp if v != "x"])
    offsets = np.array([inp.index(str(v)) for v in buses])

def find_xi(ni, m): # finds xi such that (ni * xi) % m == 1
    for xi in range(ni):
        if (ni * xi) % m == 1:
            return xi

N = buses.prod()
print("N =", N)
s = 0
for i in range(len(buses)):
    b_i = 0 if offsets[i] == 0 else buses[i] - offsets[i]
    n_i = N // buses[i]
    x_i = find_xi(n_i, buses[i])
    s += b_i * n_i * x_i
    print("b%d=%d\tn%d=%d\tx%d=%d" % (i, b_i, i, n_i, i, x_i))

res = s % N
print("s=%d\ns %% N = %d" % (s, res))
