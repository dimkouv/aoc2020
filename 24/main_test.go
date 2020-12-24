package twofour

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1AndP2(t *testing.T) {
	inp := strings.NewReader(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`)

	// part1 test
	f := newFloor()
	s := bufio.NewScanner(inp)
	for s.Scan() {
		f.invert(s.Text())
	}
	assert.Len(t, f.black, 10)

	// part2 test
	for day := 0; day < 100; day++ {
		f.dailySwap()
	}
	assert.Len(t, f.black, 2208)
}

func TestRunP1(t *testing.T) {
	inp, err := os.Open("input.txt")
	assert.NoError(t, err)
	f := newFloor()
	s := bufio.NewScanner(inp)

	// part1
	for s.Scan() {
		f.invert(s.Text())
	}
	t.Log(len(f.black))

	// part2
	for day := 0; day < 100; day++ {
		f.dailySwap()
	}
	t.Log(len(f.black))
}
