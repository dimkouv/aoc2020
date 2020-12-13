package onethree

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func nextDeparture(t0, bus int) int {
	return t0 + (bus - (t0 % bus))
}

func soonestBus(t0 int, buses []int) (bus, departure int) {
	bus, departure = buses[0], nextDeparture(t0, buses[0])

	for i := 1; i < len(buses); i++ {
		if buses[i] == -1 {
			continue
		}
		dep := nextDeparture(t0, buses[i])
		if dep < departure {
			bus, departure = buses[i], dep
		}
	}

	return bus, departure
}

func runP2(busses []int) (num int) {
	offsets := make([]int, 0)
	for i := len(busses) - 1; i >= 0; i-- {
		if busses[i] != -1 {
			offsets = append([]int{i}, offsets...)
		} else {
			busses = append(busses[:i], busses[i+1:]...)
		}
	}

	incr := 1
	t := busses[0]

	// after matching :patternN final busses two times update incr
	patternN := int(math.Min(float64(len(busses)-3), 4))
	patternT := 0

	for {
		numMatched := 0
		for i := len(offsets) - 1; i >= 0; i-- {
			if mod := (t + offsets[i]) % busses[i]; mod != 0 {
				break
			}
			numMatched++
			if i == 0 {
				return t
			}
		}

		if incr == 1 && numMatched >= patternN {
			switch {
			case patternT == 0:
				patternT = t
			case patternT > 0:
				incr = t - patternT
				patternN = numMatched
			}
		}

		t += incr
	}
}

func parseInput(reader io.Reader) (estimate int, buses []int, err error) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	estimate, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}

	scanner.Scan()
	parts := strings.Split(scanner.Text(), ",")
	buses = make([]int, len(parts))
	for i := range parts {
		if parts[i] == "x" {
			buses[i] = -1
		} else {
			bus, err := strconv.Atoi(parts[i])
			if err != nil {
				return 0, nil, err
			}
			buses[i] = bus
		}
	}

	return estimate, buses, nil
}
