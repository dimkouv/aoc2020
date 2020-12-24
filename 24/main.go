package twofour

import (
	"regexp"
)

var (
	dirRgx = regexp.MustCompile(`e|se|sw|w|nw|ne`)

	directions = map[string]complex128{
		"e":  0 + 2i,
		"se": 2 + 1i,
		"sw": 2 - 1i,
		"w":  0 - 2i,
		"nw": -2 - 1i,
		"ne": -2 + 1i,
	}
)

type floor struct {
	black map[complex128]struct{}
}

func (f *floor) invert(pathToTile string) {
	pos := 0 + 0i

	for _, dir := range dirRgx.FindAllString(pathToTile, -1) {
		pos += directions[dir]
	}

	if _, exists := f.black[pos]; exists {
		delete(f.black, pos)
	} else {
		f.black[pos] = struct{}{}
	}
}

func (f *floor) dailySwap() {
	blackCP := make(map[complex128]struct{}, len(f.black))
	for k := range f.black {
		blackCP[k] = struct{}{}
	}

	whiteCandidates := make(map[complex128]struct{})

	for pos := range f.black {
		blackNeibs, whiteNeibs := f.neibs(pos)

		if len(blackNeibs) == 0 || len(blackNeibs) > 2 {
			delete(blackCP, pos)
		}

		for whiteNeib := range whiteNeibs {
			whiteCandidates[whiteNeib] = struct{}{}
		}
	}

	for pos := range whiteCandidates {
		blackNeibs, _ := f.neibs(pos)
		if len(blackNeibs) == 2 {
			blackCP[pos] = struct{}{}
		}
	}

	f.black = blackCP
}

func (f *floor) neibs(pos complex128) (black, white map[complex128]bool) {
	black = make(map[complex128]bool)
	white = make(map[complex128]bool)

	for _, dir := range directions {
		neib := pos + dir
		if _, exists := f.black[neib]; exists {
			black[neib] = true
		} else {
			white[neib] = true
		}
	}

	return black, white
}

func newFloor() *floor {
	return &floor{black: make(map[complex128]struct{})}
}
