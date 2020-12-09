package zeronine

import (
	"io"
	"os"
	"strings"
	"testing"

	zeroone "github.com/dimkouv/AOC2K20/01"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)

	num, err := runP1(input, 5)
	assert.NoError(t, err)
	assert.Equal(t, 127, num)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)

	num, err := runP2(input, 5)
	assert.NoError(t, err)
	assert.Equal(t, 62, num)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	num, err := runP1(f, 25)
	assert.NoError(t, err)
	t.Log("p1>>>>", num)

	f.Seek(0, io.SeekStart)
	num, err = runP2(f, 25)
	assert.NoError(t, err)
	t.Log("p2>>>>", num)
}

func BenchmarkP2(b *testing.B) {
	f, err := os.Open("input.txt")
	assert.NoError(b, err)
	nums, err := zeroone.ReadNums(f)
	assert.NoError(b, err)
	target, err := findInvalidNum(nums, 25)
	assert.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findSequence(nums, target)
	}
}
