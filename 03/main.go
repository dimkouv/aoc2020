package zerothree

import (
	"bufio"
	"io"
)

type grid struct {
	grid [][]rune
}

func newGrid() *grid {
	return &grid{}
}

func (g *grid) fromReader(reader io.Reader) (*grid, error) {
	g.grid = make([][]rune, 0)

	scanner := bufio.NewScanner(reader)
	for i := 0; scanner.Scan(); i++ {
		g.grid = append(g.grid, []rune{})
		for _, r := range scanner.Text() {
			g.grid[i] = append(g.grid[i], r)
		}
	}

	return g, nil
}

func (g *grid) navigate(pos, slope [2]int) (numTrees int) {
	for {
		pos[1] = (pos[1] + slope[1]) % len(g.grid[pos[1]])
		pos[0] = pos[0] + slope[0]

		if g.grid[pos[0]][pos[1]] == '#' {
			numTrees++
		}

		if reachedBottom := pos[0] == len(g.grid)-1; reachedBottom {
			break
		}
	}
	return numTrees
}

func runP1(reader io.Reader) (int, error) {
	g, err := newGrid().fromReader(reader)
	if err != nil {
		return 0, err
	}
	return g.navigate([2]int{0, 0}, [2]int{1, 3}), nil
}

func runP2(reader io.Reader) (int, error) {
	g, err := newGrid().fromReader(reader)
	if err != nil {
		return 0, err
	}

	slopes := [][2]int{
		{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1},
	}
	prd := 1
	for _, slope := range slopes {
		prd *= g.navigate([2]int{0, 0}, slope)
	}
	return prd, nil
}
