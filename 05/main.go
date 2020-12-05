package zerofive

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func computeSeatID(code string) (int, error) {
	computeNumber := func(left, right int, operations string) (int, error) {
		for _, op := range operations {
			offset := int(math.Floor((float64(right) - float64(left)) / 2.0))
			switch op {
			case 'F', 'L':
				right = left + offset
			case 'B', 'R':
				left = right - offset
			default:
				return 0, fmt.Errorf("invalid op %c in code %s", op, code)
			}
		}

		return left, nil
	}

	row, err := computeNumber(0, 127, code[:7])
	if err != nil {
		return 0, err
	}

	col, err := computeNumber(0, 7, code[7:])
	return 8*row + col, err
}

func findMissingSeat(reader io.Reader) int {
	seats := make(map[int]bool)
	for seatID := range seatsReader(reader) {
		seats[seatID] = true
	}

	for seat := range seats {
		if seats[seat+2] && !seats[seat+1] {
			return seat + 1
		}
		if seats[seat-2] && !seats[seat-1] {
			return seat - 1
		}
	}

	return -1
}

func seatsReader(reader io.Reader) chan int {
	ch := make(chan int)

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			seatID, err := computeSeatID(scanner.Text())
			if err != nil {
				panic(err)
			}
			ch <- seatID
		}

		close(ch)
	}()

	return ch
}
