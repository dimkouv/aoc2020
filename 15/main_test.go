package onefive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		input   []int
		maxIter int
		exp     int
	}{
		{input: []int{0, 3, 6}, maxIter: 2020, exp: 436},
		{input: []int{1, 3, 2}, maxIter: 2020, exp: 1},
		{input: []int{2, 1, 3}, maxIter: 2020, exp: 10},
		{input: []int{1, 2, 3}, maxIter: 2020, exp: 27},
		{input: []int{2, 3, 1}, maxIter: 2020, exp: 78},
		{input: []int{3, 2, 1}, maxIter: 2020, exp: 438},
		{input: []int{3, 1, 2}, maxIter: 2020, exp: 1836},
		{input: []int{0, 3, 6}, maxIter: 30000000, exp: 175594},
		{input: []int{1, 3, 2}, maxIter: 30000000, exp: 2578},
		{input: []int{2, 1, 3}, maxIter: 30000000, exp: 3544142},
		{input: []int{1, 2, 3}, maxIter: 30000000, exp: 261214},
		{input: []int{2, 3, 1}, maxIter: 30000000, exp: 6895259},
		{input: []int{3, 2, 1}, maxIter: 30000000, exp: 18},
		{input: []int{3, 1, 2}, maxIter: 30000000, exp: 362},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("t", func(t *testing.T) {
			t.Parallel()
			res := play(tc.input, tc.maxIter)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestRun(t *testing.T) {
	t.Log(play([]int{6, 4, 12, 1, 20, 0, 16}, 2020))
	t.Log(play([]int{6, 4, 12, 1, 20, 0, 16}, 30000000))
}
