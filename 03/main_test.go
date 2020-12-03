package zerothree

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	num, err := runP1(input)
	assert.NoError(t, err)
	assert.Equal(t, 7, num)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	num, err := runP2(input)
	assert.NoError(t, err)
	assert.Equal(t, 336, num)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	num, err := runP1(f)
	assert.NoError(t, err)
	t.Log("p1>>>>", num)

	f.Seek(0, io.SeekStart)
	num, err = runP2(f)
	assert.NoError(t, err)
	t.Log("p2>>>>", num)
}
