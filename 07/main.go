package zeroseven

import (
	"bufio"
	"io"
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
		outerBagClr, innerBugClr, num := "", "", 0
		replacer := strings.NewReplacer("bags", "", "bag", "", ".", "", ",", "")
		parts := strings.Split(strings.Trim(replacer.Replace(scanner.Text()), " "), " ")

		for _, part := range parts {
			if part == "contain" {
				outerBagClr = strings.TrimRight(outerBagClr, " ")
				nodes[outerBagClr] = map[string]int{}
			} else if n, err := strconv.Atoi(part); err == nil {
				if innerBugClr != "" {
					innerBugClr = strings.TrimRight(innerBugClr, " ")
					nodes[outerBagClr][innerBugClr] = num
				}
				innerBugClr = ""
				num = n
			} else if num == 0 {
				outerBagClr += part + " "
			} else {
				innerBugClr += part + " "
			}
		}
		if _, exists := nodes[outerBagClr]; exists {
			innerBugClr = strings.TrimRight(innerBugClr, " ")
			nodes[outerBagClr][innerBugClr] = num
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
