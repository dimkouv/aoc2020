package onefive

func play(inp []int, maxIter int) int {
	pos := make(map[int][2]int)

	num := 0
	for i := 0; i < maxIter; i++ {
		if i >= len(inp) {
			prevPos := pos[num]
			switch seenAtLeastTwice := prevPos[1] > 0; seenAtLeastTwice {
			case true:
				num = prevPos[1] - prevPos[0]
			default:
				num = 0
			}
		} else {
			num = inp[i]
		}

		p, exists := pos[num]
		switch {
		case !exists: // never seen
			pos[num] = [2]int{i + 1}
		case p[1] == 0: // seen once
			pos[num] = [2]int{p[0], i + 1}
		default: // already seen twice
			pos[num] = [2]int{p[1], i + 1}
		}
	}

	return num
}
