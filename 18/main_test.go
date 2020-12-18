package oneeight

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	tcs := []struct {
		expr string
		res  int
	}{
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.res, parse(lex(tc.expr)))
	}
}

func TestRunP1(t *testing.T) {
	b, err := ioutil.ReadFile("input.txt")
	assert.NoError(t, err)
	s := 0
	for _, l := range strings.Split(string(b), "\n") {
		s += parse(lex(l))
	}
	t.Log(s)
}

func TestP2(t *testing.T) {
	tcs := []struct {
		expr string
		res  int
	}{
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.res, parseV2(lex(tc.expr)))
	}
}

func TestRunP2(t *testing.T) {
	b, err := ioutil.ReadFile("input.txt")
	assert.NoError(t, err)
	s := 0
	for _, l := range strings.Split(string(b), "\n") {
		s += parseV2(lex(l))
	}
	t.Log(s)
}
