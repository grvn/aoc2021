package day9

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 9, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day9, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
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
	risk := 0

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			up, okup := points[x][y-1]
			down, okdown := points[x][y+1]
			left, okleft := points[x-1][y]
			right, okright := points[x+1][y]
			point := points[x][y]
			if okup && point >= up {
				continue
			}
			if okdown && point >= down {
				continue
			}
			if okleft && point >= left {
				continue
			}
			if okright && point >= right {
				continue
			}
			risk += point + 1
		}
	}
	return risk
}
