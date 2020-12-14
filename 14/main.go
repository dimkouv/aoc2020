package onefour

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func runP1(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	mem := make(map[int]int)
	mask := [36]rune{}

	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "mask = ") {
			for i, r := range t[len("mask = "):] {
				mask[i] = r
			}
		} else {
			addr, err := strconv.Atoi(t[4:strings.Index(t, "]")])
			panicIfErr(err)
			val, err := strconv.Atoi(t[strings.Index(t, "= ")+2:])
			panicIfErr(err)
			for i := range mask {
				switch mask[i] {
				case '1':
					val |= 1 << (len(mask) - i - 1)
				case '0':
					val &= ^(1 << (len(mask) - i - 1))
				}
			}
			mem[addr] = val
		}
	}

	s := 0
	for _, v := range mem {
		s += v
	}
	return s
}

func runP2(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	mem := make(map[int]int)
	xPositions := make([]int, 0, 36)
	mask := make([]rune, 0, 36)

	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "mask = ") {
			mask = []rune(t[len("mask = "):])
			xPositions = make([]int, 0, 36)
			for i := range mask {
				if mask[i] == 'X' {
					xPositions = append(xPositions, i)
				}
			}
		} else {
			addr, err := strconv.Atoi(t[4:strings.Index(t, "]")])
			panicIfErr(err)
			val, err := strconv.Atoi(t[strings.Index(t, "= ")+2:])
			panicIfErr(err)

			maskCP := append([]rune{}, mask...)
			for j := 0; j < len(maskCP); j++ {
				if maskCP[len(maskCP)-j-1] == '0' && (addr>>(j))&1 == 1 {
					maskCP[len(maskCP)-j-1] = '1'
				}
			}

			for j := 0; j < 2<<(len(xPositions)-1); j++ {
				for z, xPos := range xPositions {
					maskCP[xPos] = '0'
					if (j>>z)&1 == 1 {
						maskCP[xPos] = '1'
					}
				}
				computedAddr, err := strconv.ParseInt(string(maskCP), 2, 64)
				panicIfErr(err)
				mem[int(computedAddr)] = val
			}
		}
	}

	s := 0
	for _, v := range mem {
		s += v
	}
	return s
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
