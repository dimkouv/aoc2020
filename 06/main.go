package zerosix

import (
	"bufio"
	"io"
)

type counter struct {
	count      map[rune]int
	numPersons int
}

func newCounter() *counter {
	return &counter{count: make(map[rune]int)}
}

func (c *counter) addRune(r rune) {
	c.count[r] += 1
}

func (c *counter) numRunes() int {
	return len(c.count)
}

func (c *counter) numRunesEqualToPersons() int {
	sum := 0
	for _, cnt := range c.count {
		if cnt == c.numPersons {
			sum += 1
		}
	}
	return sum
}

func (c *counter) reset() {
	c.count = make(map[rune]int)
	c.numPersons = 0
}

func (c *counter) isEmpty() bool {
	return len(c.count) == 0
}

func run(reader io.Reader, part int) int {
	cnt, sum := newCounter(), 0

	var sumIncFunc func() int
	if part == 1 {
		sumIncFunc = cnt.numRunes
	} else {
		sumIncFunc = cnt.numRunesEqualToPersons
	}

	var prevRune rune
	for currRune := range runeStream(reader) {
		if prevRune == '\n' && currRune == '\n' && !cnt.isEmpty() {
			sum += sumIncFunc()
			cnt.reset()
		} else if currRune != '\n' {
			cnt.addRune(currRune)
		} else if currRune == '\n' && part == 2 {
			cnt.numPersons++
		}

		prevRune = currRune
	}

	if prevRune != '\n' { // in case last rune is not LF
		cnt.numPersons++
	}
	sum += sumIncFunc()

	return sum
}

func runeStream(reader io.Reader) chan rune {
	ch := make(chan rune)

	go func() {
		scanner := bufio.NewReader(reader)
		for {
			r, _, err := scanner.ReadRune()
			if err != nil && err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			ch <- r
		}

		close(ch)
	}()

	return ch
}
