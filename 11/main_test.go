package oneone

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	floor, err := newFloor(input, true, 4)
	assert.NoError(t, err)

	floor.simulateUntilNoChange()
	assert.Equal(t, 37, floor.occupiedSeats)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	floor, err := newFloor(input, false, 5)
	assert.NoError(t, err)

	floor.simulateUntilNoChange()
	assert.Equal(t, 26, floor.occupiedSeats)
}

func TestRunP1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	floor, err := newFloor(f, true, 4)
	assert.NoError(t, err)

	floor.simulateUntilNoChange()
	t.Log(">>>", floor.occupiedSeats)
}

func TestRunP2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	floor, err := newFloor(f, false, 5)
	assert.NoError(t, err)

	floor.simulateUntilNoChange()
	t.Log(">>>", floor.occupiedSeats)
}
