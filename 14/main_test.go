package onefour

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`)
	assert.Equal(t, 165, runP1(input))
}

func TestRunP1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	t.Log(runP1(f))
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)
	assert.Equal(t, 208, runP2(input))
}

func TestRunP2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	t.Log(runP2(f))
}
