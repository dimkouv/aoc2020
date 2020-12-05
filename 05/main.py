seats = sorted([int(code.replace("F", "0").replace("B", "1").replace("R", "1").replace("L", "0"), 2) for code in [line for line in open("input.txt").read().split("\n") if line]])
print(max(seats), [seat+1 for i, seat in enumerate(seats[:-1]) if seats[i+1]-seats[i]>1][0])
