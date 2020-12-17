package oneseven

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	f := strings.NewReader(`.#.
..#
###`)

	grid := parseGrid(f)
	grid.cycle(6, false)
	assert.Equal(t, 112, len(grid.cubes))
}

func TestRunP1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	grid := parseGrid(f)
	grid.cycle(6, false)
	t.Log(len(grid.cubes))
}

func TestP2(t *testing.T) {
	f := strings.NewReader(`.#.
..#
###`)

	grid := parseGrid(f)
	grid.cycle(6, true)
	assert.Equal(t, 848, len(grid.cubes))
}

func TestRunP2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	grid := parseGrid(f)
	grid.cycle(6, true)
	t.Log(len(grid.cubes))
}
