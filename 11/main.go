package oneone

import (
	"bufio"
	"io"
)

const (
	itemEmptySeat    rune = 'L'
	itemNonEmptySeat rune = '#'
	itemFloor        rune = '.'
)

type grid struct {
	items         [][]rune
	occupiedSeats int

	adjacentAlg     bool // if false part2 alg is used
	personTolerance int
}

var directions = [][2]int{
	{-1, -1}, // top left
	{-1, 0},  // top center
	{-1, +1}, // top right
	{0, -1},  // left
	{0, +1},  // right
	{+1, -1}, // bot left
	{+1, 0},  // bot center
	{+1, +1}, // bot right
}

func (f *grid) simulateUntilNoChange() {
	for f.runSimulationRound() > 0 {
	}
}

func (f *grid) runSimulationRound() int {
	numChanges := 0
	f.occupiedSeats = 0
	itemsCopy := f.itemsCopy()

	for i := range f.items {
		for j := range f.items[i] {
			item := f.items[i][j]

			occupied := f.numOccupiedSeats(i, j)
			if item == itemEmptySeat && occupied == 0 {
				itemsCopy[i][j] = itemNonEmptySeat
				numChanges++
			} else if item == itemNonEmptySeat && occupied >= f.personTolerance {
				itemsCopy[i][j] = itemEmptySeat
				numChanges++
			}

			if itemsCopy[i][j] == itemNonEmptySeat {
				f.occupiedSeats++
			}
		}
	}
	f.items = itemsCopy
	return numChanges
}

func (f *grid) numOccupiedSeats(i, j int) int {
	if f.adjacentAlg {
		return f.numAdjacentOccupiedSeats(i, j)
	}
	return f.numVisibleOccupiedSeats(i, j)
}

func (f *grid) numAdjacentOccupiedSeats(i, j int) int {
	cnt := 0
	for _, dir := range directions {
		x, y := dir[0]+i, dir[1]+j
		isValid := x >= 0 && y >= 0 && x < len(f.items) && y < len(f.items[0])
		if isValid && f.items[x][y] == itemNonEmptySeat {
			cnt++
		}
	}
	return cnt
}

func (f *grid) numVisibleOccupiedSeats(i, j int) int {
	cnt := 0
	for _, dir := range directions {
		x, y := i, j
		for { // walk in the direction until you find a wall or a seat
			x, y = dir[0]+x, dir[1]+y
			isValid := x >= 0 && y >= 0 && x < len(f.items) && y < len(f.items[0])
			if !isValid || f.items[x][y] == itemEmptySeat {
				break
			} else if f.items[x][y] == itemNonEmptySeat {
				cnt++
				break
			}
		}
	}
	return cnt
}

func (f *grid) itemsCopy() [][]rune {
	itemsCopy := make([][]rune, len(f.items))
	for i := range f.items {
		itemsCopy[i] = make([]rune, len(f.items[i]))
		for j := range f.items[i] {
			itemsCopy[i][j] = f.items[i][j]
		}
	}
	return itemsCopy
}

func newFloor(reader io.Reader, adjacentAlg bool, tolerance int) (*grid, error) {
	f := &grid{items: make([][]rune, 0), adjacentAlg: adjacentAlg, personTolerance: tolerance}
	scanner := bufio.NewScanner(reader)
	for i := 0; scanner.Scan(); i++ {
		f.items = append(f.items, []rune{})
		for _, r := range scanner.Text() {
			f.items[i] = append(f.items[i], r)
		}
	}
	return f, nil
}
