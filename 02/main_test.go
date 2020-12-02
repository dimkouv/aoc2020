package zerotwo

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`)

	num, err := run(input, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, num)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`)

	num, err := run(input, 2)
	assert.NoError(t, err)
	assert.Equal(t, 1, num)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	num, err := run(f, 1)
	assert.NoError(t, err)
	t.Log("p1>>>>", num)

	f.Seek(0, io.SeekStart)
	num, err = run(f, 2)
	assert.NoError(t, err)
	t.Log("p2>>>>", num)
}
