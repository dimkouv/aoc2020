package onezero

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`)
	num, err := runP1(input)
	assert.NoError(t, err)
	assert.Equal(t, 35, num)

	input = strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`)
	num, err = runP1(input)
	assert.NoError(t, err)
	assert.Equal(t, 220, num)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`)
	num, err := runP2(input)
	assert.NoError(t, err)
	assert.Equal(t, 8, num)

	input = strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`)
	num, err = runP2(input)
	assert.NoError(t, err)
	assert.Equal(t, 19208, num)
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	num, err := runP1(f)
	assert.NoError(t, err)
	t.Log("p1>>>>", num)

	f.Seek(0, io.SeekStart)
	num, err = runP2(f)
	assert.NoError(t, err)
	t.Log("p2>>>>", num)
}
