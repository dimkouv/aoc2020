package zerofive

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		code   string
		seatID int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tc := range testCases {
		seatID, err := computeSeatID(tc.code)
		assert.NoError(t, err)
		assert.Equal(t, tc.seatID, seatID)
	}

}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	maxSeatID := 0
	for seatID := range seatsReader(f) {
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	t.Log("p1>>>>", maxSeatID)

	f.Seek(0, io.SeekStart)
	missingSeat := findMissingSeat(f)
	t.Log("p2>>>>", missingSeat)
}
