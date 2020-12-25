package twofive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	cardPub := 5764801
	subjectNumber := 7
	cardLS := findLoopSize(cardPub, subjectNumber)
	doorPub := 17807724
	doorLS := findLoopSize(doorPub, subjectNumber)
	ek1 := findEncryptionKey(cardLS, doorPub)
	ek2 := findEncryptionKey(doorLS, cardPub)
	assert.Equal(t, 8, cardLS)
	assert.Equal(t, 11, doorLS)
	assert.Equal(t, ek1, ek2)
}

func TestRunP1(t *testing.T) {
	cardPub := 9232416
	subjectNumber := 7
	cardLS := findLoopSize(cardPub, subjectNumber)
	doorPub := 14144084
	doorLS := findLoopSize(doorPub, subjectNumber)
	ek1 := findEncryptionKey(cardLS, doorPub)
	ek2 := findEncryptionKey(doorLS, cardPub)
	assert.Equal(t, ek1, ek2)
	t.Log(ek1)
}
