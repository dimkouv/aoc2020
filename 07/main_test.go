package zeroseven

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	input := strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	graph := newGraph().fromReader(input)

	const target = "shiny gold"
	cnt := 0
	for node := range graph.nodes {
		if node != target && graph.hasPath(node, target) {
			cnt++
		}
	}
	assert.Equal(t, 4, cnt)
}

func TestP2(t *testing.T) {
	input := strings.NewReader(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`)
	graph := newGraph().fromReader(input)

	const target = "shiny gold"
	assert.Equal(t, 126, graph.sumWeights(target))
}

func TestRun(t *testing.T) {
	f, err := os.Open("input.txt")
	assert.NoError(t, err)
	graph := newGraph().fromReader(f)

	const target = "shiny gold"
	cnt := 0
	for node := range graph.nodes {
		if node != target && graph.hasPath(node, target) {
			cnt++
		}
	}
	t.Log("p1>>>>", cnt)

	f.Seek(0, io.SeekStart)
	t.Log("p2>>>>", graph.sumWeights(target))
}
