package onethree

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`939
7,13,x,x,59,x,31,19
`)

	t0, buses, err := parseInput(input)
	assert.NoError(t, err)

	bus, departure := soonestBus(t0, buses)
	assert.Equal(t, 59, bus)
	assert.Equal(t, 295, bus*(departure-t0))
}

func TestRunP1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	t0, buses, err := parseInput(f)
	assert.NoError(t, err)

	bus, departure := soonestBus(t0, buses)
	t.Log("Bus: ", bus, "Departure: ", departure)
	t.Log("Num: ", bus*(departure-t0))
}

func TestP2(t *testing.T) {
	testCases := []struct {
		inp string
		exp int
	}{
		{"0\n7,13,x,x,59,x,31,19", 1068781},
		{"0\n17,x,13,19", 3417},
		{"0\n67,7,59,61", 754018},
		{"0\n67,x,7,59,61", 779210},
		{"0\n67,7,x,59,61", 1261476},
		{"0\n1789,37,47,1889", 1202161486},
	}

	for _, tc := range testCases {
		t.Run(tc.inp, func(t *testing.T) {
			input := strings.NewReader(tc.inp)
			_, buses, err := parseInput(input)
			assert.NoError(t, err)
			t1 := runP2(buses)
			assert.Equal(t, tc.exp, t1, tc)
		})
	}
}

func TestRunP2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	_, buses, err := parseInput(f)
	assert.NoError(t, err)

	num := runP2(buses)
	t.Log("Num: ", num)
}
