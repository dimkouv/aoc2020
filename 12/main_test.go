package onetwo

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`F10
N3
F7
R90
F11
`)

	ship := newShip()
	ship.processSteps(input)
	assert.Equal(t, 25, ship.manhattanDistance())
}

func TestRunP1(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	ship := newShip()
	ship.processSteps(f)
	t.Log(ship.manhattanDistance())
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`F10
N3
F7
R90
F11
`)

	ship := newShip()
	ship.processStepsWithWaypoint(input)
	assert.Equal(t, 286, ship.manhattanDistance())
}

func TestRunP2(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)

	ship := newShip()
	ship.processStepsWithWaypoint(f)
	t.Log(ship.manhattanDistance())
}
