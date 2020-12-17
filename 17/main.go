package oneseven

import (
	"bufio"
	"fmt"
	"io"
)

type coords [4]int

type grid struct {
	cubes map[coords]bool
	dir3D [26][3]int // neighbor directions
	dir4D [80][4]int // neighbor directions
}

func (g *grid) cycle(cycles int, use4D bool) {
	for i := 0; i < cycles; i++ {
		cop := g.copyCubes()
		potentiallyActivated := make(map[coords]bool)

		// loop through active cubes, disable them according to the rule and
		// detect disabled cubes that can be potentially activated.
		for c := range g.cubes {
			active, inactive := g.getNeibs(c, use4D)
			if len(active) != 2 && len(active) != 3 {
				delete(cop, c)
			}
			for ic := range inactive {
				potentiallyActivated[ic] = true
			}
		}

		// loop through disabled cubes and activate them according to the rules.
		for c := range potentiallyActivated {
			active, _ := g.getNeibs(c, use4D)
			if len(active) == 3 {
				cop[c] = true
			}
		}

		g.cubes = make(map[coords]bool)
		for k, v := range cop {
			g.cubes[k] = v
		}
	}
}

func (g *grid) printCubes() {
	for k, v := range g.cubes {
		fmt.Println(k, v)
	}
}

func (g *grid) getNeibs(c coords, use4D bool) (active, inactive map[coords]bool) {
	active, inactive = make(map[coords]bool), make(map[coords]bool)

	if !use4D {
		for _, d := range g.dir3D {
			nCoords := coords{c[0] + d[0], c[1] + d[1], c[2] + d[2], 0}
			if g.cubes[nCoords] {
				active[nCoords] = true
			} else {
				inactive[nCoords] = true
			}
		}
	} else {
		for _, d := range g.dir4D {
			nCoords := coords{c[0] + d[0], c[1] + d[1], c[2] + d[2], c[3] + d[3]}
			if g.cubes[nCoords] {
				active[nCoords] = true
			} else {
				inactive[nCoords] = true
			}
		}
	}
	return active, inactive
}

func (g *grid) copyCubes() map[coords]bool {
	cop := make(map[coords]bool, len(g.cubes))
	for k, v := range g.cubes {
		cop[k] = v
	}
	return cop
}

func parseGrid(f io.Reader) grid {
	g := grid{cubes: make(map[coords]bool)}
	s := bufio.NewScanner(f)
	x := 0
	for s.Scan() {
		for y, r := range s.Text() {
			if r == '#' {
				g.cubes[coords{x, y, 0, 0}] = true
			}
		}
		x++
	}

	// import itertools | :'(
	i3d, i4d := 0, 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				if x != 0 || y != 0 || z != 0 {
					g.dir3D[i3d] = [3]int{x, y, z}
					i3d++
				}
				for w := -1; w < 2; w++ {
					if x != 0 || y != 0 || z != 0 || w != 0 {
						g.dir4D[i4d] = coords{x, y, z, w}
						i4d++
					}
				}
			}
		}
	}

	return g
}
