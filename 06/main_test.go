package zerosix

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b`)

	sum := run(input, 1)
	assert.Equal(t, 11, sum)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b`)

	sum := run(input, 2)
	assert.Equal(t, 6, sum)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	sum := run(f, 1)
	t.Log("p1>>>>", sum)

	f.Seek(0, io.SeekStart)
	sum = run(f, 2)
	t.Log("p1>>>>", sum)
}
