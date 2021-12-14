package day9

import (
	"math"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

type point struct {
	x, y int
}

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 9, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day9, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	lines := input.GetStringSlice()
	points := make(map[int]map[int]int)
	for i := 0; i < len(lines); i++ {
		points[i] = map[int]int{}
	}

	for x, line := range lines {
		for y, value := range line {
			points[x][y] = util.AtoIorEXIT(string(value))
		}
	}

	width := len(points[0])
	height := len(points)
	lowPoints := []point{}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			up, okup := points[x][y-1]
			down, okdown := points[x][y+1]
			left, okleft := points[x-1][y]
			right, okright := points[x+1][y]
			value := points[x][y]
			if okup && value >= up {
				continue
			}
			if okdown && value >= down {
				continue
			}
			if okleft && value >= left {
				continue
			}
			if okright && value >= right {
				continue
			}
			lowPoints = append(lowPoints, point{x: x, y: y})
		}
	}

	basinSizes := make([]int, len(lowPoints))

	for i, point := range lowPoints {
		basinSizes[i] = calcBasin(points, point)
	}

	big, bigger, biggest := math.MinInt32, math.MinInt32, math.MinInt32
	for _, value := range basinSizes {
		if value > biggest {
			big = bigger
			bigger = biggest
			biggest = value
		} else if value > bigger {
			big = bigger
			bigger = value
		} else if value > big {
			big = value
		}
	}
	return big * bigger * biggest
}

func calcBasin(points map[int]map[int]int, lowPoint point) int {
	visited := map[point]bool{}
	basin := 0

	possible := util.NewQueue()
	for _, neighbour := range getNeighbours(points, lowPoint) {
		possible.Push(neighbour)
	}
	for possible.Length() > 0 {
		pp, ok := possible.Pop()
		if !ok {
			// Pop empty queue
			continue
		}
		p := pp.(point)
		if _, alreadyVisited := visited[p]; alreadyVisited {
			continue
		}
		if points[p.x][p.y] == 9 {
			// Any point with the value 9 does not belong here
			continue
		}
		basin++
		visited[p] = true
		for _, neighbour := range getNeighbours(points, p) {
			possible.Push(neighbour)
		}
	}
	return basin
}

func getNeighbours(points map[int]map[int]int, p point) (neighbours []point) {
	if _, ok := points[p.x][p.y-1]; ok {
		neighbours = append(neighbours, point{x: p.x, y: p.y - 1})
	}
	if _, ok := points[p.x][p.y+1]; ok {
		neighbours = append(neighbours, point{x: p.x, y: p.y + 1})
	}
	if _, ok := points[p.x-1][p.y]; ok {
		neighbours = append(neighbours, point{x: p.x - 1, y: p.y})
	}
	if _, ok := points[p.x+1][p.y]; ok {
		neighbours = append(neighbours, point{x: p.x + 1, y: p.y})
	}
	return neighbours
}
