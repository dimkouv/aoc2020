package onesix

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	inp := strings.NewReader(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	num := parseInputData(inp).rmInvalidNearby()
	assert.Equal(t, 71, num)
}

func TestRunP1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	t.Log(parseInputData(f).rmInvalidNearby())
}

func TestP2(t *testing.T) {
	inp := strings.NewReader(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`)

	d := parseInputData(inp)
	d.rmInvalidNearby()
	positions := d.positionOfEachField()
	assert.Equal(t, 12, d.myTicket[positions["class"]])
	assert.Equal(t, 11, d.myTicket[positions["row"]])
	assert.Equal(t, 13, d.myTicket[positions["seat"]])
}

func TestRunP2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	d := parseInputData(f)
	d.rmInvalidNearby()
	p := 1
	for field, pos := range d.positionOfEachField() {
		if strings.HasPrefix(field, "departure") {
			p *= d.myTicket[pos]
		}
	}
	t.Log(p)
}

func BenchmarkP2(b *testing.B) {
	f, err := os.Open("input.txt")
	assert.NoError(b, err)
	d := parseInputData(f)
	d.rmInvalidNearby()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.positionOfEachField()
	}
}
