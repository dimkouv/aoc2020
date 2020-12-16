package onesix

import (
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type inputData struct {
	ranges        map[string][2][2]int
	myTicket      []int
	nearbyTickets [][]int
}

// Removes invalid nearby tickets and returns the sum of their values.
func (inp *inputData) rmInvalidNearby() int {
	s := 0
	for i := len(inp.nearbyTickets) - 1; i >= 0; i-- {
		for _, num := range inp.nearbyTickets[i] {
			match := false
			for _, rng := range inp.ranges {
				r1, r2 := rng[0], rng[1]
				if (num >= r1[0] && num <= r1[1]) || (num >= r2[0] && num <= r2[1]) {
					match = true
					break
				}
			}

			if !match {
				s += num
				inp.nearbyTickets = append(inp.nearbyTickets[:i], inp.nearbyTickets[i+1:]...)
				break
			}
		}
	}

	return s
}

// Detects and returns the position of every ticket field.
// On every iteration we detect fields that have only one valid position and
// 		we assign that position to that specific field until we find all the field positions.
func (inp *inputData) positionOfEachField() map[string]int {
	validPositions := make(map[string][]int)
	rm := make(map[int]bool)
	for field := range inp.ranges {
		positions := inp.validPositions(field)
		if len(positions) == 1 {
			rm[positions[0]] = true
		}
		validPositions[field] = positions
	}

	for {
		found := 0
		for field, vps := range validPositions {
			if len(vps) <= 1 {
				found++
				continue
			}

			for i := len(vps) - 1; i >= 0; i-- {
				if rm[vps[i]] {
					validPositions[field] = append(validPositions[field][:i], validPositions[field][i+1:]...)
				}
			}
			if len(validPositions[field]) == 1 {
				rm[validPositions[field][0]] = true
			}
		}

		if found == len(validPositions) {
			break
		}
	}

	positions := make(map[string]int, len(validPositions))
	for field := range validPositions {
		positions[field] = validPositions[field][0]
	}
	return positions
}

// Returns the valid positions for a specific field, we call this once
// before starting the iterations.
func (inp *inputData) validPositions(field string) []int {
	ranges := inp.ranges[field]
	r1, r2 := ranges[0], ranges[1]
	positions := make([]int, 0, len(inp.nearbyTickets[0]))

	for pos := range inp.nearbyTickets[0] {
		match := true
		for _, tick := range inp.nearbyTickets {
			num := tick[pos]
			if !((num >= r1[0] && num <= r1[1]) || (num >= r2[0] && num <= r2[1])) {
				match = false
				break
			}
		}
		if match {
			positions = append(positions, pos)
		}
	}

	return positions
}

func parseInputData(r io.Reader) *inputData {
	inp := &inputData{ranges: make(map[string][2][2]int, 0), myTicket: make([]int, 0), nearbyTickets: make([][]int, 0)}
	b, _ := ioutil.ReadAll(r)
	s := string(b)

	rgx := regexp.MustCompile(`([a-z ]+): (\d+-\d+) or (\d+-\d+)`)
	for _, match := range rgx.FindAllStringSubmatch(s, -1) {
		parts := strings.Split(match[2], "-")
		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])
		parts = strings.Split(match[3], "-")
		n3, _ := strconv.Atoi(parts[0])
		n4, _ := strconv.Atoi(parts[1])
		inp.ranges[match[1]] = [2][2]int{{n1, n2}, {n3, n4}}
	}

	p1 := strings.Index(s, "t:\n") + 3
	p2 := p1 + strings.Index(s[p1:], "\n")
	parts := strings.Split(s[p1:p2], ",")
	for i := range parts {
		n, _ := strconv.Atoi(parts[i])
		inp.myTicket = append(inp.myTicket, n)
	}

	for _, ticketRaw := range strings.Split(s[p2:], "\n")[3:] {
		parts := strings.Split(ticketRaw, ",")
		ticket := make([]int, len(parts))
		for i := range parts {
			n, _ := strconv.Atoi(parts[i])
			ticket[i] = n
		}
		inp.nearbyTickets = append(inp.nearbyTickets, ticket)
	}

	return inp
}
