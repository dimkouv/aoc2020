package zeroone

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	f := strings.NewReader(`
1721
979
366
299
675
1456
`)
	num, err := runP1(f)
	assert.NoError(t, err)
	assert.Equal(t, 514579, num)
}

func TestP2(t *testing.T) {
	f := strings.NewReader(`
1721
979
366
299
675
1456
`)
	num, err := runP2(f)
	assert.NoError(t, err)
	assert.Equal(t, 241861950, num)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	r := bufio.NewReader(f)
	numP1, err := runP1(r)
	assert.NoError(t, err)
	t.Log("p1>>>", numP1)

	f.Seek(0, io.SeekStart)
	r = bufio.NewReader(f)
	numP2, err := runP2(r)
	assert.NoError(t, err)
	t.Log("p2>>>", numP2)
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("input.txt")
		assert.NoError(b, err)
		r := bufio.NewReader(f)
		_, err = runP1(r)
		assert.NoError(b, err)

		f.Seek(0, io.SeekStart)
		r = bufio.NewReader(f)
		_, err = runP2(r)
		assert.NoError(b, err)
	}
}
