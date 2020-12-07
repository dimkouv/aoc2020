package zeroseven

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type graph struct {
	nodes map[string]map[string]int
}

func newGraph() *graph {
	return &graph{
		nodes: map[string]map[string]int{},
	}
}

func (g *graph) fromReader(reader io.Reader) *graph {
	scanner := bufio.NewScanner(reader)
	nodes := make(map[string]map[string]int)

	for scanner.Scan() {
		t := scanner.Text()
		outerBagClr := t[:strings.Index(t, " bag")]
		nodes[outerBagClr] = make(map[string]int)

		innerBagRgx := regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+)")
		innerBagMatches := innerBagRgx.FindAllStringSubmatch(t, -1)
		for i := range innerBagMatches {
			innerBagCnt, _ := strconv.Atoi(innerBagMatches[i][1])
			innerBagClr := innerBagMatches[i][2]
			nodes[outerBagClr][innerBagClr] = innerBagCnt
		}
	}

	g.nodes = nodes
	return g
}

func (g *graph) hasPath(from, to string) bool {
	queue := []string{from}
	var current string

	for len(queue) > 0 {
		current = queue[0]
		queue = append(queue[:0], queue[1:]...)

		for neighbor := range g.nodes[current] {
			if neighbor == to {
				return true
			}

			queue = append(queue, neighbor)
		}
	}

	return false
}

func (g *graph) sumWeights(node string) int {
	sum := 0
	for neighbor, w := range g.nodes[node] {
		sum += w + w*g.sumWeights(neighbor)
	}

	return sum
}
