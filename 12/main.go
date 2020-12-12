package onetwo

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

const (
	east int = iota // (!) order of declarations matters
	south
	west
	north

	left
	right
	forward
)

type stepType struct {
	action int
	value  int
}

type ship struct {
	direction int
	position  [4]int
}

func (s *ship) processSteps(reader io.Reader) {
	for step := range stepStream(reader) {
		switch step.action {
		case forward:
			s.position[s.direction] += step.value
		case north, east, south, west:
			s.position[step.action] += step.value
		case right, left:
			rotations := (step.value / 90) % (north + 1)
			if step.action == left {
				rotations = (north + 1) - rotations // for not messing with negative rotations
			}
			s.direction = (s.direction + rotations) % (north + 1)
		}
	}
}

func (s *ship) processStepsWithWaypoint(reader io.Reader) {
	waypoint := [4]int{east: 10, south: 0, west: 0, north: 1}
	waypointCheckpoint := [4]int{east: 10, south: 0, west: 0, north: 1}

	for step := range stepStream(reader) {
		switch step.action {
		case forward:
			for dir := east; dir <= north; dir++ {
				s.position[dir] += waypoint[dir] * step.value
			}
		case east, south, west, north:
			waypoint[step.action] += step.value
		case right, left:
			rotations := (step.value / 90) % (north + 1)
			if step.action == left {
				rotations = (north + 1) - rotations // for not messing with negative rotations
			}
			for dir := east; dir <= north; dir++ {
				waypointCheckpoint[dir] = waypoint[dir]
			}
			for dir := east; dir <= north; dir++ {
				waypoint[(dir+rotations)%(north+1)] = waypointCheckpoint[dir]
			}
		}
	}
}

func (s *ship) manhattanDistance() int {
	dx := math.Abs(float64(s.position[east] - s.position[west]))
	dy := math.Abs(float64(s.position[south] - s.position[north]))
	return int(dx + dy)
}

func newShip() *ship {
	return &ship{}
}

func stepStream(reader io.Reader) chan stepType {
	ch := make(chan stepType)

	mapp := map[uint8]int{'W': west, 'S': south, 'N': north, 'E': east, 'L': left, 'R': right, 'F': forward}
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			row := scanner.Text()
			value, err := strconv.Atoi(row[1:])
			if err != nil {
				panic(err)
			}
			ch <- stepType{action: mapp[row[0]], value: value}
		}
		close(ch)
	}()

	return ch
}
