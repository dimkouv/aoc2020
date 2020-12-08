package zeroeight

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	bc, err := newBootCode(input)
	assert.NoError(t, err)

	err = bc.run()
	assert.Equal(t, errInfLoop, err)
	assert.Equal(t, 5, bc.accumulator)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	bc, err := newBootCode(input)
	assert.NoError(t, err)
	bc.fixInfLoop()
	assert.Equal(t, 8, bc.accumulator)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	bc, err := newBootCode(f)
	assert.NoError(t, err)
	err = bc.run()
	assert.Equal(t, errInfLoop, err)
	t.Log("p1>>>>", bc.accumulator)

	f.Seek(0, io.SeekStart)
	bc.fixInfLoop()
	t.Log("p2>>>>", bc.accumulator)
}
